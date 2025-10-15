package expressions_test

import (
	"testing"

	"taurino.com/numerals/internal/types/roman/expressions"
	"taurino.com/numerals/internal/types/roman/internals"
)

var unitsSuccessCases = []internals.TestCase{
	{Name: "I value 1", From: "I", To: 1},
	{Name: "II value 2", From: "II", To: 2},
	{Name: "III value 3", From: "III", To: 3},
	{Name: "IV value 4", From: "IV", To: 4},
	{Name: "V value 5", From: "V", To: 5},
	{Name: "VI value 6", From: "VI", To: 6},
	{Name: "VII value 7", From: "VII", To: 7},
	{Name: "VIII value 8", From: "VIII", To: 8},
	{Name: "IX value 9", From: "IX", To: 9},
}

var unitsFailCases = []internals.TestCase{
	{Name: "IIII", From: "IIII"},
	{Name: "VV", From: "VV"},
	{Name: "IIV", From: "IIV"},
}

func TestUnits(t *testing.T) {
	t.Run("units right cases", func(t *testing.T) {
		testRightCases(t, expressions.SolveUnit, unitsSuccessCases, zeroSuccessCases)
	})

	t.Run("units wrong cases", func(t *testing.T) {
		testWrongCases(t, expressions.SolveUnit, unitsFailCases, zeroFailCases, tensSuccessCases)
	})
}
