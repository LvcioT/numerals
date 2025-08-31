package expressions

import (
	"fmt"
)

func SolveVinculum(c Context) (Context, error) {
	cVinculum, cPlain, err := c.splitByVinculum()
	if err != nil {
		return Context{}, fmt.Errorf("cannot split by vinculum '%s': '%w'", c.from, err)
	}

	var resultVinculum, resultPlain Context
	if cVinculum.from != "" {
		resultVinculum, err = SolveVinculum(cVinculum)
		if err != nil {
			return Context{}, fmt.Errorf("cannot solve vinculum '%s': '%w'", cVinculum.from, err)
		}

		resultPlain, err = SolveHundreds(cPlain)
	} else {
		resultPlain, err = SolveThousands(cPlain)
	}

	if err != nil {
		return Context{}, fmt.Errorf("cannot solve '%s': '%w'", cPlain.from, err)
	}

	result, err := NewContextFromVinculumConcatenation(resultVinculum, resultPlain)
	if err != nil {
		return Context{}, fmt.Errorf("cannot concatenate with vinculum '%s'+'%s': '%w'", resultVinculum.from, resultPlain.from, err)
	}

	return result, nil
}
