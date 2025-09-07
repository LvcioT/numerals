package expressions_test

import (
	"slices"
	"testing"

	"taurino.com/numerals/internal/types/roman/expressions"
	"taurino.com/numerals/internal/types/roman/internals"
)

var thousandsSuccessCases = []internals.TestCase{
	{Name: "M to 1000", From: "M", To: 1000},
	{Name: "MM to 2000", From: "MM", To: 2000},
	{Name: "MMM to 3000", From: "MMM", To: 3000},
	{Name: "MCMLXXXIX to 1989", From: "MCMLXXXIX", To: 1989},
	{Name: "MMXXIV to 2024", From: "MMXXIV", To: 2024},
	{Name: "biggest without winculum", From: "MMMCMXCIX", To: 3999},
}

var thousandsFailCases = []internals.TestCase{
	{Name: "MMMM", From: "MMMM"},
}

func TestThousands(t *testing.T) {
	successCases := slices.Concat(thousandsSuccessCases, hundredsSuccessCases)

	t.Run("thousands right cases", func(t *testing.T) {
		for _, tc := range successCases {
			t.Run(tc.Name, func(t *testing.T) {
				c := expressions.NexContextFromString(tc.From)
				result, err := expressions.ThousandsExpression{}.Solve(c)
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

	failCases := slices.Concat(thousandsFailCases, hundredsFailCases)

	t.Run("thousands wrong cases", func(t *testing.T) {
		for _, tc := range failCases {
			t.Run(tc.Name, func(t *testing.T) {
				c := expressions.NexContextFromString(tc.From)
				result, err := expressions.ThousandsExpression{}.Solve(c)
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
