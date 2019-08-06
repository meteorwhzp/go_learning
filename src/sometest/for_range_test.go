package sometest

import (
	"fmt"
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
