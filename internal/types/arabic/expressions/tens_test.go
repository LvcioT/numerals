package expressions_test

import (
	"testing"

	"taurino.com/numerals/internal/types/arabic/expressions"
	"taurino.com/numerals/internal/types/arabic/internals"
)

var tensSuccessCases = []internals.TestCase{
	{Name: "X value 10", To: "X", From: 10},
	{Name: "XX value 20", To: "XX", From: 20},
	{Name: "XXX value 30", To: "XXX", From: 30},
	{Name: "XL value 40", To: "XL", From: 40},
	{Name: "L value 50", To: "L", From: 50},
	{Name: "LX value 60", To: "LX", From: 60},
	{Name: "LXX value 70", To: "LXX", From: 70},
	{Name: "LXXX value 80", To: "LXXX", From: 80},
	{Name: "XC value 90", To: "XC", From: 90},
	{Name: "XIX value 19", To: "XIX", From: 19},
	{Name: "XLIV value 44", To: "XLIV", From: 44},
	{Name: "LIX value 59", To: "LIX", From: 59},
	{Name: "XCIX value 99", To: "XCIX", From: 99},
}

func TestTens(t *testing.T) {
	t.Run("tens right cases", func(t *testing.T) {
		testRightCases(t, expressions.SolveTens, tensSuccessCases, unitsSuccessCases)
	})

	t.Run("tens wrong cases", func(t *testing.T) {
		testWrongCases(t, expressions.SolveTens, hundredsSuccessCases)
	})
}
