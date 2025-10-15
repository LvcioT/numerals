package expressions_test

import (
	"slices"
	"testing"

	"taurino.com/numerals/internal/types/arabic/expressions"
	"taurino.com/numerals/internal/types/arabic/internals"
)

func testRightCases(t *testing.T, solveExp expressions.ExpressionSolve, testCases ...[]internals.TestCase) {
	t.Helper()

	aggregateCases := slices.Concat(testCases...)

	for _, tc := range aggregateCases {
		t.Run(tc.Name, func(t *testing.T) {
			c := expressions.NewContextFromValue(tc.From)
			result, err := solveExp(c)
			if err != nil {
				t.Fatal(err)
			}

			value := result.GetTo()

			if value != tc.To {
				t.Errorf("Expected %s, got %s", tc.To, value)
			}
		})
	}
}

func testWrongCases(t *testing.T, solveExp expressions.ExpressionSolve, testCases ...[]internals.TestCase) {
	t.Helper()

	aggregateCases := slices.Concat(testCases...)

	for _, tc := range aggregateCases {
		t.Run(tc.Name, func(t *testing.T) {
			c := expressions.NewContextFromValue(tc.From)
			result, err := solveExp(c)
			if err != nil {
				return
			}

			value := result.GetTo()

			t.Errorf("Expected %s value fail but got %s", tc.To, value)
		})
	}
}
