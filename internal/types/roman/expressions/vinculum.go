package expressions

import (
	"fmt"
	"strings"

	"taurino.com/numerals/tools"
)

type VinculumExpression struct{}

func (exp VinculumExpression) Solve(c Context) (Context, error) {
	if !strings.HasPrefix("(", c.from) {
		return delegateToPlain(c)
	}

	closingIndex := strings.Index(c.from, ")")
	if closingIndex == -1 {
		return Context{}, fmt.Errorf("malformed vinculum in: '%s'", c.from)
	}

	vinculumI := Context{from: c.from[1:closingIndex]}
	vinculumO, e := VinculumExpression{}.Solve(vinculumI)
	if e != nil {
		return Context{}, tools.NewWrappedError(vinculumI.from, VinculumExpression{}, VinculumExpression{}, e)
	}

	plainI := Context{from: c.from[(closingIndex + 1):]}
	plainO, e := PlainExpression{}.Solve(plainI)
	if e != nil {
		return Context{}, tools.NewWrappedError(plainI.from, VinculumExpression{}, PlainExpression{}, e)
	}

	valueLeft, e := tools.MultiplySafe(vinculumO.to, 1000)
	if e != nil {
		return Context{}, fmt.Errorf("value to big: '%w'", e)
	}
	valueWhole, e := tools.AddSafe(valueLeft, plainO.to)
	if e != nil {
		return Context{}, fmt.Errorf("value to big: '%w'", e)
	}

	return Context{
		from: c.from,
		to:   valueWhole,
	}, nil
}

func delegateToPlain(c Context) (Context, error) {
	r, e := PlainExpression{}.Solve(c)
	if e != nil {
		return Context{}, tools.NewWrappedError(c.from, VinculumExpression{}, PlainExpression{}, e)
	}

	return r, nil
}
