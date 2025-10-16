package expressions_test

import (
	"testing"

	"taurino.com/numerals/internal/types/arabic/expressions"
	"taurino.com/numerals/internal/types/arabic/internals"
)

var zeroSuccessCases = []internals.TestCase{
	{Name: "empty string value 0", From: "", To: 0},
}

var zeroFailCases = []internals.TestCase{
	{Name: "other symbols", From: "A"},
	{Name: "allowed symbols", From: "X"},
	{Name: "spaces", From: " "},
}

func TestZeros(t *testing.T) {
	t.Run("zero right cases", func(t *testing.T) {
		testRightCases(t, expressions.SolveZero, zeroSuccessCases)
	})

	t.Run("zero wrong cases", func(t *testing.T) {
		testWrongCases(t, expressions.SolveZero, zeroFailCases, unitsSuccessCases)
	})
}
