package gotest

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestStringTrimRight(t *testing.T) {
	s := "滴滴、百度、腾讯、"
	strings.TrimRight(s, "、")
	t.Log(s)
	s = strings.TrimRight(s, "、")
	t.Log(s)
}


func TestInt64ToString(t *testing.T) {
	logId := time.Now().UnixNano()
	logIdStr := string(logId)
	logIdByte, _ := json.Marshal(logId)
	t.Log(logIdStr)
	t.Log(string(logIdByte))

}


func StructToString(v interface{}) string {
	resByte, _ := json.Marshal(v)
	return string(resByte)
}

func TestTransToFloat64(t *testing.T) {
	str := "0.65"
	if strFloat, err := TransToFloat64(str); err != nil {
		t.Error(err)
	} else {
		t.Log(strFloat)
	}
}

func TransToFloat64(data interface{}) (res float64, err error) {
	val := reflect.ValueOf(data)
	switch data.(type) {
	case int, int8, int16, int32, int64:
		res = float64(val.Int())
	case uint, uint8, uint16, uint32, uint64:
		res = float64(val.Uint())
	case float64:
		res = data.(float64)
	case float32:
		res = float64(data.(float32))
	case string:
		res, err = strconv.ParseFloat(strings.TrimSpace(data.(string)), 64)
	default:
		res, err = strconv.ParseFloat(fmt.Sprintf("%v", data), 64)
	}
	return
}