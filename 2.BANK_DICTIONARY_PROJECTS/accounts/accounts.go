package accounts

import "errors"

// Account struct
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Can't withdraw")

// constructor 을 다음과같이 만듬
//NewAccount Creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amout on yout account
func (account *Account) Deposit(amount int) {
	account.balance += amount
}

// Balance
func (account Account) Balance() int {
	return account.balance
}

// Withdraw   -> error exceoption
func (account *Account) Withdraw(amount int) error {
	if account.balance < amount {
		// return errors.New("Can't widthdraw amount is more than yout balance")
		return errNoMoney
	}
	account.balance -= amount
	return nil
	// nill is null or None

}
