package skin

import (
	"bytes"
	"fmt"
	"math"

	// Used to embed the default skin images as a variable
	_ "embed"
	"image"
	"image/draw"
	"image/png"
	"strings"
)

var (
	//go:embed steve.png
	rawSteveSkinData []byte
	//go:embed alex.png
	rawAlexSkinData []byte
	steveSkin       *image.NRGBA = nil
	alexSkin        *image.NRGBA = nil
	zeroPoint       image.Point  = image.Point{}
)

func init() {
	{
		rawSteveSkin, err := png.Decode(bytes.NewReader(rawSteveSkinData))

		if err != nil {
			panic(err)
		}

		steveSkin = image.NewNRGBA(rawSteveSkin.Bounds())
		draw.Draw(steveSkin, rawSteveSkin.Bounds(), rawSteveSkin, image.Pt(0, 0), draw.Src)
	}

	{
		rawAlexSkin, err := png.Decode(bytes.NewReader(rawAlexSkinData))

		if err != nil {
			panic(err)
		}

		alexSkin = image.NewNRGBA(rawAlexSkin.Bounds())
		draw.Draw(alexSkin, rawAlexSkin.Bounds(), rawAlexSkin, image.Pt(0, 0), draw.Src)
	}
}

// IsOldSkin returns a boolean which will be true if the skin is a legacy skin, which contains missing information about the skin overlay except for the head.
func IsOldSkin(img image.Image) bool {
	return img.Bounds().Dy() < 64
}

// IsSlimFromUUID returns whether the skin is a slim variant from the UUID.
// Credit: https://github.com/LapisBlue/Lapitar/blob/55ede80ce4ebb5ecc2b968164afb40f61b4cc509/mc/uuid.go#L34
func IsSlimFromUUID(uuid string) bool {
	uuid = strings.ReplaceAll(uuid, "-", "")

	return (isEven(uuid[7]) != isEven(uuid[23])) != (isEven(uuid[15]) != isEven(uuid[31]))
}

// GetDefaultSkin returns the default skin for either a regular or slim variant of a Minecraft skin.
func GetDefaultSkin(slim bool) *image.NRGBA {
	if slim {
		return alexSkin
	}

	return steveSkin
}

func extract(img *image.NRGBA, r image.Rectangle) *image.NRGBA {
	output := image.NewNRGBA(image.Rect(0, 0, r.Dx(), r.Dy()))

	for x := r.Min.X; x < r.Max.X; x++ {
		for y := r.Min.Y; y < r.Max.Y; y++ {
			index := y*img.Stride + x*4
			inputColor := img.Pix[index : index+4]

			index = (y-r.Min.Y)*output.Stride + (x-r.Min.X)*4
			output.Pix[index] = inputColor[0]
			output.Pix[index+1] = inputColor[1]
			output.Pix[index+2] = inputColor[2]
			output.Pix[index+3] = inputColor[3]
		}
	}

	return output
}

func scale(img *image.NRGBA, scale int) *image.NRGBA {
	if scale < 2 {
		return img
	}

	bounds := img.Bounds().Size()
	output := image.NewNRGBA(image.Rect(0, 0, bounds.X*scale, bounds.Y*scale))

	for x := 0; x < bounds.X; x++ {
		for y := 0; y < bounds.Y; y++ {
			i := y*img.Stride + x*4
			color := img.Pix[i : i+4]

			for sx := 0; sx < scale; sx++ {
				for sy := 0; sy < scale; sy++ {
					i = (y*scale+sy)*output.Stride + (x*scale+sx)*4
					output.Pix[i] = color[0]
					output.Pix[i+1] = color[1]
					output.Pix[i+2] = color[2]
					output.Pix[i+3] = color[3]
				}
			}
		}
	}

	return output
}

func removeTransparency(img *image.NRGBA) *image.NRGBA {
	output := clone(img)

	for i, l := 0, len(output.Pix); i < l; i += 4 {
		output.Pix[i+3] = math.MaxUint8
	}

	return output
}

func composite(bottom, top *image.NRGBA, dx, dy int) *image.NRGBA {
	output := clone(bottom)

	outputBounds := output.Bounds()
	srcBounds := top.Bounds()

	for x := srcBounds.Min.X; x < srcBounds.Max.X; x++ {
		for y := srcBounds.Min.Y; y < srcBounds.Max.Y; y++ {
			if dx+x < outputBounds.Min.X || dy+y < outputBounds.Min.Y || dx+x >= outputBounds.Max.X || dy+y >= outputBounds.Max.Y {
				continue
			}

			index := y*top.Stride + x*4
			sourceColor := top.Pix[index : index+4]
			sourceAlpha := uint32(sourceColor[3]) * 0x101

			index = (dy+y)*output.Stride + (dx+x)*4
			outputColor := output.Pix[index : index+4]

			alphaOffset := ((1<<16 - 1) - sourceAlpha) * 0x101

			outputColor[0] = uint8((uint32(outputColor[0])*alphaOffset/(1<<16-1) + (uint32(sourceColor[0]) * sourceAlpha / 0xff)) >> 8)
			outputColor[1] = uint8((uint32(outputColor[1])*alphaOffset/(1<<16-1) + (uint32(sourceColor[1]) * sourceAlpha / 0xff)) >> 8)
			outputColor[2] = uint8((uint32(outputColor[2])*alphaOffset/(1<<16-1) + (uint32(sourceColor[2]) * sourceAlpha / 0xff)) >> 8)
			outputColor[3] = uint8((uint32(outputColor[3])*alphaOffset/(1<<16-1) + sourceAlpha) >> 8)
		}
	}

	return output
}

func flipHorizontal(src *image.NRGBA) *image.NRGBA {
	bounds := src.Bounds()
	output := image.NewNRGBA(bounds)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			index := y*src.Stride + x*4
			inputColor := src.Pix[index : index+4]

			index = y*output.Stride + (bounds.Max.X-x-1)*4
			output.Pix[index] = inputColor[0]
			output.Pix[index+1] = inputColor[1]
			output.Pix[index+2] = inputColor[2]
			output.Pix[index+3] = inputColor[3]
		}
	}

	return output
}

func flipVertically(src *image.NRGBA) *image.NRGBA {
	bounds := src.Bounds()
	output := image.NewNRGBA(bounds)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			index := y*src.Stride + x*4
			inputColor := src.Pix[index : index+4]

			index = (bounds.Max.Y-y-1)*output.Stride + x*4
			output.Pix[index] = inputColor[0]
			output.Pix[index+1] = inputColor[1]
			output.Pix[index+2] = inputColor[2]
			output.Pix[index+3] = inputColor[3]
		}
	}

	return output
}

func fixTransparency(img *image.NRGBA) *image.NRGBA {
	checkColor := img.Pix[0:4]

	if checkColor[3] == 0 {
		return img
	}

	output := clone(img)

	for i, l := 0, len(output.Pix); i < l; i += 4 {
		if !isEqualSlice(checkColor, output.Pix[i:i+4]) {
			continue
		}

		output.Pix[i+3] = 0
	}

	return output
}

func clone(img *image.NRGBA) *image.NRGBA {
	bounds := img.Bounds()
	output := image.NewNRGBA(bounds)

	draw.Draw(output, bounds, img, zeroPoint, draw.Src)

	return output
}

func getSlimOffset(slim bool) int {
	if slim {
		return 1
	}

	return 0
}

// This function is a whole mess of code that I do not want to touch, but it
// seems to work very well. Most of this code was influenced by code in the
// `go/x/image` package, but with a lot less redundancy. The color mixing
// code was taken from the built-in Go method draw.Draw() from the
// `image/draw` package.
func compositeTransform(dst, src *image.NRGBA, m matrix2x2, outputX, outputY float64) *image.NRGBA {
	sourceBounds := src.Bounds()

	output := clone(dst)
	outputBounds := output.Bounds()

	im := m.Inverse()
	dr := transformRect(m, src.Bounds())
	dox, doy := translateCoordinatesWithMatrix(outputX, outputY, m)

	for boundX := dr.Min.X; boundX < dr.Max.X; boundX++ {
		for boundY := dr.Min.Y; boundY < dr.Max.Y; boundY++ {
			outputX, outputY := boundX+int(dox), boundY+int(doy)

			if outputX < outputBounds.Min.X || outputY < outputBounds.Min.Y || outputX >= outputBounds.Max.X || outputY >= outputBounds.Max.Y {
				continue
			}

			sourceX, sourceY := translateCoordinatesWithMatrix(float64(boundX), float64(boundY), im)

			if int(sourceX) < sourceBounds.Min.X || int(sourceY) < sourceBounds.Min.Y || int(sourceX) >= sourceBounds.Max.X || int(sourceY) >= sourceBounds.Max.Y {
				continue
			}

			index := int(sourceY)*src.Stride + int(sourceX)*4
			sourceColor := src.Pix[index : index+4]
			sourceAlpha := uint32(sourceColor[3]) * 0x101

			index = outputY*output.Stride + outputX*4
			outputColor := output.Pix[index : index+4]

			alphaOffset := ((1<<16 - 1) - sourceAlpha) * 0x101

			outputColor[0] = uint8((uint32(outputColor[0])*alphaOffset/(1<<16-1) + (uint32(sourceColor[0]) * sourceAlpha / 0xff)) >> 8)
			outputColor[1] = uint8((uint32(outputColor[1])*alphaOffset/(1<<16-1) + (uint32(sourceColor[1]) * sourceAlpha / 0xff)) >> 8)
			outputColor[2] = uint8((uint32(outputColor[2])*alphaOffset/(1<<16-1) + (uint32(sourceColor[2]) * sourceAlpha / 0xff)) >> 8)
			outputColor[3] = uint8((uint32(outputColor[3])*alphaOffset/(1<<16-1) + sourceAlpha) >> 8)
		}
	}

	return output
}

func rotate90(img *image.NRGBA) *image.NRGBA {
	bounds := img.Bounds().Size()
	output := image.NewNRGBA(image.Rect(0, 0, bounds.Y, bounds.X))

	for x := 0; x < bounds.X; x++ {
		for y := 0; y < bounds.Y; y++ {
			index := y*img.Stride + x*4
			inputColor := img.Pix[index : index+4]

			index = int(x)*output.Stride + int(y)*4 // Intentionally flipped X and Y
			output.Pix[index] = inputColor[0]
			output.Pix[index+1] = inputColor[1]
			output.Pix[index+2] = inputColor[2]
			output.Pix[index+3] = inputColor[3]
		}
	}

	return output
}

func rotate270(img *image.NRGBA) *image.NRGBA {
	bounds := img.Bounds().Size()
	output := image.NewNRGBA(image.Rect(0, 0, bounds.Y, bounds.X))

	for x := 0; x < bounds.X; x++ {
		for y := 0; y < bounds.Y; y++ {
			index := y*img.Stride + x*4
			inputColor := img.Pix[index : index+4]

			index = (bounds.X-int(x)-1)*output.Stride + int(y)*4 // Intentionally flipped X and Y
			output.Pix[index] = inputColor[0]
			output.Pix[index+1] = inputColor[1]
			output.Pix[index+2] = inputColor[2]
			output.Pix[index+3] = inputColor[3]
		}
	}

	return output
}

// Credit: https://github.com/LapisBlue/Lapitar/blob/55ede80ce4ebb5ecc2b968164afb40f61b4cc509/mc/uuid.go#L23
func isEven(c uint8) bool {
	switch {
	case c >= '0' && c <= '9':
		return (c & 1) == 0
	case c >= 'a' && c <= 'f':
		return (c & 1) == 1
	default:
		panic(fmt.Errorf("invalid character: %c", c))
	}
}

func isEqualSlice[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	for i, l := 0, len(a); i < l; i++ {
		if a[i] == b[i] {
			continue
		}

		return false
	}

	return true
}
