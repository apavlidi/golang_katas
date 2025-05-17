package bank

import (
	"awesomeProject2/bank/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccount(t *testing.T) {

	var acc service.Account = service.PersonalAccount{}

	acc.Deposit(1000)
	acc.Deposit(2000)
	acc.Withdraw(500)
	statement := acc.PrintStatement()
	assert.Equal(t, "Date       || Amount || Balance\n"+
		"14/01/2012 || -500   || 2500\n"+
		"13/01/2012 || 2000   || 3000\n"+
		"10/01/2012 || 1000   || 1000", statement)
}
