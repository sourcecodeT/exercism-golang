package account

import "sync"

func Open(amt int) *BankAccount {
	if amt < 0 {
		return nil
	}
	return &BankAccount{balance: amt, open: true}
}

type BankAccount struct {
	mux     sync.Mutex
	balance int
	open    bool
}

func (a *BankAccount) Close() (int, bool) {
	a.mux.Lock()
	defer a.mux.Unlock()
	if a.open {
		a.open = false
		return a.balance, true
	}
	return 0, false
}

func (a *BankAccount) Balance() (int, bool) {
	if a.open {
		return a.balance, true
	}
	return 0, false
}

func (a *BankAccount) Deposit(value int) (int, bool) {
	a.mux.Lock()
	defer a.mux.Unlock()

	if value+a.balance < 0 {
		return 0, false
	}

	if a.open {
		a.balance = a.balance + value
		return a.balance, true
	}

	return 0, false
}
