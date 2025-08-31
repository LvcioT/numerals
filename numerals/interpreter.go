package numerals

import "strings"

// Expression is the interface for our interpreter.
type Expression interface {
	Interpret(context *Context)
}

// TerminalExpression implements the Expression interface for a terminal symbol.
type TerminalExpression struct {
	Roman  string
	Arabic int
}

// Interpret checks if the context's input starts with the Roman numeral.
// If it does, it adds the Arabic value to the output and removes the
// Roman numeral from the input.
func (te *TerminalExpression) Interpret(context *Context) {
	if strings.HasPrefix(context.Input, te.Roman) {
		context.Output += te.Arabic
		context.Input = strings.TrimPrefix(context.Input, te.Roman)
	}
}

// VinculumExpression implements the Expression interface for a vinculum symbol.
type VinculumExpression struct {
	Roman  string
	Arabic int
}

// Interpret checks if the context's input starts with the Roman numeral in parenthesis.
// If it does, it adds the Arabic value multiplied by 1000 to the output and removes the
// Roman numeral from the input.
func (ve *VinculumExpression) Interpret(context *Context) {
	vinculumRoman := "(" + ve.Roman + ")"
	if strings.HasPrefix(context.Input, vinculumRoman) {
		context.Output += ve.Arabic * 1000
		context.Input = strings.TrimPrefix(context.Input, vinculumRoman)
	}
}

func Interpret(roman string) int {
	roman = strings.ToUpper(roman)
	context := &Context{Input: roman}

	expressions := []Expression{
		&VinculumExpression{Roman: "M", Arabic: 1000},
		&VinculumExpression{Roman: "CM", Arabic: 900},
		&VinculumExpression{Roman: "D", Arabic: 500},
		&VinculumExpression{Roman: "CD", Arabic: 400},
		&VinculumExpression{Roman: "C", Arabic: 100},
		&VinculumExpression{Roman: "XC", Arabic: 90},
		&VinculumExpression{Roman: "L", Arabic: 50},
		&VinculumExpression{Roman: "XL", Arabic: 40},
		&VinculumExpression{Roman: "X", Arabic: 10},
		&VinculumExpression{Roman: "IX", Arabic: 9},
		&VinculumExpression{Roman: "V", Arabic: 5},
		&VinculumExpression{Roman: "IV", Arabic: 4},
		&VinculumExpression{Roman: "I", Arabic: 1},
		&TerminalExpression{Roman: "M", Arabic: 1000},
		&TerminalExpression{Roman: "CM", Arabic: 900},
		&TerminalExpression{Roman: "D", Arabic: 500},
		&TerminalExpression{Roman: "CD", Arabic: 400},
		&TerminalExpression{Roman: "C", Arabic: 100},
		&TerminalExpression{Roman: "XC", Arabic: 90},
		&TerminalExpression{Roman: "L", Arabic: 50},
		&TerminalExpression{Roman: "XL", Arabic: 40},
		&TerminalExpression{Roman: "X", Arabic: 10},
		&TerminalExpression{Roman: "IX", Arabic: 9},
		&TerminalExpression{Roman: "V", Arabic: 5},
		&TerminalExpression{Roman: "IV", Arabic: 4},
		&TerminalExpression{Roman: "I", Arabic: 1},
	}

	for _, expression := range expressions {
		for {
			tempContext := &Context{Input: context.Input}
			expression.Interpret(tempContext)
			if tempContext.Input != context.Input {
				context.Input = tempContext.Input
				context.Output += tempContext.Output
			} else {
				break
			}
		}
	}

	if context.Input != "" {
		return 0 // Invalid characters remaining
	}

	return context.Output
}
