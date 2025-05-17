package ui

import (
	"github.com/stretchr/testify/assert"
	"golang_katas/bank/domain"
	"testing"
	"time"
)

func TestPersonalAccountPrinter_PrintEmptyStatement(t *testing.T) {
	printer := PersonalAccountPrinter{}

	statement := printer.PrintStatement([]domain.Transaction{})
	assert.Equal(t, statement, "Date       || Amount || Balance\n")
}

func TestPersonalAccountPrinter_PrintOneTransactionStatement(t *testing.T) {
	printer := PersonalAccountPrinter{}

	statement := printer.PrintStatement([]domain.Transaction{
		{
			Amount: 1000,
			Date:   time.Date(2012, time.January, 14, 0, 0, 0, 0, time.UTC),
		},
	})
	assert.Equal(t, "Date       || Amount || Balance\n"+
		"14/01/2012 || 1000  ||  1000\n", statement)
}

func TestPersonalAccountPrinter_PrintTwoTransactionStatement(t *testing.T) {
	printer := PersonalAccountPrinter{}

	statement := printer.PrintStatement([]domain.Transaction{
		{
			Amount: 1000,
			Date:   time.Date(2012, time.January, 10, 0, 0, 0, 0, time.UTC),
		},
		{
			Amount: 2000,
			Date:   time.Date(2012, time.January, 13, 0, 0, 0, 0, time.UTC),
		},
	})
	assert.Equal(t, "Date       || Amount || Balance\n"+
		"13/01/2012 || 2000  ||  3000\n"+
		"10/01/2012 || 1000  ||  1000\n", statement)
}

func TestPersonalAccountPrinter_PrintThreeTransactionStatement(t *testing.T) {
	printer := PersonalAccountPrinter{}

	statement := printer.PrintStatement([]domain.Transaction{
		{
			Amount: 1000,
			Date:   time.Date(2012, time.January, 10, 0, 0, 0, 0, time.UTC),
		},
		{
			Amount: 2000,
			Date:   time.Date(2012, time.January, 13, 0, 0, 0, 0, time.UTC),
		},
		{
			Amount: -500,
			Date:   time.Date(2012, time.January, 14, 0, 0, 0, 0, time.UTC),
		},
	})
	assert.Equal(t, "Date       || Amount || Balance\n"+
		"14/01/2012 || -500  ||  2500\n"+
		"13/01/2012 || 2000  ||  3000\n"+
		"10/01/2012 || 1000  ||  1000\n", statement)
}
