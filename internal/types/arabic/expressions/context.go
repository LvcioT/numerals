package expressions

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"

	"taurino.com/numerals/tools"
)

type Context struct {
	value uint64
	to    string
}

func (c Context) GetFrom() string {
	return strconv.Itoa(int(c.value))
}

func (c Context) GetValue() uint64 {
	return c.value
}

func (c Context) GetTo() string {
	return c.to
}

// splitByPrefix splits the Context based on matching prefixes from the provided list of Contexts (patterns).
// Returns a left Context that matches the prefix and a right Context with the remaining value.
// If no prefix matches, the left Context is empty, and the right Context contains the full original value.
func (c Context) splitByPrefix(patterns []Context) (Context, Context) {
	// sorts prefixes by value descending, it is necessary to match the biggest first and avoid partial matches
	slices.SortFunc(patterns, func(a, b Context) int {
		return cmp.Compare(b.value, a.value)
	})

	for _, pattern := range patterns {
		prefix := pattern.to
		value := pattern.value

		if c.value >= value {
			left := Context{to: prefix, value: value}
			right := Context{value: c.value - value}

			return left, right
		}
	}

	// no prefix matched, move all the content value the right
	return Context{}, c
}

func NewContextFromValue(value uint64) Context {
	return Context{value: value}
}

func NewContextFromAdd(left Context, right Context) (Context, error) {
	value, err := tools.AddSafe(left.value, right.value)
	if err != nil {
		return Context{}, fmt.Errorf("value value big for %d+%d: '%w'", left.value, right.value, err)
	}

	return Context{to: left.to + right.to, value: value}, nil
}

func NewContextFromThousandMultiply(thousand Context, base Context) (Context, error) {
	valueThousand, err := tools.MultiplySafe(thousand.value, 1_000)
	if err != nil {
		return Context{}, fmt.Errorf("value value big for %d as thousand: '%w'", thousand.value, err)
	}

	var toThousand string

	if thousand.value > 1 {
		toThousand = fmt.Sprintf("(%s)", thousand.to)
	} else {
		toThousand = ""
	}

	leftBase := Context{
		to:    toThousand,
		value: valueThousand,
	}

	return NewContextFromAdd(leftBase, base)
}
