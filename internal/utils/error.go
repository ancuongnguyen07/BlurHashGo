package utils

import "fmt"

// ErrEmptyFile occurs when a file which was opened by I/O is empty
type ErrEmptyFile string

func (e ErrEmptyFile) Error() string {
	return fmt.Sprintf("empty file: %s", string(e))
}

// ErrUnsupportedFile occurs when a file other than PNG or JPG is read
// to do blurhash encode
type ErrUnsupportedFile string

func (e ErrUnsupportedFile) Error() string {
	return fmt.Sprintf("unsupported file: %s", string(e))
}

// ErrUnknownFileType occurs when the type of a specific file
// could not be detected
type ErrUnknownFileType string

func (e ErrUnknownFileType) Error() string {
	return fmt.Sprintf("unknown file type: %s", string(e))
}
