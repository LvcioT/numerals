package expressions

import (
	"fmt"
)

func SolveHundreds(c Context) (Context, error) {
	patterns := []Context{
		{from: "C", value: 100},
		{from: "CC", value: 200},
		{from: "CCC", value: 300},
		{from: "CD", value: 400},
		{from: "D", value: 500},
		{from: "DC", value: 600},
		{from: "DCC", value: 700},
		{from: "DCCC", value: 800},
		{from: "CM", value: 900},
	}

	result, err := solveExpression(patterns, c, SolveTens)
	if err != nil {
		return Context{}, fmt.Errorf("cannot solve hundreds for '%s': '%w'", c.from, err)
	}

	return result, nil
}
