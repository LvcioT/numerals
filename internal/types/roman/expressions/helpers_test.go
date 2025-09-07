package expressions_test

import (
	"slices"
	"testing"

	"taurino.com/numerals/internal/types/roman/expressions"
	"taurino.com/numerals/internal/types/roman/internals"
)

type TestExpression interface {
	Solve(expressions.Context) (expressions.Context, error)
}

func testRightCases(t *testing.T, exp TestExpression, testCases ...[]internals.TestCase) {
	t.Helper()

	aggreateCases := slices.Concat(testCases...)

	for _, tc := range aggreateCases {
		t.Run(tc.Name, func(t *testing.T) {
			c := expressions.NexContextFromString(tc.From)
			result, err := exp.Solve(c)
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
}

func testWrongCases(t *testing.T, exp TestExpression, testCases ...[]internals.TestCase) {
	t.Helper()

	aggreateCases := slices.Concat(testCases...)

	for _, tc := range aggreateCases {
		t.Run(tc.Name, func(t *testing.T) {
			c := expressions.NexContextFromString(tc.From)
			result, err := exp.Solve(c)
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
}
