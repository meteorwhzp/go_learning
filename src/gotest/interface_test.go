package gotest

import (
	"fmt"
	"reflect"
	"testing"
)

type Phone interface {
	Call()
}

type IPhone struct {

}


func (phone IPhone)Call(){
	fmt.Println("phone")
}

func TestInterface(t *testing.T) {
	var a Phone
	b := &IPhone{}
	b = nil
	a = b

	t.Log(reflect.TypeOf(a))
	t.Log(reflect.ValueOf(a).Kind())
	t.Log(reflect.TypeOf(a))
	t.Log(reflect.ValueOf(a))

	if a == nil {
		t.Log("true")
	}
}
