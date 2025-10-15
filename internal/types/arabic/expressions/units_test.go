package expressions_test

import (
	"testing"

	"taurino.com/numerals/internal/types/arabic/expressions"
	"taurino.com/numerals/internal/types/arabic/internals"
)

var unitsSuccessCases = []internals.TestCase{
	{Name: "I value 1", To: "I", From: 1},
	{Name: "II value 2", To: "II", From: 2},
	{Name: "III value 3", To: "III", From: 3},
	{Name: "IV value 4", To: "IV", From: 4},
	{Name: "V value 5", To: "V", From: 5},
	{Name: "VI value 6", To: "VI", From: 6},
	{Name: "VII value 7", To: "VII", From: 7},
	{Name: "VIII value 8", To: "VIII", From: 8},
	{Name: "IX value 9", To: "IX", From: 9},
}

func TestUnits(t *testing.T) {
	t.Run("units right cases", func(t *testing.T) {
		testRightCases(t, expressions.SolveUnit, unitsSuccessCases, zeroSuccessCases)
	})

	t.Run("units wrong cases", func(t *testing.T) {
		testWrongCases(t, expressions.SolveUnit, tensSuccessCases)
	})
}
