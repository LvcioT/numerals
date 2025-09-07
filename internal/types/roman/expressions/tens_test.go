package expressions_test

import (
	"testing"

	"taurino.com/numerals/internal/types/roman/expressions"
	"taurino.com/numerals/internal/types/roman/internals"
)

var tensSuccessCases = []internals.TestCase{
	{Name: "X to 10", From: "X", To: 10},
	{Name: "XX to 20", From: "XX", To: 20},
	{Name: "XXX to 30", From: "XXX", To: 30},
	{Name: "XL to 40", From: "XL", To: 40},
	{Name: "L to 50", From: "L", To: 50},
	{Name: "LX to 60", From: "LX", To: 60},
	{Name: "LXX to 70", From: "LXX", To: 70},
	{Name: "LXXX to 80", From: "LXXX", To: 80},
	{Name: "XC to 90", From: "XC", To: 90},
	{Name: "XIX to 19", From: "XIX", To: 19},
	{Name: "XLIV to 44", From: "XLIV", To: 44},
	{Name: "LIX to 59", From: "LIX", To: 59},
	{Name: "XCIX to 99", From: "XCIX", To: 99},
}

var tensFailCases = []internals.TestCase{
	{Name: "XXXX", From: "XXXX"},
	{Name: "LL", From: "LL"},
	{Name: "XXL", From: "XXL"},
}

func TestTens(t *testing.T) {
	t.Run("tens right cases", func(t *testing.T) {
		testRightCases(t, expressions.TensExpression{}, tensSuccessCases, zeroSuccessCases)
	})

	t.Run("tens wrong cases", func(t *testing.T) {
		testWrongCases(t, expressions.TensExpression{}, tensFailCases, unitsFailCases, hundredsSuccessCases)
	})
}
