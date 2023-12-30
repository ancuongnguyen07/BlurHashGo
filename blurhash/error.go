package blurhash

import "fmt"

// ErrInvalidComps occurs when the x-component or y-component
// is not in the range of 1 and 9 exclusively
type ErrInvalidComps int

func (e ErrInvalidComps) Error() string {
	return fmt.Sprintf("invalid axis component: %d", int(e))
}

// ErrInvalidHashLength occurs when the length of blurhash is less
// than 6 or not equal (4 + 2 * numX * numY), where numX and numY
// are the number of x/y-compoments respectively
type ErrInvalidHashLength int

func (e ErrInvalidHashLength) Error() string {
	return fmt.Sprintf("invalid hash length: %d", int(e))
}
