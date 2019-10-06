package gotest

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeRange(t *testing.T) {
	timeNow := time.Now().Unix()
	fmt.Println("time now is start", timeNow)
	ch1 := make(chan int, 1)
	go func() {
		count := 0
		d := 2 * time.Second
		ticker := time.Tick(d)
		for range ticker {
			if count > 10 {
				ch1 <- 1
				break
			}
			fmt.Printf("time now is %d, %v \n", count, time.Now().Unix())
			count++
		}
	}()
	select {
	case <-ch1:
		fmt.Println("time now stop")
	}
}
