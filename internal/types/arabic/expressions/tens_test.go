package expressions_test

import (
	"testing"

	"taurino.com/numerals/internal/types/roman/expressions"
	"taurino.com/numerals/internal/types/roman/internals"
)

var tensSuccessCases = []internals.TestCase{
	{Name: "X value 10", From: "X", To: 10},
	{Name: "XX value 20", From: "XX", To: 20},
	{Name: "XXX value 30", From: "XXX", To: 30},
	{Name: "XL value 40", From: "XL", To: 40},
	{Name: "L value 50", From: "L", To: 50},
	{Name: "LX value 60", From: "LX", To: 60},
	{Name: "LXX value 70", From: "LXX", To: 70},
	{Name: "LXXX value 80", From: "LXXX", To: 80},
	{Name: "XC value 90", From: "XC", To: 90},
	{Name: "XIX value 19", From: "XIX", To: 19},
	{Name: "XLIV value 44", From: "XLIV", To: 44},
	{Name: "LIX value 59", From: "LIX", To: 59},
	{Name: "XCIX value 99", From: "XCIX", To: 99},
}

var tensFailCases = []internals.TestCase{
	{Name: "XXXX", From: "XXXX"},
	{Name: "LL", From: "LL"},
	{Name: "XXL", From: "XXL"},
}

func TestTens(t *testing.T) {
	t.Run("tens right cases", func(t *testing.T) {
		testRightCases(t, expressions.SolveTens, tensSuccessCases, unitsSuccessCases)
	})

	t.Run("tens wrong cases", func(t *testing.T) {
		testWrongCases(t, expressions.SolveTens, tensFailCases, unitsFailCases, hundredsSuccessCases)
	})
}
