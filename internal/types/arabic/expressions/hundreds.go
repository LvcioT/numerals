package expressions

import (
	"fmt"
)

func SolveHundreds(c Context) (Context, error) {
	patterns := []Context{
		{to: "C", value: 100},
		{to: "CC", value: 200},
		{to: "CCC", value: 300},
		{to: "CD", value: 400},
		{to: "D", value: 500},
		{to: "DC", value: 600},
		{to: "DCC", value: 700},
		{to: "DCCC", value: 800},
		{to: "CM", value: 900},
	}

	result, err := solveExpressionByTenth(patterns, c, SolveTens)
	if err != nil {
		return Context{}, fmt.Errorf("cannot solve hundreds for '%d': '%w'", c.value, err)
	}

	return result, nil
}
