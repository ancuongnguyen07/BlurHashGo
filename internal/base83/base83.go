package base83

import (
	"math"
	"strings"
)

// The list of valid characters in base83
const characters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz#$%*+,-.:;=?@[]^_{|}~"

// Encode encodes the given integer `value` into a base83 string
// with the given `length`. If `length` is too short, an
func Encode(value, length int) (string, error) {
	if length < 0 {
		return "", ErrInvalidLength(length)
	} else if length == 0 {
		return "", nil
	}
	encodedValue := make([]rune, length)
	divisor := int(math.Pow(83, float64(length-1)))
	for i := 0; i < length; i++ {
		digit := (value / divisor) % 83
		divisor /= 83
		encodedValue[i] = rune(characters[digit])
	}
	return string(encodedValue), nil
}

// Decode decodes the given base83 string into a integer value
func Decode(str string) (value int, err error) {
	for _, letter := range str {
		base83Index := strings.IndexRune(characters, letter)
		if base83Index == -1 {
			// invalid base83 character
			return 0, ErrInvalidCha(letter)
		}
		value = value*len(characters) + base83Index
	}
	return value, nil
}
