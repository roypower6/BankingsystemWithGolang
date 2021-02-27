package main

import (
	"fmt"

	accounts "./banking"
)

func main() {
	account := accounts.NewAccount("roy")
	account.Deposit(10)
	fmt.Println(account)
	// err := account.Withdraw(5)
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(account.Balance(), account.Owner())
}
