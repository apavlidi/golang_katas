package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestServiceDeposit(t *testing.T) {
	mockRepo := new(MockPersonalAccountRepository)
	amount := 1000
	mockRepo.On("Deposit", amount)

	var acc Account = PersonalAccount{repository: mockRepo}
	acc.Deposit(amount)
	mockRepo.AssertCalled(t, "Deposit", amount)
}

func TestServiceWithdraw(t *testing.T) {
	mockRepo := new(MockPersonalAccountRepository)
	amount := 500
	mockRepo.On("Withdraw", amount)

	acc := PersonalAccount{repository: mockRepo}
	acc.Withdraw(amount)
	mockRepo.AssertCalled(t, "Withdraw", amount)
}

func TestPrintStatement(t *testing.T) {
	mockRepo := new(MockPersonalAccountRepository)
	mockPrinter := new(MockPersonalAccountPrinter)
	statement := "14/01/2012 || -500   || 2500"
	mockPrinter.On("PrintStatement").Return(statement)

	acc := PersonalAccount{repository: mockRepo, printer: mockPrinter}
	assert.Equal(t, statement, acc.PrintStatement())
}

type MockPersonalAccountPrinter struct {
	mock.Mock
}

func (m *MockPersonalAccountPrinter) PrintStatement() string {
	args := m.Called()
	return args.String(0)
}

type MockPersonalAccountRepository struct {
	mock.Mock
}

func (m *MockPersonalAccountRepository) Deposit(amount int) {
	m.Called(amount)
}

func (m *MockPersonalAccountRepository) Withdraw(amount int) {
	m.Called(amount)
}
