package main

import (
	"errors"
	"fmt"
)

// ErrInsufficientFunds describes the error value for over withdrawal
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Bitcoin represent digital currency type
type Bitcoin int

// Wallet represents the data structure for crypto wallets
type Wallet struct {
	balance Bitcoin
}

// Stringer will define how we want to print Bitcoin related info
type Stringer interface {
	String() string
}

// String func returns the info of the wallet in a specific format
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Deposit func takes in amount of type int and adds it to the Wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Withdraw takes in a Bitcoin val and deducts it from the wallet balance
func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

// Balance func returns the balance in the Bitcoin
func (w *Wallet) Balance() (balance Bitcoin) {
	return w.balance
}
