package arabic_test

import (
	"testing"

	"taurino.com/numerals/data"
	"taurino.com/numerals/internal/types/arabic"
)

func TestInterpreter(t *testing.T) {
	t.Run("BaseCases", func(t *testing.T) {
		for _, tc := range data.BaseCases {
			t.Run(tc.Name, func(t *testing.T) {
				result, err := arabic.Interpreter{}.Interpret(tc.Arabic)
				if err != nil {
					t.Fatalf("Expected no error, but got %v", err)
				}
				val := result.GetTo()
				if val != tc.Roman {
					t.Errorf("Expected %s, got %s", tc.Roman, val)
				}
			})
		}
	})

	t.Run("WithVinculumCases", func(t *testing.T) {
		for _, tc := range data.WithVinculumCases {
			t.Run(tc.Name, func(t *testing.T) {
				result, err := arabic.Interpreter{}.Interpret(tc.Arabic)
				if err != nil {
					t.Fatalf("Expected no error, but got %v", err)
				}
				val := result.GetTo()
				if val != tc.Roman {
					t.Errorf("Expected %s, got %s", tc.Roman, val)
				}
			})
		}
	})
}
