package blurhash_test

import (
	"image"
	"image/png"
	"os"
	"testing"

	"github.com/ancuongnguyen07/BlurHashGo/blurhash"
	"github.com/ancuongnguyen07/BlurHashGo/internal/utils"
)

func TestEncodeImgFile(t *testing.T) {
	testCases := []struct {
		description  string
		filePath     string
		xComp, yComp int
		expected     string
	}{
		{
			description: "Ice Cream",
			filePath:    "../imgs/ice-cream.png",
			xComp:       4,
			yComp:       3,
			expected:    "LORnoQT}}gs6?aw@M{N0xBN2J=xt",
		},
		{
			description: "Pizza",
			filePath:    "../imgs/pizza.png",
			xComp:       4,
			yComp:       3,
			expected:    "LVGtcy~9IqoLxYo3afWB0#EMs.WC",
		},
		{
			description: "Salad",
			filePath:    "../imgs/salad.png",
			xComp:       4,
			yComp:       3,
			expected:    "LOD+VQIU0mWU^MNHI[WYa0bIShj[",
		},
		{
			// reference from https://github.com/bbrks/go-blurhash/tree/master
			description: "Dall-E",
			filePath:    "../imgs/dalle.png",
			xComp:       5,
			yComp:       5,
			expected:    "eaF#5R0#WBjYR+58-nWCWBn~bIsTbbayjFWof8jFj[WX-nNHR*jss.",
		},
		{
			// reference from https://github.com/bbrks/go-blurhash/tree/master
			description: "Octocat",
			filePath:    "../imgs/octocat.png",
			xComp:       4,
			yComp:       3,
			expected:    "LNAdApj[00aymkj[TKay9}ay-Sj[",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			img, err := utils.ReadImgFile(tc.filePath)
			if err != nil {
				t.Fatalf("failed to read PNG file: %v", err)
			}
			hash, err := blurhash.Encode(tc.xComp, tc.yComp, img)
			if err != nil {
				t.Fatalf("failed to encode img into blurhash")
			}
			if hash != tc.expected {
				t.Fatalf("invalid hash output | wanted %s | got %s", tc.expected, hash)
			}

		})
	}
}

func TestWrongComponents(t *testing.T) {
	testCases := []struct {
		description string
		xComp       int
		yComp       int
	}{
		{
			description: "too small x-component",
			xComp:       0,
			yComp:       4,
		},
		{
			description: "too small y-component",
			xComp:       3,
			yComp:       -4,
		},
		{
			description: "too big x-component",
			xComp:       20,
			yComp:       7,
		},
		{
			description: "too big y-component",
			xComp:       6,
			yComp:       12,
		},
	}

	fooImg := image.NewNRGBA(image.Rect(0, 0, 120, 120))

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			_, err := blurhash.Encode(tc.xComp, tc.yComp, fooImg)
			if err == nil {
				t.Fatalf("invalid component pair x-y (%d,%d) should raise error", tc.xComp, tc.yComp)
			}
			err, ok := err.(blurhash.ErrInvalidComps)
			if !ok {
				t.Fatalf("error type ErrInvalidComps should be through | got %v", err)
			}
		})
	}
}

func BenchmarkEncode(b *testing.B) {
	filePath := "imgs/salad.png"
	f, err := os.Open(filePath)
	if err != nil {
		b.Fatalf("failed to open %s", filePath)
	}
	defer f.Close()

	img, err := png.Decode(f)
	if err != nil {
		b.Fatal("failed to decode the image")
	}

	for i := 0; i < b.N; i++ {
		_, _ = blurhash.Encode(4, 3, img)
	}
}
