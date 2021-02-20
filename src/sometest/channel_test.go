package sometest

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	for i := 0; i < 10 ; i++ {
		goChannel()
	}
	return
}

func goChannel () {
	//ch1 := make(chan struct{})
	ch1 := make(chan struct{}, 1)
	go func() {
		defer func(){
			fmt.Println("get ch1 before")
			//此处，如果ch1为阻塞channel，超时后协程会在此阻塞
			ch1 <- struct{}{}
			fmt.Println("get ch1 done")
		}()
		time.Sleep(time.Duration(1600) * time.Millisecond)
		fmt.Println("get ch1 result")
	}()

	select {
	case <-ch1:
		return
	case <-time.After(time.Duration(1500) * time.Millisecond):
		fmt.Println("time out ")
		return
		/* default:
		   fmt.Println("time default")*/

	}
	return
}
