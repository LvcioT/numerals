package pkg

import (
	"fmt"

	"taurino.com/numerals/internal/types/roman"
	"taurino.com/numerals/tools"
)

func NewInterpreter(from tools.NumeralType) (tools.Interpreter, error) {
	switch from {
	case tools.Roman:
		return roman.Interpreter{}, nil
	default:
		return nil, fmt.Errorf("cannot instanciate the proper interpreter for %d", from)
	}
}
