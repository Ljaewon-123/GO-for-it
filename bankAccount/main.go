package main

import (
	"fmt"

	"github.com/Ljaewon-123/GO-for-it/bankAccount/accounts"
)

func main() {
	account := accounts.NewAccount("jaewon")
	account.Deposit(10)
	fmt.Println(account.Balance())
}
