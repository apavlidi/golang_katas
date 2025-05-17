package ui

type AccountPrinter interface {
	PrintStatement() string
}

type PersonalAccountPrinter struct {
}

func (p *PersonalAccountPrinter) PrintStatement() string {
	panic("Not implemented")
}
