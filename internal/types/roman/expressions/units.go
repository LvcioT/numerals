package expressions

import (
	"fmt"
	"strings"

	"taurino.com/numerals/tools"
)

type UnitsExpression struct{}

func (exp UnitsExpression) Solve(c Context) (Context, error) {
	// must be ordered by length
	prefixes := map[string]uint64{
		"VIII": 8,
		"III":  3,
		"VII":  7,
		"II":   2,
		"IV":   4,
		"VI":   6,
		"IX":   9,
		"I":    1,
		"V":    5,
	}

	for prefix, value := range prefixes {
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
