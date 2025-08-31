package expressions

import (
	"fmt"
	"strings"

	"taurino.com/numerals/tools"
)

type HundredsExpression struct{}

func (exp HundredsExpression) Solve(c Context) (Context, error) {
	// must be ordered by length
	prefixes := map[string]uint64{
		"DCCC": 800,
		"CCC":  300,
		"DCC":  700,
		"CC":   200,
		"CD":   400,
		"DC":   600,
		"CM":   900,
		"C":    100,
		"D":    500,
	}

	for prefix, value := range prefixes {
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
