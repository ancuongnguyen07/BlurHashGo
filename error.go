package blurhash

import "fmt"

// ErrInvalidComps occurs when the x-component or y-component
// is not in the range of 1 and 9 exclusively
type ErrInvalidComps int

func (e ErrInvalidComps) Error() string {
	return fmt.Sprintf("Invalid axis component: %d", int(e))
}
