package gotest

import (
	"math"
	"testing"
)

func TestGetOrderList(t *testing.T) {
	t.Log("Order List is [123, 456, 789]")
}

func TestGetDriverList(t *testing.T) {
	t.Log("Driver List is [123, 456, 789]")
}

func TestAbs(t *testing.T) {
	got := Abs(-1)
	if got != 2 {
		t.Errorf("Abs(-1) = %f; want 1", got)
	}
}

func Abs(x float64) float64 {
	return math.Float64frombits(math.Float64bits(x) &^ (1 << 63))
}