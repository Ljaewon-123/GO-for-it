package accounts

// Account struct
type Account struct {
	owner   string
	balance int
}

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

func (a Account) Balance() int {
	return a.balance
}
