package ui

import (
	"fmt"
	"golang_katas/bank/domain"
	"slices"
	"strings"
)

const DateLayout = "02/01/2006"

type AccountPrinter interface {
	PrintStatement(transactions []domain.Transaction) string
}

type PersonalAccountPrinter struct {
}

func (p *PersonalAccountPrinter) PrintStatement(transactions []domain.Transaction) string {
	header := "Date       || Amount || Balance\n"
	var statementsWithBalance []string
	total := 0
	for _, val := range transactions {
		total += val.Amount
		statementsWithBalance = append(statementsWithBalance, fmt.Sprintf("%s || %d  ||  %d\n", val.Date.Format(DateLayout), val.Amount, total))
	}
	slices.Reverse(statementsWithBalance)
	allStatementsReversed := strings.Join(statementsWithBalance, "")

	return header + allStatementsReversed
}
