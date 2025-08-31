package expressions

import (
	"errors"
)

type PlainExpression struct{}

func (e PlainExpression) Solve(c Context) (Context, error) {
	return c, errors.New("not yet implemented")
}
