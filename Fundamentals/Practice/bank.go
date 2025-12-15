package main

import (
	"errors"
	"fmt"
	"sync"
)

// ------------- strukturalar ------------- //

type User struct {
	Name string
	Age  int
}

type Account struct {
	Owner   *User
	Balance int
	mu      sync.Mutex // race conditionga qarshi
}

type Transaction struct {
	Account *Account
	Amount  int
	Type    string
}

// ------------- Interface ------------- //

type Notifier interface {
	Notify(message string)
}
type EmailNotifier struct{}

func (e EmailNotifier) Notify(message string) {
	fmt.Println("[EMAIL]", message)
}

// ------------- Xatoliklar ------------- //

var ErrMoneyNotEnought = errors.New("Kechirasiz, mablag' yetartli emas")
var ErrInvalidAmount = errors.New("Noto'g'ri summa")

// ------------- Funksiyalar ------------- //

func (a *Account) Deposit(amount int) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}
	a.mu.Lock()
	defer a.mu.Unlock()
	a.Balance += amount

	return nil
}

func (a *Account) Withdraw(amount int) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	if amount > a.Balance {
		return ErrMoneyNotEnought
	}

	a.Balance -= amount
	return nil
}

func TransactionWorker(id int, txChan <-chan Transaction, wg *sync.WaitGroup) {
	// email := EmailNotifier()

	defer wg.Done()

	for tx := range txChan {
		var err error
		if tx.Type == "deposit" {
			err = tx.Account.Deposit(tx.Amount)
		} else if tx.Type == "withdraw" {
			err = tx.Account.Withdraw(tx.Amount)
		} else {
			fmt.Println("--[X]-- | No'malum tranzaksiya turi!")
			continue
		}
		if err != nil {
			fmt.Println("Xatolik: ", err)
		} else {
			msg := fmt.Sprintf("[Worker %d] OK | %s | Amount: %d | Balance: %d | User: %s\n",
				id,
				tx.Type,
				tx.Amount,
				tx.Account.Balance,
				tx.Account.Owner.Name,
			)
			fmt.Print(msg)
			// SendNotification(email, msg)
		}
	}
}

func SendNotification(n Notifier, msg string) {
	n.Notify(msg)
}

func main() {
	user := &User{Name: "Mexroj", Age: 17}
	acc := &Account{
		Owner:   user,
		Balance: 100,
	}

	txChan := make(chan Transaction)
	var wg sync.WaitGroup

	workerCount := 2
	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go TransactionWorker(i, txChan, &wg)
	}

	txChan <- Transaction{Account: acc, Amount: 50, Type: "deposit"}
	txChan <- Transaction{Account: acc, Amount: 30, Type: "withdraw"}
	txChan <- Transaction{Account: acc, Amount: 500, Type: "withdraw"}
	txChan <- Transaction{Account: acc, Amount: -10, Type: "deposit"}
	close(txChan)

	wg.Wait()
}
