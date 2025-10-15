package expressions

import (
	"fmt"
)

func SolveUnit(c Context) (Context, error) {
	patterns := []Context{
		{to: "I", value: 1},
		{to: "II", value: 2},
		{to: "III", value: 3},
		{to: "IV", value: 4},
		{to: "V", value: 5},
		{to: "VI", value: 6},
		{to: "VII", value: 7},
		{to: "VIII", value: 8},
		{to: "IX", value: 9},
	}

	result, err := solveExpressionByTenth(patterns, c, SolveZero)
	if err != nil {
		return Context{}, fmt.Errorf("cannot solve units for '%d': '%w'", c.value, err)
	}

	return result, nil
}
