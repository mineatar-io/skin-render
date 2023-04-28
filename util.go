package skin

import (
	"bytes"
	"math"

	// Used to embed the default skin images as a variable
	_ "embed"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"strings"

	drw "golang.org/x/image/draw"
	"golang.org/x/image/math/f64"
)

var (
	skewA            float64 = 26.0 / 45.0
	skewB            float64 = skewA * 2.0
	transformForward matrix3 = matrix3{
		XX: 1, YX: -skewA,
		XY: 0, YY: skewB,
		X0: 0, Y0: skewA,
	}
	transformUp matrix3 = matrix3{
		XX: 1, YX: -skewA,
		XY: 1, YY: skewA,
		X0: 0, Y0: 0,
	}
	transformRight matrix3 = matrix3{
		XX: 1, YX: skewA,
		XY: 0, YY: skewB,
		X0: 0, Y0: 0,
	}

	//go:embed steve.png
	rawSteveSkinData []byte

	//go:embed alex.png
	rawAlexSkinData []byte

	steveSkin image.Image = nil
	alexSkin  image.Image = nil
)

func init() {
	var err error

	if steveSkin, err = png.Decode(bytes.NewReader(rawSteveSkinData)); err != nil {
		panic(err)
	}

	if alexSkin, err = png.Decode(bytes.NewReader(rawAlexSkinData)); err != nil {
		panic(err)
	}
}

// IsOldSkin returns a boolean which will be true if the skin is a legacy skin, which contains missing information about the skin overlay.
func IsOldSkin(img image.Image) bool {
	return img.Bounds().Max.Y < 64
}

// IsSlimFromUUID returns whether the skin is a slim variant from the UUID.
// Credit: https://github.com/LapisBlue/Lapitar/blob/55ede80ce4ebb5ecc2b968164afb40f61b4cc509/mc/uuid.go#L23
func IsSlimFromUUID(uuid string) bool {
	uuid = strings.ReplaceAll(uuid, "-", "")

	return (isEven(uuid[7]) != isEven(uuid[16+7])) != (isEven(uuid[15]) != isEven(uuid[16+15]))
}

// GetDefaultSkin returns the default skin for either a regular or slim variant of a Minecraft skin.
func GetDefaultSkin(slim bool) image.Image {
	if slim {
		return alexSkin
	}

	return steveSkin
}

func extract(img image.Image, x, y, width, height int) image.Image {
	output := image.NewNRGBA(image.Rect(0, 0, width, height))

	draw.Draw(output, output.Bounds(), img, image.Pt(x, y), draw.Src)

	return output
}

func scale(img image.Image, scale int) image.Image {
	if scale == 1 {
		return img
	}

	bounds := img.Bounds().Max
	output := image.NewNRGBA(image.Rect(0, 0, bounds.X*scale, bounds.Y*scale))

	for x := 0; x < bounds.X; x++ {
		for y := 0; y < bounds.Y; y++ {
			color := img.At(x, y)

			for sx := 0; sx < scale; sx++ {
				for sy := 0; sy < scale; sy++ {
					output.Set(x*scale+sx, y*scale+sy, color)
				}
			}
		}
	}

	return output
}

func removeTransparency(img image.Image) image.Image {
	output := image.NewNRGBA(img.Bounds())
	bounds := img.Bounds().Size()

	for x := 0; x < bounds.X; x++ {
		for y := 0; y < bounds.Y; y++ {
			c := img.At(x, y)

			r, g, b, _ := c.RGBA()

			output.Set(x, y, color.NRGBA64{
				R: uint16(r),
				G: uint16(g),
				B: uint16(b),
				A: math.MaxUint16,
			})
		}
	}

	return output
}

func composite(bottom, top image.Image, x, y int) image.Image {
	output := image.NewNRGBA(bottom.Bounds())

	topBounds := top.Bounds().Max

	draw.Draw(output, bottom.Bounds(), bottom, image.Pt(0, 0), draw.Src)
	draw.Draw(output, image.Rect(0, 0, topBounds.X+x, topBounds.Y+y), top, image.Pt(-x, -y), draw.Over)

	return output
}

func flipHorizontal(img image.Image) image.Image {
	output := image.NewNRGBA(img.Bounds())
	bounds := img.Bounds().Size()

	for x := 0; x < bounds.X; x++ {
		for y := 0; y < bounds.Y; y++ {
			output.Set(bounds.X-x-1, y, img.At(x, y))
		}
	}

	return output
}

func fixTransparency(img image.Image) image.Image {
	ir, ig, ib, ia := img.At(0, 0).RGBA()

	if ia == 0 {
		return img
	}

	output := clone(img)
	bounds := output.Bounds().Size()

	for x := 0; x < bounds.X; x++ {
		for y := 0; y < bounds.Y; y++ {
			pr, pg, pb, pa := img.At(x, y).RGBA()

			if pr != ir || pg != ig || pb != ib || pa != ia {
				continue
			}

			output.Set(x, y, color.NRGBA64{
				R: uint16(pr),
				G: uint16(pg),
				B: uint16(pb),
				A: 0,
			})
		}
	}

	return output
}

func clone(img image.Image) *image.NRGBA {
	bounds := img.Bounds()
	output := image.NewNRGBA(bounds)

	draw.Draw(output, bounds, img, image.Pt(0, 0), draw.Src)

	return output
}

func getSlimOffset(slim bool) int {
	if slim {
		return 1
	}

	return 0
}

func compositeTransform(bottom, top image.Image, mat matrix3, x, y float64) image.Image {
	output := image.NewNRGBA(bottom.Bounds())

	draw.Draw(output, bottom.Bounds(), bottom, image.Pt(0, 0), draw.Src)

	transformer := drw.NearestNeighbor

	fx, fy := float64(x), float64(y)

	m := mat.Translate(fx, fy)

	transformer.Transform(output, f64.Aff3{m.XX, m.XY, m.X0, m.YX, m.YY, m.Y0}, top, top.Bounds(), draw.Over, nil)

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
		panic("Invalid digit " + string(c))
	}
}
