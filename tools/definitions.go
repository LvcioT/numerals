package tools

type NumeralType int

// supported numerals
const (
	Arabic NumeralType = iota
	Roman
)

type Context interface {
	GetFrom() string
	GetValue() uint64
	GetTo() string
}

type Interpreter interface {
	Interpret(input string) (Context, error)
}
