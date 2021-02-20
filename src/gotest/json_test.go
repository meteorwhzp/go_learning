package gotest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestJsonNumber(t *testing.T) {
	s := "{\"a\":6673221165400540161}"
	d := make(map[string]interface{})
	d, _ = TransToMap(s)

	fmt.Println(TransToInt64(d["a"]))
}

func TransToMap(param interface{}) (map[string]interface{}, error) {
	resultMap := make(map[string]interface{})
	switch s_type := param.(type) {
	case string:
		decoder := json.NewDecoder(bytes.NewReader([]byte(string(s_type))))
		decoder.UseNumber()
		err := decoder.Decode(&resultMap)
		if err != nil {
			return nil, err
		}
	case map[string]interface{}:
		resultMap = param.(map[string]interface{})
	default:
		return nil, errors.New(fmt.Sprintf("ToMap is err. input not string, not map type:%v", s_type))
	}
	return resultMap, nil
}

func TransToInt64(data interface{}) (res int64, err error) {
	fmt.Println(reflect.TypeOf(data))
	switch data.(type) {
	case float64:
		res = int64(data.(float64))
	case int:
		res = int64(data.(int))
	case uint:
		res = int64(data.(uint))
	case int64:
		res = data.(int64)
	case json.Number:
		res, err = strconv.ParseInt(strings.TrimSpace(string(data.(json.Number))), 10, 64)
	case string:
		res, err = strconv.ParseInt(strings.TrimSpace(data.(string)), 10, 64)
	}
	return
}