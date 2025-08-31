package expressions

import (
	"fmt"
)

// ZeroExpression is the only terminal expression and has the responsibility to check if the string is correct
type ZeroExpression struct{}

func (exp ZeroExpression) Solve(c Context) (Context, error) {
	// if not allowed characters are present, this expression will receive them
	if len(c.from) != 0 {
		return Context{}, fmt.Errorf("not allowed expression for '%s'", c.from)
	}

	return Context{from: "", to: 0}, nil
}
