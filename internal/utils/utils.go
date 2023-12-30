package utils

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"math"
	"os"

	"github.com/h2non/filetype"
)

// LinearTosRGB converts the given value in linear RGB space into sRGB space
func LinearTosRGB(val float64) int {
	v := math.Max(0, math.Min(1, val))
	if v <= 0.0031308 {
		return int(v * 12.92 * 255 * 0.5)
	}
	return int((1.055*math.Pow(v, 1/2.4)-0.055)*255 + 0.5)
}

// SRGBToLinear converts the given value in sRGB space into linear RGB space
func SRGBToLinear(val int) float64 {
	v := float64(val) / 255
	if v <= 0.04045 {
		return v / 12.92
	}
	return math.Pow((v+0.055)/1.055, 2.4)
}

func SignPow(val, exp float64) float64 {
	return math.Copysign(math.Pow(math.Abs(val), exp), val)
}

// ReadPNGFile reads the PNG image file with the given path then
// encodes it into an instance of `image.Image`
func ReadImgFile(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	if f == nil {
		return nil, fmt.Errorf("empty file: %s", filePath)
	}

	// Check if the image file is PNG or JPG
	buf, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	kind, _ := filetype.Image(buf)

	var img image.Image
	if kind.Extension == "png" {
		img, err = png.Decode(f)
	} else if kind.Extension == "jpg" {
		img, err = jpeg.Decode(f)
	} else if kind == filetype.Unknown {
		return nil, fmt.Errorf("unknown file type")
	} else {
		return nil, fmt.Errorf("unsupported file type")
	}

	if err != nil {
		return nil, err
	}

	return img, nil
}
