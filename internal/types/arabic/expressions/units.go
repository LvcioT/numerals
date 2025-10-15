package expressions

import (
	"fmt"
)

func SolveUnit(c Context) (Context, error) {
	patterns := []Context{
		{from: "I", value: 1},
		{from: "II", value: 2},
		{from: "III", value: 3},
		{from: "IV", value: 4},
		{from: "V", value: 5},
		{from: "VI", value: 6},
		{from: "VII", value: 7},
		{from: "VIII", value: 8},
		{from: "IX", value: 9},
	}

	result, err := solveExpression(patterns, c, SolveZero)
	if err != nil {
		return Context{}, fmt.Errorf("cannot solve units for '%s': '%w'", c.from, err)
	}

	return result, nil
}
