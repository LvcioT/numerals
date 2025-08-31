package expressions_test

import (
	"testing"

	"taurino.com/numerals/internal/types/roman/expressions"
	"taurino.com/numerals/internal/types/roman/internals"
)

var vinculumSuccessCases = []internals.TestCase{
	{Name: "(I) value to 1k", From: "(I)", To: 1_000},
	{Name: "(IV) value to 4k", From: "(IV)", To: 4_000},
	{Name: "(X) value to 10k", From: "(X)", To: 10_000},
	{Name: "(IX) value to 9k", From: "(IX)", To: 9_000},
	{Name: "(C) value to 100k", From: "(C)", To: 100_000},
	{Name: "(CM) value to 900k", From: "(CM)", To: 900_000},
	{Name: "(MC)LV value to 1100055", From: "(MC)LV", To: 1_100_055},
	{Name: "((MC)LV)CV value to 1_100_055_105", From: "((MC)LV)CV", To: 1_100_055_105},
}

var vinculumFailCases = []internals.TestCase{
	{Name: "thousands in plain", From: "(C)MI"},
	{Name: "unmatched parenthesis", From: "(II"},
	{Name: "overflow", From: "((((((M))))))"},
}

func TestVinculum(t *testing.T) {
	t.Run("vinculum right cases", func(t *testing.T) {
		testRightCases(t, expressions.SolveVinculum, vinculumSuccessCases, hundredsSuccessCases)
	})

	t.Run("vinculum wrong cases", func(t *testing.T) {
		testWrongCases(t, expressions.SolveVinculum, vinculumFailCases, hundredsFailCases)
	})
}
