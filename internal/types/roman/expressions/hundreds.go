package expressions

import (
	"fmt"
	"strings"

	"taurino.com/numerals/tools"
)

type HundredsExpression struct{}

func (exp HundredsExpression) Solve(c Context) (Context, error) {
	// must be ordered by length
	patterns := []Context{
		{from: "DCCC", to: 800},
		{from: "CCC", to: 300},
		{from: "DCC", to: 700},
		{from: "CC", to: 200},
		{from: "CD", to: 400},
		{from: "DC", to: 600},
		{from: "CM", to: 900},
		{from: "C", to: 100},
		{from: "D", to: 500},
	}

	for _, pattern := range patterns {
		prefix := pattern.from
		value := pattern.to

		if strings.HasPrefix(c.from, prefix) {
			valueLeft := value

			contextI := Context{from: strings.TrimPrefix(c.from, prefix)}
			contextO, e := TensExpression{}.Solve(contextI)
			if e != nil {
				return Context{}, tools.NewWrappedError(contextO.from, HundredsExpression{}, TensExpression{}, e)
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

	// no hundreds present, check for smaller quantity
	r, e := TensExpression{}.Solve(c)
	if e != nil {
		return Context{}, tools.NewWrappedError(c.from, HundredsExpression{}, TensExpression{}, e)
	}

	return r, nil
}
