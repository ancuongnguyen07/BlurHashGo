package base83_test

import "github.com/ancuongnguyen07/BlurHashGo/internal/base83"

type result[T any] struct {
	out T
	err error
}

type base83EncodeTest[T any] struct {
	description string
	input       int
	length      int
	expected    result[T]
}

type base83DecodeTest[T any] struct {
	description string
	input       string
	expected    result[T]
}

// Test cases are referenced from:
// - https://github.com/buckket/go-blurhash/blob/master/base83/base83_test.go
var encodeTests = []base83EncodeTest[string]{
	{
		description: "zero input",
		input:       0,
		length:      0,
		expected: result[string]{
			out: "",
			err: nil,
		},
	},
	{
		description: "only digit input",
		input:       163902429697,
		length:      6,
		expected: result[string]{
			out: "foobar",
			err: nil,
		},
	},
	{
		description: "negative length",
		input:       123456789,
		length:      -3,
		expected: result[string]{
			out: "",
			err: base83.ErrInvalidLength(-3),
		},
	},
	{
		description: "percent symbol",
		input:       255172974336,
		length:      6,
		expected: result[string]{
			out: "%%%%%%",
			err: nil,
		},
	},
	{
		description: "random",
		input:       6869,
		length:      2,
		expected: result[string]{
			out: "~$",
			err: nil,
		},
	},
}

var decodeTests = []base83DecodeTest[int]{
	{
		description: "empty string",
		input:       "",
		expected: result[int]{
			out: 0,
			err: nil,
		},
	},
	{
		description: "only alphabet",
		input:       "foobar",
		expected: result[int]{
			out: 163902429697,
			err: nil,
		},
	},
	{
		description: "all valid characters",
		input:       "LFE.@D9F01_2%L%MIVD*9Goe-;WB",
		expected: result[int]{
			out: -1597651267176502418,
			err: nil,
		},
	},
	{
		description: "invalid character case",
		input:       "LFE.@D9F01_2%L%MIVD*9Goe-;Wµ",
		expected: result[int]{
			out: 0,
			err: base83.ErrInvalidCha('µ'),
		},
	},
	{
		description: "percent symbols",
		input:       "%%%%%%",
		expected: result[int]{
			out: 255172974336,
			err: nil,
		},
	},
	{
		description: "random",
		input:       "~$",
		expected: result[int]{
			out: 6869,
			err: nil,
		},
	},
}
