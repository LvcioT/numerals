package expressions

import (
	"fmt"
)

func solveExpression(patterns []Context, c Context, solveSub ExpressionSolve) (Context, error) {
	cLeft, cRight := c.splitByPrefix(patterns)

	resultRight, err := solveSub(cRight)
	if err != nil {
		return Context{}, fmt.Errorf("cannot solve '%s': '%w'", cRight.from, err)
	}

	result, err := NewContextFromConcatenation(cLeft, resultRight)
	if err != nil {
		return Context{}, fmt.Errorf("cannot concatenate '%s'+'%s': '%w'", cLeft.from, resultRight.from, err)
	}

	return result, nil
}
