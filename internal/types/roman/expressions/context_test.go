package expressions

import (
	"math"
	"testing"
)

func TestContextGetters(t *testing.T) {
	c := Context{from: "X", value: 42}

	if got := c.GetFrom(); got != "X" {
		t.Errorf("GetFrom expected 'X', got %q", got)
	}
	if got := c.GetValue(); got != 42 {
		t.Errorf("GetValue expected 42, got %d", got)
	}
	if got := c.GetTo(); got != "42" {
		t.Errorf("GetTo expected '42', got %q", got)
	}
}

func TestNewContextFromString(t *testing.T) {
	c := NewContextFromString("XIV")
	if c.from != "XIV" {
		t.Errorf("from expected 'XIV', got %q", c.from)
	}
	if c.value != 0 {
		t.Errorf("value expected 0, got %d", c.value)
	}
}

func TestSplitByPrefix_LongestMatch(t *testing.T) {
	c := Context{from: "XIV"}
	patterns := []Context{{from: "X", value: 10}, {from: "XI", value: 11}}

	left, right := c.splitByPrefix(patterns)

	if left.from != "XI" || left.value != 11 {
		t.Errorf("left expected {from:'XI', value:11}, got {from:%q, value:%d}", left.from, left.value)
	}
	if right.from != "V" || right.value != 0 {
		t.Errorf("right expected {from:'V', value:0}, got {from:%q, value:%d}", right.from, right.value)
	}
}

func TestSplitByPrefix_NoMatch(t *testing.T) {
	c := Context{from: "ABC", value: 7}
	patterns := []Context{{from: "X", value: 10}, {from: "V", value: 5}}

	left, right := c.splitByPrefix(patterns)

	if left.from != "" || left.value != 0 {
		t.Errorf("left expected empty Context, got {from:%q, value:%d}", left.from, left.value)
	}
	if right.from != "ABC" || right.value != 7 {
		t.Errorf("right expected original Context, got {from:%q, value:%d}", right.from, right.value)
	}
}

func TestSplitByVinculum_Valid(t *testing.T) {
	c := Context{from: "(X)IV"}
	left, right, err := c.splitByVinculum()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if left.from != "X" {
		t.Errorf("left.from expected 'X', got %q", left.from)
	}
	if right.from != "IV" {
		t.Errorf("right.from expected 'IV', got %q", right.from)
	}
}

func TestSplitByVinculum_NoVinculum(t *testing.T) {
	c := Context{from: "XIV"}
	left, right, err := c.splitByVinculum()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if left.from != "" || left.value != 0 {
		t.Errorf("left expected empty Context, got {from:%q, value:%d}", left.from, left.value)
	}
	if right.from != "XIV" {
		t.Errorf("right.from expected 'XIV', got %q", right.from)
	}
}

func TestSplitByVinculum_Malformed(t *testing.T) {
	c := Context{from: "(XIV"}
	_, _, err := c.splitByVinculum()
	if err == nil {
		t.Errorf("expected error for malformed vinculum, got nil")
	}
}

func TestNewContextFromConcatenation(t *testing.T) {
	left := Context{from: "X", value: 10}
	right := Context{from: "IV", value: 4}

	res, err := NewContextFromConcatenation(left, right)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if res.from != "XIV" {
		t.Errorf("from expected 'XIV', got %q", res.from)
	}
	if res.value != 14 {
		t.Errorf("value expected 14, got %d", res.value)
	}
}

func TestNewContextFromConcatenation_Overflow(t *testing.T) {
	left := Context{from: "A", value: math.MaxUint64}
	right := Context{from: "B", value: 1}

	_, err := NewContextFromConcatenation(left, right)
	if err == nil {
		t.Errorf("expected overflow error, got nil")
	}
}

func TestNewContextFromVinculumConcatenation(t *testing.T) {
	vinc := Context{from: "X", value: 10} // 10 * 1000 = 10000
	plain := Context{from: "IV", value: 4}

	res, err := NewContextFromVinculumConcatenation(vinc, plain)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if res.value != 10004 {
		t.Errorf("value expected 10004, got %d", res.value)
	}
}

func TestNewContextFromVinculumConcatenation_Overflow(t *testing.T) {
	vinc := Context{from: "Z", value: (math.MaxUint64 / 1000) + 1}
	plain := Context{from: "I", value: 1}

	_, err := NewContextFromVinculumConcatenation(vinc, plain)
	if err == nil {
		t.Errorf("expected overflow error in MultiplySafe, got nil")
	}
}

func TestNewContextFromVinculumConcatenation_ZeroVinculum(t *testing.T) {
	vinc := Context{from: "", value: 0} // 0 * 1000 = 0 by MultiplySafe contract
	plain := Context{from: "V", value: 5}

	res, err := NewContextFromVinculumConcatenation(vinc, plain)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if res.value != 5 {
		t.Errorf("value expected 5, got %d", res.value)
	}
}
