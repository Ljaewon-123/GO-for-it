package accounts

import "errors"

// Account struct
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("can't withdraw")

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// 호출한 object사용
// Deposit x amount on you account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// return balance
func (a Account) Balance() int {
	return a.balance
}

// Withdraw x amout from you account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}

	a.balance -= amount
	return nil
}

// ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// get owner
func (a Account) Owner() string {
	return a.owner
}

// struct에서 Go가 자동으로 호출하는 매서드중 하나
func (a Account) String() string {
	return "whatever you want"
}
