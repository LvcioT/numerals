package expressions_test

import (
	"slices"
	"testing"

	"taurino.com/numerals/internal/types/roman/expressions"
	"taurino.com/numerals/internal/types/roman/internals"
)

func testRightCases(t *testing.T, solveExp expressions.ExpressionSolve, testCases ...[]internals.TestCase) {
	t.Helper()

	aggreateCases := slices.Concat(testCases...)

	for _, tc := range aggreateCases {
		t.Run(tc.Name, func(t *testing.T) {
			c := expressions.NewContextFromString(tc.From)
			result, err := solveExp(c)
			if err != nil {
				t.Fatal(err)
			}

			value := result.GetValue()

			if value != tc.To {
				t.Errorf("Expected %d, got %d", tc.To, value)
			}
		})
	}
}

func testWrongCases(t *testing.T, solveExp expressions.ExpressionSolve, testCases ...[]internals.TestCase) {
	t.Helper()

	aggreateCases := slices.Concat(testCases...)

	for _, tc := range aggreateCases {
		t.Run(tc.Name, func(t *testing.T) {
			c := expressions.NewContextFromString(tc.From)
			result, err := solveExp(c)
			if err != nil {
				return
			}

			value := result.GetValue()

			t.Errorf("Expected %d value fail but got %d", tc.To, value)
		})
	}
}
