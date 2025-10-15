package expressions

import (
	"fmt"
)

func SolveThousands(c Context) (Context, error) {
	patterns := []Context{
		{from: "M", value: 1000},
		{from: "MM", value: 2000},
		{from: "MMM", value: 3000},
	}

	result, err := solveExpression(patterns, c, SolveHundreds)
	if err != nil {
		return Context{}, fmt.Errorf("cannot solve thousands for '%s': '%w'", c.from, err)
	}

	return result, nil
}
