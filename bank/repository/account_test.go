package repository

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestPersonalAccountRepository_Deposit(t *testing.T) {
	mockClock := new(MockClock)
	repo := PersonalAccountRepository{Clock: mockClock}
	date := time.Date(1996, time.October, 9, 0, 0, 0, 0, time.UTC)
	mockClock.On("Now").Return(date)

	repo.Deposit(1000)
	transactions := repo.GetAll()
	assert.Equal(t, len(transactions), 1)
	assert.Equal(t, transactions[0].Amount, 1000)
	assert.Equal(t, transactions[0].Date, date)
}

func TestPersonalAccountRepository_Withdraw(t *testing.T) {
	mockClock := new(MockClock)
	repo := PersonalAccountRepository{Clock: mockClock}
	date := time.Date(1996, time.October, 10, 0, 0, 0, 0, time.UTC)
	mockClock.On("Now").Return(date)

	repo.Withdraw(1000)
	transactions := repo.GetAll()
	assert.Equal(t, len(transactions), 1)
	assert.Equal(t, transactions[0].Amount, -1000)
	assert.Equal(t, transactions[0].Date, date)
}

type MockClock struct {
	mock.Mock
}

func (m *MockClock) Now() time.Time {
	args := m.Called()
	return args.Get(0).(time.Time)
}
