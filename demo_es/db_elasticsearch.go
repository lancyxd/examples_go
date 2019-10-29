package demo_es

import (
	"encoding/json"
	"examples_go/demo_kafka/common/logging"
	"fmt"
	"gopkg.in/olivere/elastic.v3"
)

var (
	ClientES *elastic.Client
)

type queryItem struct {
	field string
	text  string
	boost float64
	value interface{}
	btime string
	etime string
}

type QueryList struct {
	ItemArray []queryItem
	fieldAggs string
	fieldTime string
	Scale     string
	Offset    string
	BoostMode string
	timeFrom  string
	timeTo    string
	From      int
	Size      int
	decay     float64
}

//init
func InitES(urls []string, username, password string) error {
	var err error
	ClientES, err = elastic.NewClient(elastic.SetBasicAuth(username, password), elastic.SetURL(urls...))
	if err != nil {
		return err
	}
	return nil
}

func (queryIn *QueryList) Query() ([]map[string]interface{}, error) {
	var bQuery *elastic.BoolQuery
	highlight := elastic.NewHighlight()
	order := OrderBy(queryIn.ItemArray[0].field)
	for _, v := range queryIn.ItemArray {
		highlight.Field(v.field)
		bQuery = ShouldQuery(v.field, v.text, v.boost, v.value)
	}

	aggsQ := elastic.NewTermsAggregation().Field(queryIn.fieldAggs).Size(queryIn.Size)

	//print es
	funcQ := elastic.NewFunctionScoreQuery().Query(bQuery).
		AddScoreFunc(elastic.NewGaussDecayFunction().FieldName(queryIn.fieldTime).
			Origin(queryIn.timeFrom).Offset(queryIn.Offset).Scale(queryIn.Scale).
			Decay(queryIn.decay)).BoostMode(queryIn.BoostMode)

	searcher1 := elastic.NewSearchSource()
	searcher1.Query(funcQ).Highlight(highlight). // specify the query
							SortBy(order).
							From(queryIn.From).
							Size(queryIn.Size).Aggregation(queryIn.fieldAggs, aggsQ)
	jSearch1, err := searcher1.Source()
	if err != nil {
		return nil, err
	}
	debug1, _ := json.Marshal(jSearch1)
	logging.Debug("searcher1 query body: %s", string(debug1))

	//searcher := elastic.NewFunctionScoreQuery()
	order2 := OrderBy("_score")
	//order3 := OrderScript("make")
	searcher := elastic.NewSearchSource()
	searcher.
		Query(bQuery).
		Highlight(highlight). // specify the query
		SortBy(order, order2).
		From(queryIn.From).
		Size(queryIn.Size).Aggregation(queryIn.fieldAggs, aggsQ)

	jSearch, err := searcher.Source()
	if err != nil {
		return nil, err
	}
	debug, _ := json.Marshal(jSearch)
	logging.Debug("searcher query body: %s", string(debug))

	// search
	res, err := ClientES.Search().Highlight(highlight).
		Index("cars", "index2"). // search in index "twitter"
		Query(bQuery).           // specify the query
		SortBy(order, order2).
		From(0).
		Size(2).Aggregation("colors", aggsQ).
		Pretty(true). // pretty print request and response JSON
		Do()          // executed

	//查询结果不为0时，进行处理至doc中
	docList := []map[string]interface{}{}

	if res != nil {
		if len(res.Hits.Hits) > 0 {
			for _, v := range res.Hits.Hits {
				doc := map[string]interface{}{}
				json.Unmarshal(*v.Source, &doc)
				fmt.Println(doc)
				docList = append(docList, doc)
			}
		}
	}

	return docList, nil
}

/*
query sentence
*/

//should
func ShouldQuery(field, text string, boost float64, value interface{}) *elastic.BoolQuery {
	Bquery := elastic.NewBoolQuery()
	matchQ := elastic.NewMatchQuery(field, text).Boost(boost)
	matchPhraseQ := elastic.NewMatchPhraseQuery(field, text)
	matchPrefixQ := elastic.NewPrefixQuery(field, text)
	matchPhrasePrefixQ := elastic.NewMatchPhrasePrefixQuery(field, text)
	matchWildcardQ := elastic.NewWildcardQuery(field, "*"+text+"*")
	multiMatchQ := elastic.NewMultiMatchQuery(text, field, field).Boost(boost)
	termCQ := elastic.NewConstantScoreQuery(elastic.NewTermQuery(field, value)).Boost(boost) //termQ,value:"440606"
	termQ := elastic.NewTermQuery(field, value).Boost(boost)
	Bquery.Should(matchQ, matchPhraseQ, matchPrefixQ, matchPhrasePrefixQ, matchWildcardQ, multiMatchQ, termCQ, termQ)
	return Bquery
}

//order
func OrderBy(field string) elastic.FieldSort {
	order := elastic.NewFieldSort(field).Desc()
	return order
}

//highlight
func Highlight(field1, field2 string) *elastic.Highlight {
	highlight := elastic.NewHighlight()
	highlight.Field(field1)
	highlight.Field(field2)
	return highlight
}

//script
func OrderScript(field string) elastic.ScriptSort {
	q := elastic.NewScriptSort(elastic.NewScript("doc['"+field+"']"+".value * factor").Param("factor", 1.1), "number").Desc()
	return q
}

//dismax_query
func disMaxQuery() *elastic.DisMaxQuery {
	q := elastic.NewDisMaxQuery()
	q = q.Query(elastic.NewTermQuery("age", 34), elastic.NewTermQuery("age", 35)).Boost(1.2).TieBreaker(0.7)

	return q
}

//must
func MustQuery(field, from, to string) *elastic.BoolQuery {
	Bquery := elastic.NewBoolQuery()
	id := []interface{}{"1", "2"}
	termQ := elastic.NewTermsQuery(field, id...)
	RangeQ := elastic.NewRangeQuery(field).From(from).To(to)
	existQ := elastic.NewExistsQuery(field)

	//filterQ
	Squery := elastic.NewBoolQuery()
	Squery.Filter(elastic.NewMatchQuery(field, "guang"))
	q := elastic.NewGeoDistanceQuery("location")
	q = q.GeoPoint(elastic.GeoPointFromLatLon(40, -70))
	q = q.Distance("200km")
	Squery.Filter(q)

	Bquery.Must(termQ, RangeQ, existQ, Squery)

	return Bquery
}

//must_not
func MustNotQuery(brands []string) *elastic.BoolQuery {
	Bquery := elastic.NewBoolQuery()
	termsNQ := elastic.NewTermsQuery("brand", brands[0], brands[1])
	Bquery.MustNot(termsNQ)
	return Bquery
}

//创建索引别名
func CreateIndexAlias(srcIndex []string, alias string) error {
	//判断该别名是否存在，存在无法创建
	indexExists, err := ClientES.IndexExists(alias).Do()
	if err != nil {
		logging.Error("[CreateIndexAlias] IndexExists error,err=%+v", err)
		return err
	}
	if indexExists {
		logging.Error("[CreateIndexAlias] index %s should not exist, but does\n", alias)
		return err
	}

	//创建索引别名
	for _, v := range srcIndex {
		aliasCreate, err := ClientES.Alias().
			Add(v, alias).
			Do()
		if err != nil {
			logging.Error("[CreateIndexAlias]  Alias failed,err=%+v", err)
			return err
		}

		if !aliasCreate.Acknowledged {
			logging.Error("[CreateIndexAlias] aliasCreate.Acknowledged failed,err=%+v", err)
			return err
		}
	}
	logging.Debug("[CreateIndexAlias] ok,alias=%s", alias)
	return nil
}

/*
// Add both indices to a new alias
aliasCreate, err := EsClient.Alias().
	Add("xxxx_srcname", "xxxx_bf").
	Action(elastic.NewAliasAddAction("xxxx_bf").Index("xxxx_srcname_test")).
	Do()
if err != nil {
	logging.Error("EsClient.Alias failed,err=%+v",err)
	return
}
if !aliasCreate.Acknowledged {
	logging.Error("expected AliasResult.Acknowledged %v; got %v", true, aliasCreate.Acknowledged)
	return
}
*/
