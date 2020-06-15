package sometest

import (
	"fmt"
	"strings"
	"testing"
)

func TestForRangeErr1(t *testing.T) {
	src := []int{1,2,3,4,5}
	var dst2 []*int
	for _, i := range src {
		dst2 = append(dst2, &i)
	}
	res := ""
	for _, p := range dst2 {
		res = fmt.Sprintf("%s%d", res, *p)
	}
	t.Log(res)
}

func TestForRangeRight(t *testing.T) {
	src := []int{1,2,3,4,5}
	var dst2 []*int
	for _, i := range src {
		tmp := i
		dst2 = append(dst2, &tmp)
	}
	res := ""
	for _, p := range dst2 {
		res = fmt.Sprintf("%s%d", res, *p)
	}
	t.Log(res)
}

func TestEid(t *testing.T) {
	eid := "\"e6962c3f14341e03a71f607b7bddaa64\""
	eid = "e6962c3f14341e03a71f607b7bddaa64"
	t.Log(eid)
	reqEstimateId := strings.Trim(eid, "\"")
	t.Log(reqEstimateId)
}

func TestEe(t *testing.T) {
	Dim1 := 1e-16
	Dim2 := 1E-16

	t.Log(Dim1)
	t.Log(Dim2)
}

