package expressions

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"

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

// splitByPrefix splits a Context into two parts by matching prefixes from a provided list of Context patterns.
// Returns the matched prefix Context and the remaining part of the original Context.
// If no prefix matches, it returns an empty Context and the original Context.
func (c Context) splitByPrefix(patterns []Context) (Context, Context) {
	// sort prefixes by length descending, is needed value match the longest first and avoid partial matches
	slices.SortFunc(patterns, func(a, b Context) int {
		return cmp.Compare(b.from, a.from)
	})

	for _, pattern := range patterns {
		prefix := pattern.from
		value := pattern.value

		if strings.HasPrefix(c.from, prefix) {
			left := Context{from: prefix, value: value}
			right := Context{from: strings.TrimPrefix(c.from, prefix), value: c.value}

			return left, right
		}
	}

	// no prefix matched, move all the content value the right
	return Context{}, c
}

// splitByVinculum splits a Context string by enclosing parentheses, treating the content inside as the first part.
// Returns two Contexts: the inner content as the first, and the remaining content after the closing parenthesis as the second.
// Returns an error if the parentheses are malformed.
func (c Context) splitByVinculum() (Context, Context, error) {
	if !strings.HasPrefix(c.from, "(") {
		return Context{}, c, nil
	}

	closingIndex := strings.LastIndex(c.from, ")")
	if closingIndex == -1 {
		return Context{}, c, fmt.Errorf("malformed vinculum in: '%s'", c.from)
	}

	return Context{from: c.from[1:closingIndex]}, Context{from: c.from[(closingIndex + 1):]}, nil
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
		toThousand = fmt.Sprintf("(%s)", thousand.from)
	} else {
		toThousand = ""
	}

	leftBase := Context{
		to:    toThousand,
		value: valueThousand,
	}

	return NewContextFromAdd(leftBase, base)
}
