package blurhash

import (
	"image"
	"image/color"
	"math"

	"github.com/ancuongnguyen07/BlurHashGo/internal/base83"
	"github.com/ancuongnguyen07/BlurHashGo/internal/utils"
)

func decodeDC(val int) [3]float64 {
	return [3]float64{
		utils.SRGBToLinear(val >> 16),        // R value
		utils.SRGBToLinear((val >> 8) & 255), // G value
		utils.SRGBToLinear(val & 255),        // B value
	}
}

func decodeAC(val int, maxVal float64) [3]float64 {
	// Each value represents a float-point value between -1 and 1.
	// 0-8 represent negative values, 9 represents 0, zand 10-18 represents
	// positive values. Positive values are encoded as
	// ((X-9) / 9) ^ 2, while negative values are encoded as
	// -((9-X) / 9) ^ 2. This value is then multiplied by the maximum
	// AC component value
	quantR := int(math.Floor(float64(val) / (19 * 19)))
	quantG := math.Mod(math.Floor(float64(val)/19), 19)
	quantB := math.Mod(float64(val), 19)

	return [3]float64{
		utils.SignPow((float64(quantR)-9)/9, 2.0) * maxVal, // R value
		utils.SignPow((float64(quantG)-9)/9, 2.0) * maxVal, // G value
		utils.SignPow((float64(quantB)-9)/9, 2.0) * maxVal, // B value
	}
}

// Components decodes and returns the x/y component of the given hash string
func Components(blurhash string) (numX, numY int, err error) {
	hashLength := len(blurhash)
	if len(blurhash) < 6 {
		return 0, 0, ErrInvalidHashLength(hashLength)
	}

	// Number of components
	// For a BlurHash with nx components along the X axis and ny components
	// along the Y axis, this is equal to (nx-1) + (ny-1) * 9
	sizeFlag, err := base83.Decode(string(blurhash[0]))
	if err != nil {
		return 0, 0, err
	}
	numY = (sizeFlag / 9) + 1
	numX = (sizeFlag % 9) + 1

	if hashLength != 4+2*numX*numY {
		return 0, 0, ErrInvalidHashLength(hashLength)
	}
	return numX, numY, nil
}

// Decode decodes the given hash string into the blur image.
// `punch` argument is the factor indicating how much the contrast is.
func Decode(blurhash string, width, height, punch int) (image.Image, error) {
	numX, numY, err := Components(blurhash)
	if err != nil {
		return nil, err
	}

	if punch < 1 {
		punch = 1
	}

	quantizedMaxVal, err := base83.Decode(string(blurhash[1]))
	if err != nil {
		return nil, err
	}
	// All AC components are scaled by this value. It represents a
	// floating-point value of (max+1) / 166
	maxVal := float64(quantizedMaxVal+1) / 166

	colors := make([][3]float64, numX*numY)
	for i := range colors {
		if i == 0 {
			intVal, err := base83.Decode(blurhash[2:6])
			if err != nil {
				return nil, err
			}
			colors[i] = decodeDC(intVal)
		} else {
			intVal, err := base83.Decode(blurhash[4+i*2 : 6+i*2])
			if err != nil {
				return nil, err
			}
			colors[i] = decodeAC(intVal, maxVal*float64(punch))
		}
	}

	imgNRGBA := image.NewNRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			var r, g, b float64
			for j := 0; j < numY; j++ {
				for i := 0; i < numX; i++ {
					basis := math.Cos((math.Pi*float64(x)*float64(i))/float64(width)) * math.Cos((math.Pi*float64(y)*float64(j))/float64(height))
					idx := i + j*numX
					r += colors[idx][0] * basis
					g += colors[idx][1] * basis
					b += colors[idx][2] * basis
				}
			}

			intR := utils.LinearTosRGB(r)
			intG := utils.LinearTosRGB(g)
			intB := utils.LinearTosRGB(b)

			imgNRGBA.SetNRGBA(x, y, color.NRGBA{uint8(intR), uint8(intG), uint8(intB), 255})
		}
	}

	return imgNRGBA, nil
}
