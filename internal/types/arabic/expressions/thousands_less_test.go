package expressions_test

import (
	"testing"

	"taurino.com/numerals/internal/types/arabic/expressions"
	"taurino.com/numerals/internal/types/arabic/internals"
)

var thousandsLessSuccessCases = []internals.TestCase{
	{Name: "M value 1000", To: "M", From: 1000},
	{Name: "MM value 2000", To: "MM", From: 2000},
	{Name: "MMM value 3000", To: "MMM", From: 3000},
	{Name: "MCMLXXXIX value 1989", To: "MCMLXXXIX", From: 1989},
	{Name: "MMXXIV value 2024", To: "MMXXIV", From: 2024},
	{Name: "biggest without winculum", To: "MMMCMXCIX", From: 3999},
}

func TestThousands(t *testing.T) {
	t.Run("thousands more right cases", func(t *testing.T) {
		testRightCases(t, expressions.SolveThousandLess, thousandsLessSuccessCases, hundredsSuccessCases)
	})

	t.Run("thousands more wrong cases", func(t *testing.T) {
		testWrongCases(t, expressions.SolveThousandLess, thousandsMoreSuccessCases)
	})
}
