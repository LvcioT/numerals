package expressions_test

import (
	"testing"

	"taurino.com/numerals/internal/types/arabic/expressions"
	"taurino.com/numerals/internal/types/arabic/internals"
)

var hundredsSuccessCases = []internals.TestCase{
	{Name: "C value 100", To: "C", From: 100},
	{Name: "CC value 200", To: "CC", From: 200},
	{Name: "CCC value 300", To: "CCC", From: 300},
	{Name: "CD value 400", To: "CD", From: 400},
	{Name: "D value 500", To: "D", From: 500},
	{Name: "DC value 600", To: "DC", From: 600},
	{Name: "DCC value 700", To: "DCC", From: 700},
	{Name: "DCCC value 800", To: "DCCC", From: 800},
	{Name: "CM value 900", To: "CM", From: 900},
	{Name: "CIX value 109", To: "CIX", From: 109},
	{Name: "CDXLIV value 444", To: "CDXLIV", From: 444},
	{Name: "DCLXVI value 666", To: "DCLXVI", From: 666},
	{Name: "CMXCIX value 999", To: "CMXCIX", From: 999},
}

func TestHundreds(t *testing.T) {
	t.Run("hundreds right cases", func(t *testing.T) {
		testRightCases(t, expressions.SolveHundreds, hundredsSuccessCases, tensSuccessCases)
	})

	t.Run("hundreds wrong cases", func(t *testing.T) {
		testWrongCases(t, expressions.SolveHundreds, thousandsLessSuccessCases)
	})
}
