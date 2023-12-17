package base83

import (
	"fmt"
)

// ErrInvalidLength is retuned when the length is too short or less than or equal 0
type ErrInvalidLength int

func (e ErrInvalidLength) Error() string {
	return fmt.Sprintf("base83 output: invalid length (%d) error", int(e))
}

// ErrInvalidCha is retuned when the given character is not in the list of base83 characters
type ErrInvalidCha rune

func (e ErrInvalidCha) Error() string {
	return fmt.Sprintf("invalid base83 character: %c", rune(e))
}
