package tools_test

import (
	"math"
	"testing"

	"taurino.com/numerals/tools"
)

func TestAdd(t *testing.T) {
	t.Run("no overflow", func(t *testing.T) {
		sum, error := tools.AddSafe(1, 2)
		if error != nil {
			t.Error("Expected no overflow, but got overflow")
		}
		if sum != 3 {
			t.Errorf("Expected sum to be 3, but got %d", sum)
		}
	})

	t.Run("max uint64", func(t *testing.T) {
		sum, error := tools.AddSafe(math.MaxUint64-1, 1)
		if error != nil {
			t.Error("Expected no overflow, but got overflow")
		}
		if sum != math.MaxUint64 {
			t.Errorf("Expected sum to be %d, but got %d", uint64(math.MaxUint64), sum)
		}
	})

	t.Run("overflow", func(t *testing.T) {
		_, error := tools.AddSafe(math.MaxUint64, 1)
		if error == nil {
			t.Error("Expected overflow, but got no overflow")
		}
	})
}

func TestMultiplySafe(t *testing.T) {
	t.Run("no overflow", func(t *testing.T) {
		product, err := tools.MultiplySafe(2, 3)
		if err != nil {
			t.Fatalf("expected no error, but got: %v", err)
		}
		if product != 6 {
			t.Errorf("expected 6, but got: %d", product)
		}
	})

	t.Run("overflow", func(t *testing.T) {
		_, err := tools.MultiplySafe(math.MaxUint64, 2)
		if err == nil {
			t.Fatal("expected an error, but got none")
		}
	})

	t.Run("no overflow with max value", func(t *testing.T) {
		product, err := tools.MultiplySafe(math.MaxUint64/2, 2)
		if err != nil {
			t.Fatalf("expected no error, but got: %v", err)
		}
		if product != math.MaxUint64/2*2 {
			t.Errorf("expected %d, but got: %d", uint64(math.MaxUint64/2*2), product)
		}
	})
}
