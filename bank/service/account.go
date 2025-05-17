package service

type Account interface {
	Deposit(amount int)
	Withdraw(amount int)
	PrintStatement() string
}

type PersonalAccount struct {
}

func (acc PersonalAccount) Deposit(amount int) {
	panic("Not implemented")
}

func (acc PersonalAccount) Withdraw(amount int) {
	panic("Not implemented")
}

func (acc PersonalAccount) PrintStatement() string {
	panic("Not implemented")
}
