package skin

import (
	"bytes"
	// Used to embed the default skin images as a variable
	_ "embed"
	"image"
	"image/draw"
	"image/png"
	"log"
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
	rawSteveData []byte

	//go:embed alex.png
	rawAlexData []byte

	steveSkin *image.NRGBA = image.NewNRGBA(image.Rect(0, 0, 64, 64))
	alexSkin  *image.NRGBA = image.NewNRGBA(image.Rect(0, 0, 64, 64))
)

func init() {
	{
		rawSteveSkin, err := png.Decode(bytes.NewReader(rawSteveData))

		if err != nil {
			log.Fatal(err)
		}

		draw.Draw(steveSkin, rawSteveSkin.Bounds(), rawSteveSkin, image.Pt(0, 0), draw.Src)
	}

	{
		rawAlexSkin, err := png.Decode(bytes.NewReader(rawAlexData))

		if err != nil {
			log.Fatal(err)
		}

		draw.Draw(alexSkin, rawAlexSkin.Bounds(), rawAlexSkin, image.Pt(0, 0), draw.Src)
	}
}

// IsOldSkin returns a boolean which will be true if the skin is a legacy skin, which contains missing information about the skin overlay.
func IsOldSkin(img *image.NRGBA) bool {
	return img.Bounds().Max.Y < 64
}

// IsSlimFromUUID returns whether the skin is a slim variant from the UUID.
// Credit: https://github.com/LapisBlue/Lapitar/blob/55ede80ce4ebb5ecc2b968164afb40f61b4cc509/mc/uuid.go#L23
func IsSlimFromUUID(uuid string) bool {
	uuid = strings.ReplaceAll(uuid, "-", "")

	return (isEven(uuid[7]) != isEven(uuid[16+7])) != (isEven(uuid[15]) != isEven(uuid[16+15]))
}

// GetDefaultSkin returns the default skin for either a regular or slim variant of a Minecraft skin.
func GetDefaultSkin(slim bool) *image.NRGBA {
	if slim {
		return alexSkin
	}

	return steveSkin
}

func extract(img *image.NRGBA, x, y, width, height int) *image.NRGBA {
	output := image.NewNRGBA(image.Rect(0, 0, width, height))

	draw.Draw(output, output.Bounds(), img, image.Pt(x, y), draw.Src)

	return output
}

func scale(img *image.NRGBA, scale int) *image.NRGBA {
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

func removeTransparency(img *image.NRGBA) *image.NRGBA {
	output := image.NewNRGBA(img.Bounds())

	for i, l := 0, len(img.Pix); i < l; i += 4 {
		output.Pix[i] = img.Pix[i]
		output.Pix[i+1] = img.Pix[i+1]
		output.Pix[i+2] = img.Pix[i+2]
		output.Pix[i+3] = 255
	}

	return output
}

func composite(bottom, top *image.NRGBA, x, y int) *image.NRGBA {
	output := image.NewNRGBA(bottom.Bounds())

	topBounds := top.Bounds().Max

	draw.Draw(output, bottom.Bounds(), bottom, image.Pt(0, 0), draw.Src)
	draw.Draw(output, image.Rect(0, 0, topBounds.X+x, topBounds.Y+y), top, image.Pt(-x, -y), draw.Over)

	return output
}

func flipHorizontal(img *image.NRGBA) *image.NRGBA {
	data := img.Pix
	bounds := img.Bounds()

	output := image.NewNRGBA(bounds)

	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			fx := bounds.Max.X - x - 1
			fi := fx*4 + y*4*bounds.Max.X
			ix := x*4 + y*4*bounds.Max.X

			for i := 0; i < 4; i++ {
				output.Pix[ix+i] = data[fi+i]
			}
		}
	}

	return output
}

func fixTransparency(img *image.NRGBA) *image.NRGBA {
	a := img.Pix[0:4]

	if a[3] == 0 {
		return img
	}

	output := clone(img)

	for i, l := 0, len(output.Pix); i < l; i += 4 {
		if output.Pix[i+0] != a[0] || output.Pix[i+1] != a[1] || output.Pix[i+2] != a[2] || output.Pix[i+3] != a[3] {
			continue
		}

		output.Pix[i+3] = 0
	}

	return output
}

func clone(img *image.NRGBA) *image.NRGBA {
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

func compositeTransform(bottom, top *image.NRGBA, mat matrix3, x, y float64) *image.NRGBA {
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
