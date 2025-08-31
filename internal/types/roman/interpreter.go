package roman

import (
	"fmt"

	"taurino.com/numerals/internal/types/roman/expressions"
	"taurino.com/numerals/tools"
)

type Interpreter struct {
}

func (i Interpreter) Interpret(input string) (tools.Context, error) {
	c := expressions.NexContextFromString(input)
	r, e := expressions.VinculumExpression{}.Solve(c)
	if e != nil {
		return nil, fmt.Errorf("cannot interpret '%s': '%w", input, e)
	}

	return r, nil
}
