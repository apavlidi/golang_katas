package repository

import (
	"golang_katas/bank/domain"
	"golang_katas/bank/utils"
)

type AccountRepository interface {
	Deposit(amount int)
	Withdraw(amount int)
	GetAll() []domain.Transaction
}

type PersonalAccountRepository struct {
	transactions []domain.Transaction
	Clock        utils.Clock
}

func (acc *PersonalAccountRepository) Deposit(amount int) {
	acc.transactions = append(acc.transactions, domain.Transaction{Amount: amount, Date: acc.Clock.Now()})
}

func (acc *PersonalAccountRepository) Withdraw(amount int) {
	acc.transactions = append(acc.transactions, domain.Transaction{Amount: -amount, Date: acc.Clock.Now()})
}

func (acc *PersonalAccountRepository) GetAll() []domain.Transaction {
	return acc.transactions
}
