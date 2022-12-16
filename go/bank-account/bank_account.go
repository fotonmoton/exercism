package account

import "sync"

type Account interface {
	Open(initialDeposit int64) *Account
	Close() (payout int64, ok bool)
	Balance() (balance int64, ok bool)
	Deposit(amount int64) (newBalance int64, ok bool)
}

type account struct {
	open    bool
	balance int64
	sync.Mutex
}

func Open(intialDeposit int64) *account {
	if intialDeposit < 0 {
		return nil
	}

	return &account{balance: intialDeposit, open: true}
}

func (acc *account) Close() (payout int64, ok bool) {
	acc.Lock()
	defer acc.Unlock()

	if !acc.open {
		return 0, false
	}

	acc.open = false

	return acc.balance, true
}

func (acc *account) Balance() (balance int64, ok bool) {
	if !acc.open {
		return 0, false
	}

	return acc.balance, acc.open
}

func (acc *account) Deposit(amount int64) (payout int64, ok bool) {
	acc.Lock()
	defer acc.Unlock()

	if !acc.open {
		return 0, false
	}

	if (acc.balance + amount) < 0 {
		return acc.balance, false

	}

	acc.balance += amount

	return acc.balance, true
}
