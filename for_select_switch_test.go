package examples_go

import (
	"fmt"
	"testing"
)

func TestForSelect(t *testing.T) {
	forSelect()

}

func TestForSelectBreak(t *testing.T) {
	var i int = 0
	forSelectBreak(i)

}

func TestForSelectContinue(t *testing.T) {
	var i int = 0
	forSelectContinue(i)

}

func TestForBreak(t *testing.T) {
	var i int = 0
	forBreak(i)
	fmt.Println("TestForBreak finish")
}

func TestForContinue(t *testing.T) {
	var i int = 0
	forContinue(i)
	fmt.Println("TestForContinue finish")
}

func TestSwitch(t *testing.T) {
	i := 3
	Switch(i)
}

func TestForSwitchBreak(t *testing.T) {
	forSwitchBreak()
}

func TestForSwitchContinue(t *testing.T) {
	forSwitchContinue()
}

func TestBreakOut(t *testing.T) {
	forBreakOut()
}
