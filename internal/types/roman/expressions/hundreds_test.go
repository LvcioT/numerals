package expressions_test

import (
	"testing"

	"taurino.com/numerals/internal/types/roman/expressions"
	"taurino.com/numerals/internal/types/roman/internals"
)

var hundredsSuccessCases = []internals.TestCase{
	{Name: "C to 100", From: "C", To: 100},
	{Name: "CC to 200", From: "CC", To: 200},
	{Name: "CCC to 300", From: "CCC", To: 300},
	{Name: "CD to 400", From: "CD", To: 400},
	{Name: "D to 500", From: "D", To: 500},
	{Name: "DC to 600", From: "DC", To: 600},
	{Name: "DCC to 700", From: "DCC", To: 700},
	{Name: "DCCC to 800", From: "DCCC", To: 800},
	{Name: "CM to 900", From: "CM", To: 900},
	{Name: "CIX to 109", From: "CIX", To: 109},
	{Name: "CDXLIV to 444", From: "CDXLIV", To: 444},
	{Name: "DCLXVI to 666", From: "DCLXVI", To: 666},
	{Name: "CMXCIX to 999", From: "CMXCIX", To: 999},
}

var hundredsFailCases = []internals.TestCase{
	{Name: "CCCC", From: "CCCC"},
	{Name: "DD", From: "DD"},
	{Name: "CCD", From: "CCD"},
}

func TestHundreds(t *testing.T) {
	t.Run("hundreds right cases", func(t *testing.T) {
		testRightCases(t, expressions.HundredsExpression{}, hundredsSuccessCases, tensSuccessCases)
	})

	t.Run("hundreds wrong cases", func(t *testing.T) {
		testWrongCases(t, expressions.HundredsExpression{}, hundredsFailCases, tensFailCases)
	})
}
