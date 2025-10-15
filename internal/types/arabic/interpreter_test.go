package roman_test

import (
	"testing"

	"taurino.com/numerals/data"
	"taurino.com/numerals/internal/types/roman"
)

func TestInterpreter(t *testing.T) {
	t.Run("BaseCases", func(t *testing.T) {
		for _, tc := range data.BaseCases {
			t.Run(tc.Name, func(t *testing.T) {
				result, err := roman.Interpreter{}.Interpret(tc.From)
				if err != nil {
					t.Fatalf("Expected no error, but got %v", err)
				}
				val := result.GetValue()
				if val != tc.To {
					t.Errorf("Expected %d, got %d", tc.To, val)
				}
			})
		}
	})

	t.Run("WithVinculumCases", func(t *testing.T) {
		for _, tc := range data.WithVinculumCases {
			t.Run(tc.Name, func(t *testing.T) {
				result, err := roman.Interpreter{}.Interpret(tc.From)
				if err != nil {
					t.Fatalf("Expected no error, but got %v", err)
				}
				val := result.GetValue()
				if val != tc.To {
					t.Errorf("Expected %d, got %d", tc.To, val)
				}
			})
		}
	})

	t.Run("ErrorCases", func(t *testing.T) {
		for _, tc := range data.ErrorCases {
			t.Run(tc.Name, func(t *testing.T) {
				_, err := roman.Interpreter{}.Interpret(tc.From)
				if err == nil {
					t.Errorf("Expected an error, but got none on '%s'", tc.From)
				}
			})
		}
	})
}
