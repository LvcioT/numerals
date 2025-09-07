package expressions

import (
	"fmt"
	"strings"

	"taurino.com/numerals/tools"
)

type UnitsExpression struct{}

func (exp UnitsExpression) Solve(c Context) (Context, error) {
	// must be ordered by length
	patterns := []Context{
		Context{from: "VIII", to: 8},
		Context{from: "III", to: 3},
		Context{from: "VII", to: 7},
		Context{from: "II", to: 2},
		Context{from: "IV", to: 4},
		Context{from: "VI", to: 6},
		Context{from: "IX", to: 9},
		Context{from: "I", to: 1},
		Context{from: "V", to: 5},
	}

	for _, pattern := range patterns {
		prefix := pattern.from
		value := pattern.to

		if strings.HasPrefix(c.from, prefix) {
			valueLeft := value

			contextI := Context{from: strings.TrimPrefix(c.from, prefix)}
			contextO, e := ZeroExpression{}.Solve(contextI)
			if e != nil {
				return Context{}, tools.NewWrappedError(contextO.from, UnitsExpression{}, ZeroExpression{}, e)
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

	// no units present, check if empty
	r, e := ZeroExpression{}.Solve(c)
	if e != nil {
		return Context{}, tools.NewWrappedError(c.from, UnitsExpression{}, ZeroExpression{}, e)
	}

	return r, nil
}
