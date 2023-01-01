package main

import "fmt"

type Account struct {
	AccountNo string
	Pwd       string
	Balance   float64
}

// deposit
func (account *Account) Deposit(money float64, pwd string) {
	if pwd != account.Pwd {
		fmt.Println("pwd error")
		return
	}

	if money <= 0 {
		fmt.Println("money error")
		return
	}

	account.Balance += money
	fmt.Println("Deposit successfully")
}

// withdraw
func (account *Account) Withdraw(money float64, pwd string) {
	if pwd != account.Pwd {
		fmt.Println("pwd error")
		return
	}

	if money <= 0 || money > account.Balance {
		fmt.Println("money error")
		return
	}

	account.Balance -= money
	fmt.Println("Withdraw successfully")
}

// check balance
func (account *Account) Query(pwd string) {
	if pwd != account.Pwd {
		fmt.Println("pwd error")
		return
	}
	fmt.Printf("Account=%v Balance=%v \n", account.AccountNo, account.Balance)
}
func main() {
	account := Account{
		AccountNo: "abc11111",
		Pwd:       "1111",
		Balance:   100.0,
	}
	account.Query("1111")
	account.Deposit(200, "1111")
	account.Query("1111")
	account.Withdraw(150, "1111")
	account.Query("1111")
}
