package expressions

import (
	"fmt"
)

func solveExpressionByTenth(patterns []Context, c Context, solveSub ExpressionSolve) (Context, error) {
	cLeft, cRight := c.splitByPrefix(patterns)

	resultRight, err := solveSub(cRight)
	if err != nil {
		return Context{}, fmt.Errorf("cannot solve '%d': '%w'", cRight.value, err)
	}

	result, err := NewContextFromAdd(cLeft, resultRight)
	if err != nil {
		return Context{}, fmt.Errorf("cannot concatenate '%d+%d': '%w'", cLeft.value, resultRight.value, err)
	}

	return result, nil
}
