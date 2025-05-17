package repository

import "golang_katas/bank/domain"

type AccountRepository interface {
	Deposit(amount int)
	Withdraw(amount int)
	GetAll() []domain.Transaction
}

type PersonalAccountRepository struct {
	transactions []domain.Transaction
}

func (acc *PersonalAccountRepository) Deposit(amount int) {
	acc.transactions = append(acc.transactions, domain.Transaction{Amount: amount})
}

func (acc *PersonalAccountRepository) Withdraw(amount int) {
	acc.transactions = append(acc.transactions, domain.Transaction{Amount: -amount})
}

func (acc *PersonalAccountRepository) GetAll() []domain.Transaction {
	return acc.transactions
}
