package main

import (
	"fmt"
	"log"

	"github.com/Ljaewon-123/GO-for-it/bankAccount/accounts"
)

func main() {
	account := accounts.NewAccount("jaewon")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(account.Balance())
}
