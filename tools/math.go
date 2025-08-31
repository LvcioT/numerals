package tools

import (
	"errors"
	"math"
)

// AddSafe performs the addition of two uint64 numbers and returns a boolean indicating if an overflow occurred.
func AddSafe(a, b uint64) (uint64, error) {
	if math.MaxUint64-a < b {
		return 0, errors.New("cannot sum %d and %d: overflow")
	}
	return a + b, nil
}

// MultiplySafe performs the product of two uint64 numbers and returns a boolean indicating if an overflow occurred.
func MultiplySafe(a, b uint64) (uint64, error) {
	if a == 0 || b == 0 {
		return 0, nil
	}

	if math.MaxUint64/a < b {
		return 0, errors.New("cannot multiply %d and %d: overflow")
	}
	return a * b, nil
}
