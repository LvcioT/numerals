package expressions

import (
	"fmt"
)

func SolveTens(c Context) (Context, error) {
	patterns := []Context{
		{to: "X", value: 10},
		{to: "XX", value: 20},
		{to: "XXX", value: 30},
		{to: "XL", value: 40},
		{to: "L", value: 50},
		{to: "LX", value: 60},
		{to: "LXX", value: 70},
		{to: "LXXX", value: 80},
		{to: "XC", value: 90},
	}

	result, err := solveExpressionByTenth(patterns, c, SolveUnit)
	if err != nil {
		return Context{}, fmt.Errorf("cannot solve tens for '%d': '%w'", c.value, err)
	}

	return result, nil
}
