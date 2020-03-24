package examples_go

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"encoding/json"
	"github.com/labstack/echo"
)

type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
	//Name  string `json:"name,omitempty"`
	//Email string `json:"email,omitempty"`
}

func main() {
	e := echo.New()

	e.HTTPErrorHandler = customHTTPErrorHandler

	//handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/users/:id", GetUser) //请求参数 http://localhost:1323/users/Joe?team=x-men&member=wolverine

	e.POST("/save", save)
	e.POST("/savef", saveF)
	e.POST("/user/set", setUser)

	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &CustomContext{c}
			return h(cc)
		}
	})
	e.GET("/custom", CustomMid)

	e.Logger.Fatal(e.Start(":1323"))

}

func GetUser(c echo.Context) error {
	id := c.Param("id")
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, id+":"+team+":"+member)
}

//表单 application/x-www-form-urlencoded
// curl -F "name=Joe Smith" -F "email=joe@labstack.com" http://localhost:1323/save
func save(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:"+name+", email:"+email)
}

//表单  multipart/form-data
//curl -F "name=Joe Smith" -F "pic=@/home/xxx/project/go/src/pic.jpg" http://localhost:1323/savef
func saveF(c echo.Context) error {
	name := c.FormValue("name")
	pic, err := c.FormFile("pic")
	if err != nil {
		return err
	}
	src, err := pic.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(pic.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, "<b>Thank you! "+name+"</b>")
}

//请求body为json
//curl -XPOST 'http://localhost:1323/users' -d'{"name":"lily","email":"Smith@qq.com"}'
func setUser(c echo.Context) error {
	//body, _ := c.Get("req_body").([]byte)
	body, _ := ioutil.ReadAll(c.Request().Body)
	payload := new(User)
	//payload := &User{}
	json.Unmarshal(body, payload)
	return c.JSON(http.StatusOK, payload)

	// or
	// return c.XML(http.StatusCreated, u)
}

//自定义context
type CustomContext struct {
	echo.Context
}

func (c *CustomContext) Foo() {
	println("foo excute")
}

func (c *CustomContext) Bar() {
	println("bar excute")
}

func CustomMid(c echo.Context) error {
	cc := c.(*CustomContext)
	cc.Foo()
	cc.Bar()
	return cc.String(200, "OK")
}

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	errorPage := fmt.Sprintf("%d.html", code)
	if err := c.File(errorPage); err != nil {
		fmt.Println("err:", err.Error())
		c.Logger().Error(err)
	}
	c.Logger().Error(err)
}
