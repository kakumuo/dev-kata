package banking

import (
	"fmt"
	"slices"
	"strings"
	"time"
)

type TransactionType int

const (
	TT_Deposit TransactionType = iota
	TT_Withdraw
)

type Transaction struct {
	Amount    int
	Balance   int
	TransType TransactionType
	Date      time.Time
}

type Account struct {
	Balance      int
	Transactions []Transaction
}

func NewAccount() *Account {
	return &Account{
		Balance:      0,
		Transactions: make([]Transaction, 0),
	}
}

func (acc *Account) Deposit(amount int) {
	if amount <= 0 {
		fmt.Printf("Cannot deposit non-positve value: %d\n", amount)
		return
	}

	acc.Balance += amount
	trans := Transaction{
		Balance: acc.Balance, Amount: amount, TransType: TT_Deposit,
		Date: time.Now(),
	}
	acc.Transactions = append(acc.Transactions, trans)

	time.Sleep(time.Second / 4)
}

func (acc *Account) Withdraw(amount int) {
	if amount <= 0 {
		fmt.Printf("Cannot withdraw non-positive value: %d\n", amount)
		return
	} else if acc.Balance < 0 {
		fmt.Printf("Cannot withdraw from overdrafted account: (balance = %d)\n", acc.Balance)
		return
	}

	acc.Balance -= amount
	trans := Transaction{
		Balance: acc.Balance, Amount: amount, TransType: TT_Withdraw,
		Date: time.Now(),
	}
	acc.Transactions = append(acc.Transactions, trans)

	time.Sleep(time.Second / 4)
}

func (acc *Account) String() string {
	const GAP = 4

	tmpTrans := acc.Transactions
	slices.SortFunc(tmpTrans, func(a, b Transaction) int {
		return -1 * a.Date.Compare(b.Date)
	})

	table := make([][]string, 0)
	table = append(table, []string{"Date", "Amount", "Balance"})
	numCols := len(table[0])
	widths := make([]int, numCols)

	for i := 0; i < numCols; i++ {
		widths[i] = len(table[0][i])
	}

	var output strings.Builder = strings.Builder{}

	for _, trans := range tmpTrans {
		row := make([]string, numCols)

		row[0] = trans.Date.Format(time.RFC3339)

		if trans.TransType == TT_Deposit {
			row[1] = fmt.Sprintf("+%d", trans.Amount)
		} else {
			row[1] = fmt.Sprintf("-%d", trans.Amount)
		}

		row[2] = fmt.Sprintf("%d", trans.Balance)

		for i := 0; i < numCols; i++ {
			widths[i] = max(widths[i], len(row[i]))
		}

		table = append(table, row)
	}

	for _, row := range table {
		for i, col := range row {
			output.WriteString(col)
			output.WriteString(strings.Repeat(" ", widths[i]-len(col)+GAP))
		}
		output.WriteByte('\n')
	}

	return output.String()
}
