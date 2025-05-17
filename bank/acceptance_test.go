package bank

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang_katas/bank/repository"
	"golang_katas/bank/service"
	"golang_katas/bank/ui"
	"testing"
	"time"
)

func TestAccount(t *testing.T) {
	mockClock := new(MockClock)
	mockClock.On("Now").Return(time.Date(2012, time.January, 10, 0, 0, 0, 0, time.UTC)).Once()
	mockClock.On("Now").Return(time.Date(2012, time.January, 13, 0, 0, 0, 0, time.UTC)).Once()
	mockClock.On("Now").Return(time.Date(2012, time.January, 14, 0, 0, 0, 0, time.UTC)).Once()

	var acc service.Account = &service.PersonalAccount{
		Repository: &repository.PersonalAccountRepository{Clock: mockClock},
		Printer:    &ui.PersonalAccountPrinter{},
	}

	acc.Deposit(1000)
	acc.Deposit(2000)
	acc.Withdraw(500)
	statement := acc.PrintStatement()
	assert.Equal(t, "Date       || Amount || Balance\n"+
		"14/01/2012 || -500  ||  2500\n"+
		"13/01/2012 || 2000  ||  3000\n"+
		"10/01/2012 || 1000  ||  1000\n", statement)
}

type MockClock struct {
	mock.Mock
}

func (m *MockClock) Now() time.Time {
	args := m.Called()
	return args.Get(0).(time.Time)
}
