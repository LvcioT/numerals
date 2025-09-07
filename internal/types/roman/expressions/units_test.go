package expressions_test

import (
	"slices"
	"testing"

	"taurino.com/numerals/internal/types/roman/expressions"
	"taurino.com/numerals/internal/types/roman/internals"
)

var unitsSuccessCases = []internals.TestCase{
	{Name: "I to 1", From: "I", To: 1},
	{Name: "II to 2", From: "II", To: 2},
	{Name: "III to 3", From: "III", To: 3},
	{Name: "IV to 4", From: "IV", To: 4},
	{Name: "V to 5", From: "V", To: 5},
	{Name: "VI to 6", From: "VI", To: 6},
	{Name: "VII to 7", From: "VII", To: 7},
	{Name: "VIII to 8", From: "VIII", To: 8},
	{Name: "IX to 9", From: "IX", To: 9},
}

var unitsFailCases = []internals.TestCase{
	{Name: "IIII", From: "IIII"},
	{Name: "VV", From: "VV"},
	{Name: "IIV", From: "IIV"},
}

func TestUnits(t *testing.T) {
	successCases := slices.Concat(unitsSuccessCases, zeroSuccessCases)

	t.Run("units right cases", func(t *testing.T) {
		for _, tc := range successCases {
			t.Run(tc.Name, func(t *testing.T) {
				c := expressions.NexContextFromString(tc.From)
				result, err := expressions.UnitsExpression{}.Solve(c)
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

	failCases := slices.Concat(unitsFailCases, zeroFailCases, tensSuccessCases)

	t.Run("units wrong cases", func(t *testing.T) {
		for _, tc := range failCases {
			t.Run(tc.Name, func(t *testing.T) {
				c := expressions.NexContextFromString(tc.From)
				result, err := expressions.UnitsExpression{}.Solve(c)
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
