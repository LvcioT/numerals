package expressions

import (
	"fmt"
)

func SolveThousandLess(c Context) (Context, error) {
	patterns := []Context{
		{to: "M", value: 1_000},
		{to: "MM", value: 2_000},
		{to: "MMM", value: 3_000},
	}

	result, err := solveExpressionByTenth(patterns, c, SolveHundreds)
	if err != nil {
		return Context{}, fmt.Errorf("cannot solve thousands for '%d': '%w'", c.value, err)
	}

	return result, nil
}
