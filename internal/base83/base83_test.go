package base83_test

import (
	"testing"

	"github.com/ancuongnguyen07/BlurHashGo/internal/base83"
)

func TestEncode(t *testing.T) {
	for _, tc := range encodeTests {
		t.Run(tc.description, func(t *testing.T) {
			out, err := base83.Encode(tc.input, tc.length)
			if err != tc.expected.err {
				t.Fatal(err)
			}
			if out != tc.expected.out {
				t.Fatalf("Encode(%d,%d) = %s, want: %s", tc.input, tc.length, out, tc.expected.out)
			}
		})
	}
}

func BenchmarkEncode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = base83.Encode(1234543, 7)
	}
}

func TestDecode(t *testing.T) {
	for _, tc := range decodeTests {
		t.Run(tc.description, func(t *testing.T) {
			out, err := base83.Decode(tc.input)
			if err != tc.expected.err {
				t.Fatal(err)
			}
			if out != tc.expected.out {
				t.Fatalf("Decode(%s) = %d, want: %d", tc.input, out, tc.expected.out)
			}
		})
	}
}

func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = base83.Decode("Hello World!")
	}
}
