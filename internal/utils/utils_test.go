package utils_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/ancuongnguyen07/BlurHashGo/internal/utils"
)

func TestLinearToRGB(t *testing.T) {
	for _, tc := range linearTosRGBTestCases {
		t.Run(fmt.Sprintf("Input: %f", tc.in), func(t *testing.T) {
			rgb := utils.LinearTosRGB(tc.in)
			if rgb != tc.out {
				t.Fatalf("wanted %d, got %d", tc.out, rgb)
			}
		})
	}
}

func TestSRGBToLinear(t *testing.T) {
	for _, tc := range sRGBToLinearTestCases {
		t.Run(fmt.Sprintf("Input: %d", tc.in), func(t *testing.T) {
			linear := utils.SRGBToLinear(tc.in)
			if math.Abs(linear-tc.out) > 0.05 {
				t.Fatalf("wanted %f, got %f", tc.out, linear)
			}
		})
	}
}

func TestSignPow(t *testing.T) {
	for _, tc := range signPowTestCases {
		t.Run(fmt.Sprintf("Input val: %f | Exp: %f", tc.in, tc.exp), func(t *testing.T) {
			signPow := utils.SignPow(tc.in, tc.exp)
			if signPow != tc.out {
				t.Fatalf("wanted %f, got %f", tc.out, signPow)
			}
		})
	}
}
