package repository

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPersonalAccountRepository_Deposit(t *testing.T) {
	repo := PersonalAccountRepository{}

	repo.Deposit(1000)
	transactions := repo.GetAll()
	assert.Equal(t, len(transactions), 1)
	assert.Equal(t, transactions[0].Amount, 1000)
}

func TestPersonalAccountRepository_Withdraw(t *testing.T) {
	repo := PersonalAccountRepository{}

	repo.Withdraw(1000)
	transactions := repo.GetAll()
	assert.Equal(t, len(transactions), 1)
	assert.Equal(t, transactions[0].Amount, -1000)
}
