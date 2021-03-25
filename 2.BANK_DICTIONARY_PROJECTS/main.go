package main

import (
	"fmt"

	"github.com/Donghee-L/BANK_DICTIONARY_PROJECTS/accounts"
)

func main() {
	account := accounts.NewAccount("dongs")
	fmt.Println(*account)
	account.Deposit((10))
	fmt.Println(account.Balance())
	fmt.Println(account)
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	}

}
