package main

import (
	"fmt"
	"time"
)

/*
#include <stdio.h>
void printint(int v) {
	printf("printint: %d\n", v);
}
*/

import "C"


func main() {
	//v := 42
	//c.printint(C.int(v))
}

func main1() {
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