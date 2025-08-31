package expressions

import (
	"fmt"
	"strings"

	"taurino.com/numerals/tools"
)

type TensExpression struct{}

func (exp TensExpression) Solve(c Context) (Context, error) {
	// must be ordered by length
	prefixes := map[string]uint64{
		"LXXX": 80,
		"XXX":  30,
		"LXX":  70,
		"XX":   20,
		"XL":   40,
		"LX":   60,
		"XC":   90,
		"X":    10,
		"L":    50,
	}

	for prefix, value := range prefixes {
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
