package expressions

import (
	"fmt"
	"strings"

	"taurino.com/numerals/tools"
)

type ThousandsExpression struct{}

func (exp ThousandsExpression) Solve(c Context) (Context, error) {
	// must be ordered by length
	patterns := []Context{
		{from: "MMM", to: 3000},
		{from: "MM", to: 2000},
		{from: "M", to: 1000},
	}

	for _, pattern := range patterns {
		prefix := pattern.from
		value := pattern.to

		if strings.HasPrefix(c.from, prefix) {
			valueLeft := value

			contextI := Context{from: strings.TrimPrefix(c.from, prefix)}
			contextO, e := HundredsExpression{}.Solve(contextI)
			if e != nil {
				return Context{}, tools.NewWrappedError(contextO.from, ThousandsExpression{}, HundredsExpression{}, e)
			}

			valueWhole, e := tools.AddSafe(valueLeft, contextO.to)
			if e != nil {
				return Context{}, fmt.Errorf("value to big: '%w'", e)
			}

			return Context{
				from: c.from,
				to:   valueWhole,
			}, nil
		}
	}

	// no thousands present, check for smaller quantity
	r, e := HundredsExpression{}.Solve(c)
	if e != nil {
		return Context{}, tools.NewWrappedError(c.from, ThousandsExpression{}, HundredsExpression{}, e)
	}

	return r, nil
}
