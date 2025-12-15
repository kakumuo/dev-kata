package banking_test

import (
	"banking/banking"
	"testing"
)

/*
- user can deposit into account
- user can widthdraw from account
- user can overdraft, but once overdrafted, user cannot withdraw
- user cannot depost nothing into account
- user cannot withdraw nothing from account
- all withdraws and deposits are recorded with a date
- account transaction history is printed in tabular format
*/

func TestDeposit(t *testing.T) {
	acc := banking.NewAccount()
	acc.Deposit(100)

	if acc.Balance != 100 {
		t.Error("balance should be 100")
	} else if len(acc.Transactions) != 1 {
		t.Error("Transaction should be recorded")
	}

	acc.Deposit(-100)

	if acc.Balance != 100 {
		t.Errorf("balance should not be reduced, balance :%d", acc.Balance)
	} else if len(acc.Transactions) != 1 {
		t.Error("Transaction should not be recorded")
	}

	acc.Deposit(0)

	if len(acc.Transactions) != 1 {
		t.Error("Transaction should not be recorded")
	}

	acc.Deposit(100)

	if acc.Balance != 200 {
		t.Error("Balance should be increased to 200")
	} else if len(acc.Transactions) != 2 {
		t.Error("Transaction should be recorded")
	}
}

func TestWidthdraw(t *testing.T) {
	acc := banking.NewAccount()
	acc.Balance = 100

	acc.Withdraw(99)
	if acc.Balance != 1 {
		t.Errorf("balance != 1; balance: %d", acc.Balance)
	} else if len(acc.Transactions) != 1 {
		t.Error("Transaction was not recorded")
	}

	acc.Withdraw(-100)
	if acc.Balance != 1 {
		t.Errorf("balance != 1; balance: %d", acc.Balance)
	} else if len(acc.Transactions) != 1 {
		t.Error("Transaction should have not been recorded")
	}

	acc.Withdraw(0)
	if acc.Balance != 1 {
		t.Errorf("balance != 1; balance: %d", acc.Balance)
	} else if len(acc.Transactions) != 1 {
		t.Error("Transaction should have not been recorded")
	}

	acc.Withdraw(100)
	if acc.Balance != -99 {
		t.Errorf("balance != -99; balance: %d", acc.Balance)
	} else if len(acc.Transactions) != 2 {
		t.Error("Transaction should have been recorded")
	}

	acc.Withdraw(1)
	if acc.Balance != -99 {
		t.Errorf("balance != -99; balance: %d", acc.Balance)
	} else if len(acc.Transactions) != 2 {
		t.Error("Transaction should have not been recorded")
	}
}
