package expressions_test

import (
	"testing"

	"taurino.com/numerals/internal/types/arabic/expressions"
	"taurino.com/numerals/internal/types/arabic/internals"
)

var zeroSuccessCases = []internals.TestCase{
	{Name: "empty string value 0", To: "", From: 0},
}

func TestZeros(t *testing.T) {
	t.Run("zero right cases", func(t *testing.T) {
		testRightCases(t, expressions.SolveZero, zeroSuccessCases)
	})

	t.Run("zero wrong cases", func(t *testing.T) {
		testWrongCases(t, expressions.SolveZero, unitsSuccessCases)
	})
}
