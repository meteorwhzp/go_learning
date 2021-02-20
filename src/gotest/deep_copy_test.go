package gotest

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"testing"
)

type TestStruct struct {
	Name string
	Feature map[string]string
}

func TestDeepCopy(t *testing.T) {
	str1 := &TestStruct{
		Name: "www",
		Feature: make(map[string]string),
	}
	str1.Feature["w"] = "y"
	func1(str1)

	str2 := &TestStruct{}

	DeepCopy(str2, str1)


	fmt.Printf("%p\n",str1.Feature)
	fmt.Printf("%p\n",str2.Feature)
	fmt.Println(str2)

}


func func1(dc *TestStruct) {
	fmt.Printf("%p\n",dc)
	dc.Feature["w"] = "z"
}

//对象深拷贝
func DeepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}