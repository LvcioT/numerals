package expressions_test

import (
	"slices"
	"testing"

	"taurino.com/numerals/internal/types/roman/expressions"
	"taurino.com/numerals/internal/types/roman/internals"
)

var zeroSuccessCases = []internals.TestCase{
	{Name: "empty string to 0", From: "", To: 0},
}

var zeroFailCases = []internals.TestCase{
	{Name: "other symbols", From: "A"},
	{Name: "allowed symbols", From: "X"},
	{Name: "spaces", From: " "},
}

func TestZeros(t *testing.T) {
	t.Run("zero right cases", func(t *testing.T) {
		for _, tc := range zeroSuccessCases {
			t.Run(tc.Name, func(t *testing.T) {
				c := expressions.NexContextFromString(tc.From)
				result, err := expressions.ZeroExpression{}.Solve(c)
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

	failCases := slices.Concat(zeroFailCases, unitsSuccessCases)

	t.Run("zero wrong cases", func(t *testing.T) {
		for _, tc := range failCases {
			t.Run(tc.Name, func(t *testing.T) {
				c := expressions.NexContextFromString(tc.From)
				result, err := expressions.ZeroExpression{}.Solve(c)
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
