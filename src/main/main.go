package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello world")
	var count int
	ch1 := make(chan struct{}, 1)
	fmt.Println("time now is", count)
	go func() {
		d := 1 * time.Second
		ticker := time.Tick(d)
		for range ticker {
			fmt.Println("time now is", count)
			count++
		}
	}()

	select {
	case <-ch1:
		return
	case <-time.After(time.Duration(1500) * time.Millisecond):
		fmt.Println("time out ")
		return


	}

}