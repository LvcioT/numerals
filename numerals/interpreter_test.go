package numerals_test

import (
	"testing"

	"taurino.com/numerals/data"
	"taurino.com/numerals/numerals"
)

func TestInterpreter(t *testing.T) {
	t.Run("BaseCases", func(t *testing.T) {
		for _, tc := range data.BaseCases {
			t.Run(tc.Name, func(t *testing.T) {
				result := numerals.ToArabic(tc.From)
				if uint64(result) != tc.To {
					t.Errorf("Expected %d, got %d", tc.To, result)
				}
			})
		}
	})

	t.Run("WithVinculumCases", func(t *testing.T) {
		for _, tc := range data.WithVinculumCases {
			t.Run(tc.Name, func(t *testing.T) {
				result := numerals.ToArabic(tc.From)
				if uint64(result) != tc.To {
					t.Errorf("Expected %d, got %d", tc.To, result)
				}
			})
		}
	})

	t.Run("ErrorCases", func(t *testing.T) {
		for _, tc := range data.ErrorCases {
			t.Run(tc.Name, func(t *testing.T) {
				result := numerals.ToArabic(tc.From)
				if uint64(result) != tc.To {
					t.Errorf("Expected %d, got %d", tc.To, result)
				}
			})
		}
	})
}
