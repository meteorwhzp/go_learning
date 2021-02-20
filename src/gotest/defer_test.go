package gotest

import "testing"

func TestDefer(t *testing.T) {
	i := 1
	defer func() {
		if err := recover(); err != nil {
			t.Logf("panic err is %v", err)
			t.Logf("i is %d", i)
		}
	}()


	defer func() {
/*		defer func() {
			if err := recover(); err != nil {
				t.Logf("panic err is %v", err)
			}
		}()*/
		t.Log("before panic 1")
		t.Log("before panic 1")

		panic("errno2")
	}()

	panic("errno1")
}
