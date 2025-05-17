package repository

type AccountRepository interface {
	Deposit(amount int)
	Withdraw(amount int)
}

type PersonalAccountRepository struct {
}

func (acc PersonalAccountRepository) Deposit(amount int) {
	panic("Not implemented")
}

func (acc PersonalAccountRepository) Withdraw(amount int) {
	panic("Not implemented")
}
