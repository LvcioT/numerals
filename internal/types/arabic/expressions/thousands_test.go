package expressions_test

import (
	"testing"

	"taurino.com/numerals/internal/types/roman/expressions"
	"taurino.com/numerals/internal/types/roman/internals"
)

var thousandsSuccessCases = []internals.TestCase{
	{Name: "M value 1000", From: "M", To: 1000},
	{Name: "MM value 2000", From: "MM", To: 2000},
	{Name: "MMM value 3000", From: "MMM", To: 3000},
	{Name: "MCMLXXXIX value 1989", From: "MCMLXXXIX", To: 1989},
	{Name: "MMXXIV value 2024", From: "MMXXIV", To: 2024},
	{Name: "biggest without winculum", From: "MMMCMXCIX", To: 3999},
}

var thousandsFailCases = []internals.TestCase{
	{Name: "MMMM", From: "MMMM"},
}

func TestThousands(t *testing.T) {
	t.Run("thousands right cases", func(t *testing.T) {
		testRightCases(t, expressions.SolveThousands, thousandsSuccessCases, hundredsSuccessCases)
	})

	t.Run("thousands wrong cases", func(t *testing.T) {
		testWrongCases(t, expressions.SolveThousands, thousandsFailCases, hundredsFailCases, vinculumSuccessCases)
	})
}
