package bank

import (
	"errors"
	"fmt"
)

type Customer struct {
	Name    string
	Address string
	Phone   string
}

type Account struct {
	Customer
	Number  int32
	Balance float64
}

type Bank interface {
	Statement() string
}

func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("the amount to deposit should be greater than zero")
	}

	a.Balance += amount
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("the amount to withdraw must be greater than zero")
	}

	if a.Balance < amount {
		return errors.New("the amount to withdraw must be less than account's balance")
	}

	a.Balance -= amount
	return nil
}

func (a *Account) Transfer(amount float64, dest *Account) error {
	if amount <= 0 {
		return errors.New("the amount to withdraw must be greater than zero")
	}

	if a.Balance < amount {
		return errors.New("the amount to withdraw must be less than account's balance")
	}

	a.Withdraw(amount)
	dest.Deposit(amount)
	return nil
}

func (a *Account) Statement() string {
	return fmt.Sprintf("Name: %v, Address: %v, Phone: %v, Number: %v, Balance: %v", a.Name, a.Address, a.Phone, a.Number, a.Balance)
}

func Statement(b Bank) string {
	return b.Statement()
}
