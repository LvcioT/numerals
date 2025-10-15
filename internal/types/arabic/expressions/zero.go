package expressions

import (
	"fmt"
)

// SolveZero is the only terminale expression, validates if the given Context has an empty 'from' string and assigns a value of 0 to the returned Context.
// Returns an error if 'from' is not empty.
func SolveZero(c Context) (Context, error) {
	// if not allowed characters are present, this expression will receive them
	if c.value != 0 {
		return Context{}, fmt.Errorf("not allowed expression for '%d'", c.value)
	}

	return Context{to: "", value: 0}, nil
}
