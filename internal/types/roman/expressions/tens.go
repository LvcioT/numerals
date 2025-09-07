package expressions

import (
	"fmt"
	"strings"

	"taurino.com/numerals/tools"
)

type TensExpression struct{}

func (exp TensExpression) Solve(c Context) (Context, error) {
	// ordered patterns from highest to lowest
	patterns := []Context{
		{from: "LXXX", to: 80},
		{from: "XXX", to: 30},
		{from: "LXX", to: 70},
		{from: "XX", to: 20},
		{from: "XL", to: 40},
		{from: "LX", to: 60},
		{from: "XC", to: 90},
		{from: "X", to: 10},
		{from: "L", to: 50},
	}

	for _, pattern := range patterns {
		prefix := pattern.from
		value := pattern.to

		if strings.HasPrefix(c.from, prefix) {
			valueLeft := value

			contextI := Context{from: strings.TrimPrefix(c.from, prefix)}
			contextO, e := UnitsExpression{}.Solve(contextI)
			if e != nil {
				return Context{}, tools.NewWrappedError(contextO.from, TensExpression{}, UnitsExpression{}, e)
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

	// no tens present, check for smaller quantity
	r, e := UnitsExpression{}.Solve(c)
	if e != nil {
		return Context{}, tools.NewWrappedError(c.from, TensExpression{}, UnitsExpression{}, e)
	}

	return r, nil
}
