package expressions

import (
	"math"
	"testing"
)

func TestContextGetters(t *testing.T) {
	c := Context{to: "X", value: 42}

	if got := c.GetFrom(); got != "42" {
		t.Errorf("GetFrom expected '42', got %q", got)
	}
	if got := c.GetValue(); got != 42 {
		t.Errorf("GetValue expected 42, got %d", got)
	}
	if got := c.GetTo(); got != "X" {
		t.Errorf("GetTo expected 'X', got %q", got)
	}
}

func TestNewContextFromValue(t *testing.T) {
	c := NewContextFromValue(123)
	if c.value != 123 {
		t.Errorf("value expected 123, got %d", c.value)
	}
	if c.to != "" {
		t.Errorf("to expected empty, got %q", c.to)
	}
}

func TestSplitByPrefix_LongestMatchByValue(t *testing.T) {
	c := Context{value: 14}
	patterns := []Context{{to: "X", value: 10}, {to: "XI", value: 11}}

	left, right := c.splitByPrefix(patterns)

	if left.to != "XI" || left.value != 11 {
		t.Errorf("left expected {to:'XI', value:11}, got {to:%q, value:%d}", left.to, left.value)
	}
	if right.to != "" || right.value != 3 {
		t.Errorf("right expected {to:'', value:3}, got {to:%q, value:%d}", right.to, right.value)
	}
}

func TestSplitByPrefix_NoMatch(t *testing.T) {
	c := Context{value: 4, to: "abc"}
	patterns := []Context{{to: "V", value: 5}, {to: "X", value: 10}}

	left, right := c.splitByPrefix(patterns)

	if left.to != "" || left.value != 0 {
		t.Errorf("left expected empty Context, got {to:%q, value:%d}", left.to, left.value)
	}
	if right.to != "abc" || right.value != 4 {
		t.Errorf("right expected original Context, got {to:%q, value:%d}", right.to, right.value)
	}
}

func TestNewContextFromAdd_Normal(t *testing.T) {
	left := Context{to: "A", value: 2}
	right := Context{to: "B", value: 3}

	res, err := NewContextFromAdd(left, right)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res.to != "AB" {
		t.Errorf("to expected 'AB', got %q", res.to)
	}
	if res.value != 5 {
		t.Errorf("value expected 5, got %d", res.value)
	}
}

func TestNewContextFromAdd_Overflow(t *testing.T) {
	left := Context{to: "A", value: math.MaxUint64}
	right := Context{to: "B", value: 1}

	_, err := NewContextFromAdd(left, right)
	if err == nil {
		t.Errorf("expected overflow error, got nil")
	}
}

func TestNewContextFromThousandMultiply_Normal(t *testing.T) {
	thousand := Context{to: "a", value: 2}
	base := Context{to: "b", value: 3}

	res, err := NewContextFromThousandMultiply(thousand, base)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res.value != 2003 {
		t.Errorf("value expected 2003, got %d", res.value)
	}
	if res.to != "(a)b" {
		t.Errorf("to expected '(2)b', got %q", res.to)
	}
}

func TestNewContextFromThousandMultiply_ThousandOne(t *testing.T) {
	thousand := Context{value: 1}
	base := Context{to: "b", value: 3}

	res, err := NewContextFromThousandMultiply(thousand, base)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res.value != 1003 {
		t.Errorf("value expected 1003, got %d", res.value)
	}
	if res.to != "b" {
		t.Errorf("to expected 'b', got %q", res.to)
	}
}

func TestNewContextFromThousandMultiply_ZeroThousand(t *testing.T) {
	thousand := Context{value: 0}
	base := Context{to: "b", value: 5}

	res, err := NewContextFromThousandMultiply(thousand, base)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res.value != 5 {
		t.Errorf("value expected 5, got %d", res.value)
	}
	if res.to != "b" {
		t.Errorf("to expected 'b', got %q", res.to)
	}
}

func TestNewContextFromThousandMultiply_Overflow(t *testing.T) {
	thousand := Context{value: (math.MaxUint64 / 1000) + 1}
	base := Context{to: "b", value: 1}

	_, err := NewContextFromThousandMultiply(thousand, base)
	if err == nil {
		t.Errorf("expected overflow error in MultiplySafe, got nil")
	}
}
