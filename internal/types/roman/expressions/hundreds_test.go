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
	successCases := append(hundredsSuccessCases, tensSuccessCases...)

	t.Run("hundreds right cases", func(t *testing.T) {
		for _, tc := range successCases {
			t.Run(tc.Name, func(t *testing.T) {
				c := expressions.NexContextFromString(tc.From)
				result, err := expressions.HundredsExpression{}.Solve(c)
				if err != nil {
					t.Fatal(err)
				}

				value, err := result.GetValue()
				if err != nil {
					t.Fatal(err)
				}

				if value != tc.To {
					t.Errorf("Expected %d, got %d", tc.To, value)
				}
			})
		}
	})

	failCases := append(hundredsFailCases, tensFailCases...)
	failCases = append(failCases, thousandsSuccessCases...)

	t.Run("hundreds wrong cases", func(t *testing.T) {
		for _, tc := range failCases {
			t.Run(tc.Name, func(t *testing.T) {
				c := expressions.NexContextFromString(tc.From)
				result, err := expressions.HundredsExpression{}.Solve(c)
				if err != nil {
					return
				}

				value, err := result.GetValue()
				if err != nil {
					t.Fatal(err)
				}

				t.Errorf("Expected %d to fail but got %d", tc.To, value)
			})
		}
	})
}
