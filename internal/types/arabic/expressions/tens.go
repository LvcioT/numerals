package expressions

import (
	"fmt"
)

func SolveTens(c Context) (Context, error) {
	patterns := []Context{
		{from: "X", value: 10},
		{from: "XX", value: 20},
		{from: "XXX", value: 30},
		{from: "XL", value: 40},
		{from: "L", value: 50},
		{from: "LX", value: 60},
		{from: "LXX", value: 70},
		{from: "LXXX", value: 80},
		{from: "XC", value: 90},
	}

	result, err := solveExpression(patterns, c, SolveUnit)
	if err != nil {
		return Context{}, fmt.Errorf("cannot solve tens for '%s': '%w'", c.from, err)
	}

	return result, nil
}
