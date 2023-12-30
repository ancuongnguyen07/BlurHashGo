package blurhash

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"strings"

	"github.com/ancuongnguyen07/BlurHashGo/internal/base83"
	"github.com/ancuongnguyen07/BlurHashGo/internal/utils"
)

const (
	minComponents = 1
	maxComponents = 9
)

// encodeDC encodes DC components of RGB values
func encodeDC(r, g, b float64) int {
	return (utils.LinearTosRGB(r) << 16) + (utils.LinearTosRGB(g) << 8) + utils.LinearTosRGB(b)
}

// encodeAC encodes AC components of RGB values within the given maximum value
func encodeAC(r, g, b, maxVal float64) int {
	// The AC compoments of the DCT transform, ordered by increasing X first then Y. They are
	// encoded as three values for R,G and B, each between 0 and 18. They combined together
	// as R * 19^2 + G * 19 + B, for a total range of 0 and 6859
	quantR := math.Max(0, math.Min(18, math.Floor(utils.SignPow(r/maxVal, 0.5)*9+9.5)))
	quantG := math.Max(0, math.Min(18, math.Floor(utils.SignPow(g/maxVal, 0.5)*9+9.5)))
	quantB := math.Max(0, math.Min(18, math.Floor(utils.SignPow(b/maxVal, 0.5)*9+9.5)))
	return int(quantR*19*19 + quantG*19 + quantB)
}

func multiplyBasicFunction(xComps, yComps int, rgba image.Image) ([3]float64, error) {
	var r, g, b float64
	height := rgba.Bounds().Dy()
	width := rgba.Bounds().Dx()

	normalization := 2.0
	if xComps == 0 && yComps == 0 {
		normalization = 1.0
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// loop over the pixel at (x,y) and convert the color model
			// of that point into NRGB
			c, ok := color.NRGBAModel.Convert(rgba.At(x, y)).(color.NRGBA)
			if !ok {
				return [3]float64{0.0, 0.0, 0.0}, fmt.Errorf("invalid RGB color model")
			}
			basis := math.Cos(math.Pi*float64(xComps)*float64(x)/float64(width)) *
				math.Cos(math.Pi*float64(yComps)*float64(y)/float64(height))
			r += basis * utils.SRGBToLinear(int(c.R))
			g += basis * utils.SRGBToLinear(int(c.G))
			b += basis * utils.SRGBToLinear(int(c.B))
		}
	}

	scale := normalization / float64(width*height)

	return [3]float64{
		r * scale,
		g * scale,
		b * scale,
	}, nil
}

func Encode(xComps, yComps int, rgba image.Image) (string, error) {
	if xComps < minComponents || xComps > maxComponents {
		return "", ErrInvalidComps(xComps)
	}
	if yComps < minComponents || yComps > maxComponents {
		return "", ErrInvalidComps(yComps)
	}

	hashBuilder := strings.Builder{}

	sizeFlag := (xComps - 1) + (yComps-1)*9
	hashVal, err := base83.Encode(sizeFlag, 1)
	if err != nil {
		return "", err
	}
	_, err = hashBuilder.WriteString(hashVal)
	if err != nil {
		return "", err
	}

	// The matrix of RGB pixels corresponding to the image
	factors := make([][][3]float64, yComps)
	for y := 0; y < yComps; y++ {
		factors[y] = make([][3]float64, xComps)
		for x := 0; x < xComps; x++ {
			factor, err := multiplyBasicFunction(x, y, rgba)
			if err != nil {
				return "", err
			}

			factors[y][x][0] = factor[0]
			factors[y][x][1] = factor[1]
			factors[y][x][2] = factor[2]
		}
	}

	var maxVal float64
	if xComps*yComps-1 > 0 {
		actualMaxVal := 0.0
		for y := 0; y < yComps; y++ {
			for x := 0; x < xComps; x++ {
				if x == 0 && y == 0 {
					continue
				}
				f := factors[y][x]
				actualMaxVal = math.Max(math.Abs(f[0]), actualMaxVal)
				actualMaxVal = math.Max(math.Abs(f[1]), actualMaxVal)
				actualMaxVal = math.Max(math.Abs(f[2]), actualMaxVal)
			}
		}

		quantisedMaxVal := math.Max(0, math.Min(82, math.Floor(actualMaxVal*166-0.5)))
		maxVal = (quantisedMaxVal + 1) / 166
		hashVal, err = base83.Encode(int(quantisedMaxVal), 1)
		if err != nil {
			return "", err
		}
	} else {
		maxVal = 1.0
		hashVal, err = base83.Encode(0, 1)
		if err != nil {
			return "", err
		}
	}

	_, err = hashBuilder.WriteString(hashVal)
	if err != nil {
		return "", err
	}

	dc := factors[0][0]
	hashVal, err = base83.Encode(encodeDC(
		dc[0],
		dc[1],
		dc[2],
	), 4)
	if err != nil {
		return "", err
	}

	_, err = hashBuilder.WriteString(hashVal)
	if err != nil {
		return "", err
	}

	for y := 0; y < yComps; y++ {
		for x := 0; x < xComps; x++ {
			if x == 0 && y == 0 {
				continue
			}
			f := factors[y][x]
			hashVal, err := base83.Encode(encodeAC(
				f[0], f[1], f[2], maxVal,
			), 2)
			if err != nil {
				return "", err
			}
			_, err = hashBuilder.WriteString(hashVal)
			if err != nil {
				return "", err
			}
		}
	}

	return hashBuilder.String(), nil
}
