package main

import (
	"banking/banking"
	"fmt"
)

func main() {
	acc := banking.NewAccount()
	acc.Deposit(200)
	acc.Deposit(100)
	acc.Withdraw(5)
	acc.Withdraw(3)

	fmt.Println(acc)
}
