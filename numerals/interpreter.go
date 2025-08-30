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
