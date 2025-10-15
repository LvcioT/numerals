package arabic

import (
	"fmt"
	"strconv"

	"taurino.com/numerals/internal/types/arabic/expressions"
	"taurino.com/numerals/tools"
)

type Interpreter struct {
}

func (i Interpreter) Interpret(input string) (tools.Context, error) {
	value, err := strconv.ParseUint(input, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("cannot interpret '%s' as unsigned integer: '%w", input, err)
	}

	c := expressions.NewContextFromValue(value)
	r, e := expressions.SolveThousandMore(c)
	if e != nil {
		return nil, fmt.Errorf("cannot interpret '%s': '%w", input, e)
	}

	return r, nil
}
