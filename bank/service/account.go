package service

import (
	"golang_katas/bank/repository"
	"golang_katas/bank/ui"
)

type Account interface {
	Deposit(amount int)
	Withdraw(amount int)
	PrintStatement() string
}

type PersonalAccount struct {
	Repository repository.AccountRepository
	Printer    ui.AccountPrinter
}

func (acc *PersonalAccount) Deposit(amount int) {
	acc.Repository.Deposit(amount)
}

func (acc *PersonalAccount) Withdraw(amount int) {
	acc.Repository.Withdraw(amount)
}

func (acc *PersonalAccount) PrintStatement() string {
	return acc.Printer.PrintStatement()
}
