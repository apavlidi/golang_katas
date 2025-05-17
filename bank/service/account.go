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
	repository repository.AccountRepository
	printer    ui.AccountPrinter
}

func (acc PersonalAccount) Deposit(amount int) {
	acc.repository.Deposit(amount)
}

func (acc PersonalAccount) Withdraw(amount int) {
	acc.repository.Withdraw(amount)
}

func (acc PersonalAccount) PrintStatement() string {
	return acc.printer.PrintStatement()
}
