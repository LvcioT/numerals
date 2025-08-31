package expressions_test

import (
	"testing"

	"taurino.com/numerals/internal/types/roman/expressions"
	"taurino.com/numerals/internal/types/roman/internals"
)

var hundredsSuccessCases = []internals.TestCase{
	{Name: "C value 100", From: "C", To: 100},
	{Name: "CC value 200", From: "CC", To: 200},
	{Name: "CCC value 300", From: "CCC", To: 300},
	{Name: "CD value 400", From: "CD", To: 400},
	{Name: "D value 500", From: "D", To: 500},
	{Name: "DC value 600", From: "DC", To: 600},
	{Name: "DCC value 700", From: "DCC", To: 700},
	{Name: "DCCC value 800", From: "DCCC", To: 800},
	{Name: "CM value 900", From: "CM", To: 900},
	{Name: "CIX value 109", From: "CIX", To: 109},
	{Name: "CDXLIV value 444", From: "CDXLIV", To: 444},
	{Name: "DCLXVI value 666", From: "DCLXVI", To: 666},
	{Name: "CMXCIX value 999", From: "CMXCIX", To: 999},
}

var hundredsFailCases = []internals.TestCase{
	{Name: "CCCC", From: "CCCC"},
	{Name: "DD", From: "DD"},
	{Name: "CCD", From: "CCD"},
}

func TestHundreds(t *testing.T) {
	t.Run("hundreds right cases", func(t *testing.T) {
		testRightCases(t, expressions.SolveHundreds, hundredsSuccessCases, tensSuccessCases)
	})

	t.Run("hundreds wrong cases", func(t *testing.T) {
		testWrongCases(t, expressions.SolveHundreds, hundredsFailCases, tensFailCases, thousandsSuccessCases)
	})
}
