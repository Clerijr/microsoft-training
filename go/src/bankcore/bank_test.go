package bank

import "testing"

func TestAccount(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John Doe",
			Address: "123 Main St",
			Phone:   "555-555-5555",
		},
		Number:  1001,
		Balance: 0,
	}

	if account.Name == "" {
		t.Error("can't create an account object without a name")
	}
}
