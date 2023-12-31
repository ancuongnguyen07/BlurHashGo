package blurhash_test

import (
	"testing"

	"github.com/ancuongnguyen07/BlurHashGo/blurhash"
)

type result struct {
	xcomp int
	ycomp int
	err   error
}

func TestComponent(t *testing.T) {
	testCases := []struct {
		description string
		hash        string
		expected    result
	}{
		{
			description: "Invalid hash",
			hash:        "LOD+VQIU0mWU^MNHI[WYa0bISh",
			expected: result{
				xcomp: 0,
				ycomp: 0,
				err:   blurhash.ErrInvalidHashLength(26),
			},
		},
		{
			description: "Too short hash",
			hash:        "He42",
			expected: result{
				xcomp: 0,
				ycomp: 0,
				err:   blurhash.ErrInvalidHashLength(4),
			},
		},
		{
			description: "Valid hash",
			hash:        "LOD+VQIU0mWU^MNHI[WYa0bIShj[",
			expected: result{
				xcomp: 4,
				ycomp: 3,
				err:   nil,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			x, y, err := blurhash.Components(tc.hash)
			if err != tc.expected.err {
				t.Fatalf("mismatch error | wanted %v | got %v", tc.expected.err, err)
			}
			if x != tc.expected.xcomp || y != tc.expected.ycomp {

				t.Fatalf("components are mismatched | wanted (%d,%d) | got (%d,%d)", tc.expected.xcomp, tc.expected.ycomp, x, y)
			}
		})
	}
}

func BenchmarkDecode(b *testing.B) {
	const hash = "LOD+VQIU0mWU^MNHI[WYa0bISh["
	for i := 0; i < b.N; i++ {
		_, _ = blurhash.Decode(hash, 300, 300, 2)
	}
}
