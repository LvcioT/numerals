package expressions_test

import (
	"testing"

	"taurino.com/numerals/internal/types/arabic/expressions"
	"taurino.com/numerals/internal/types/arabic/internals"
)

var thousandsMoreSuccessCases = []internals.TestCase{
	{Name: "(IV) value to 4k", To: "(IV)", From: 4_000},
	{Name: "(X) value to 10k", To: "(X)", From: 10_000},
	{Name: "(IX) value to 9k", To: "(IX)", From: 9_000},
	{Name: "(C) value to 100k", To: "(C)", From: 100_000},
	{Name: "(CM) value to 900k", To: "(CM)", From: 900_000},
	{Name: "(MC)LV value to 1100055", To: "(MC)LV", From: 1_100_055},
	{Name: "((MC)LV)CV value to 1_100_055_105", To: "((MC)LV)CV", From: 1_100_055_105},
}

func TestThousandMore(t *testing.T) {
	t.Run("thousands more right cases", func(t *testing.T) {
		testRightCases(t, expressions.SolveThousandMore, thousandsMoreSuccessCases, thousandsLessSuccessCases)
	})
}
