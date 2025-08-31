package tools

import "fmt"

func NewWrappedError(input string, eCurr any, eNext any, err error) error {
	return fmt.Errorf("%T cannot process '%s' through %T: '%w'", eCurr, input, eNext, err)
}
