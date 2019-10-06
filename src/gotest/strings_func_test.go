package gotest

import (
	"strings"
	"testing"
)

func TestStringTrimRight(t *testing.T) {
	s := "滴滴、百度、腾讯、"
	strings.TrimRight(s, "、")
	t.Log(s)
	s = strings.TrimRight(s, "、")
	t.Log(s)
}
