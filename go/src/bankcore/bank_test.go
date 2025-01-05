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

func TestDeposit(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John Doe",
			Address: "123 Main St",
			Phone:   "555-555-5555",
		},
		Number:  1001,
		Balance: 0,
	}

	account.Deposit(10)

	if account.Balance != 10 {
		t.Error("balance is not being updated after a deposit")
	}
}

func TestDepositInvalidAmmount(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John Doe",
			Address: "123 Main St",
			Phone:   "555-555-5555",
		},
		Number:  1001,
		Balance: 0,
	}

	if err := account.Deposit(-10); err == nil {
		t.Error("only positive numbers should be allowed to deposit")
	}
}

func TestWithdraw(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John Doe",
			Address: "123 Main St",
			Phone:   "555-555-5555",
		},
		Number:  1001,
		Balance: 0,
	}

	account.Deposit(10)
	account.Withdraw(10)

	if account.Balance != 0 {
		t.Error("balance is not being updated after a withdraw")
	}
}

func TestStatement(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John",
			Address: "123 Main St",
			Phone:   "555-555-5555",
		},
		Number:  1001,
		Balance: 0,
	}

	account.Deposit(100)
	statement := account.Statement()
	if statement != "1001 - John - 100" {
		t.Error("statement doesn't have the proper format")
	}
}
