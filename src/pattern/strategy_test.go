package pattern

import "testing"

func TestExamplePayByCash(t *testing.T) {
	payment := NewPayment("Ada","", 123, &Cash{})
	payment.Pay()

}

func TestExamplePayByBank(t *testing.T) {
	payment := NewPayment("Bob", "00002", 888, &Bank{})
	payment.Pay()
}