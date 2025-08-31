package tools

type NumeralType int

// supported numerals
const (
	Arabic NumeralType = iota
	Roman
)

type Context interface {
	GetFrom() (string, error)
	GetValue() (uint64, error)
	GetTo() (string, error)
}

type Interpreter interface {
	Interpret(input string) (Context, error)
}
