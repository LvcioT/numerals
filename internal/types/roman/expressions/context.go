package expressions

import "strconv"

type Context struct {
	from string
	to   uint64
}

func (c Context) GetFrom() (string, error) {
	return c.from, nil
}

func (c Context) GetValue() (uint64, error) {
	return c.to, nil
}

func (c Context) GetTo() (string, error) {
	return strconv.Itoa(int(c.to)), nil
}

func NexContextFromString(input string) Context {
	return Context{from: input}
}
