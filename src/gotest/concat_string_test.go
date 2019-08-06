package gotest

import (
	"bytes"
	"testing"
)

func TestConcatStringByAdd(t *testing.T) {
	elems := []string{"1", "2", "3", "4", "5"}
	ret := ""
	for _, elem := range elems {
		ret += elem
	}
	if ret == "12345" {
		t.Log("true")
	}
}

func TestConcatStringByBytesBuffer(t *testing.T) {
	var buf bytes.Buffer
	elems := []string{"1", "2", "3", "4", "5"}
	for _, elem := range elems {
		buf.WriteString(elem)

	}
	if  buf.String() == "12345" {
		t.Log("true")
	}
}


func BenchmarkConcatStringByAdd(b *testing.B) {

	elems := []string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ret := ""
		for _, elem := range elems {
			ret += elem
		}
	}
	b.StopTimer()
}

