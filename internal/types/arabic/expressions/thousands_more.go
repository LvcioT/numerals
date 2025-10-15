package expressions

import (
	"fmt"
)

func SolveThousandMore(c Context) (Context, error) {
	// no vinculum needed
	if c.value < 4_000 {
		return SolveThousandLess(c)
	}

	// vinculum is needed
	cThousand := Context{
		value: c.value / 1_000,
	}
	cPlain := Context{
		value: c.value % 1_000,
	}

	resultThousand, err := SolveThousandMore(cThousand)
	if err != nil {
		return Context{}, fmt.Errorf("cannot solve '%d': '%w'", cThousand.value, err)
	}
	resultPlain, err := SolveHundreds(cPlain)
	if err != nil {
		return Context{}, fmt.Errorf("cannot solve '%d': '%w'", cPlain.value, err)
	}

	result, err := NewContextFromThousandMultiply(resultThousand, resultPlain)
	if err != nil {
		return Context{}, fmt.Errorf("cannot concatenate with vinculum '%s+%s': '%w'", cThousand.to, resultPlain.to, err)
	}

	return result, nil
}
