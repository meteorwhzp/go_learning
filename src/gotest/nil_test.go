package gotest

import (
	"math/rand"
	"testing"
	"time"
)


func GetRandFloat() float64 {
	rand.Seed(time.Now().UnixNano())
	r := float64(rand.Intn(100)) / float64(100)
	return r
}

func TestFloat(t *testing.T) {
	randVal := GetRandFloat()
	t.Log(randVal)
	if randVal <= 0.9 {
		t.Log("yes")
	}
}