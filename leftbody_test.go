package skin_test

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
	"testing"

	"github.com/mineatar-io/skin-render"
)

func TestCustomSkin(t *testing.T) {
	f, err := os.Open("skin.png")

	if err != nil {
		t.Fatal(err)
	}

	rawImg, err := png.Decode(f)

	if err != nil {
		t.Fatal(err)
	}

	if err = f.Close(); err != nil {
		t.Fatal(err)
	}

	rawSkin := image.NewNRGBA(rawImg.Bounds())
	draw.Draw(rawSkin, rawImg.Bounds(), rawImg, image.ZP, draw.Src)

	res := skin.RenderLeftBody(rawSkin, skin.Options{
		Scale:   8,
		Overlay: true,
		Slim:    false,
	})

	f, err = os.OpenFile("output.png", os.O_CREATE|os.O_RDWR, 0777)

	if err != nil {
		t.Fatal(err)
	}

	if err = png.Encode(f, res); err != nil {
		t.Fatal(err)
	}

	if err = f.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestLeftBodySteve(t *testing.T) {
	rawSkin := skin.GetDefaultSkin(false)

	for i := 0; i <= 8; i++ {
		scale := 1 << i

		output := skin.RenderLeftBody(rawSkin, skin.Options{
			Scale:   scale,
			Overlay: true,
			Slim:    false,
		})

		if writeRenders {
			f, err := os.OpenFile(fmt.Sprintf("leftbody_steve_test_%d.png", scale), os.O_CREATE|os.O_RDWR, 0777)

			if err != nil {
				t.Fatal(err)
			}

			if err = png.Encode(f, output); err != nil {
				t.Fatal(err)
			}

			if err = f.Close(); err != nil {
				t.Fatal(err)
			}
		}
	}
}

func BenchmarkLeftBodySteve(b *testing.B) {
	rawSkin := skin.GetDefaultSkin(false)

	for n := 0; n <= b.N; n++ {
		skin.RenderLeftBody(rawSkin, skin.Options{
			Scale:   defaultBenchmarkRenderScale,
			Overlay: true,
			Slim:    false,
		})
	}
}

func TestLeftBodyAlex(t *testing.T) {
	rawSkin := skin.GetDefaultSkin(true)

	for i := 0; i <= 8; i++ {
		scale := 1 << i

		output := skin.RenderLeftBody(rawSkin, skin.Options{
			Scale:   scale,
			Overlay: true,
			Slim:    true,
		})

		if writeRenders {
			f, err := os.OpenFile(fmt.Sprintf("leftbody_alex_test_%d.png", scale), os.O_CREATE|os.O_RDWR, 0777)

			if err != nil {
				t.Fatal(err)
			}

			if err = png.Encode(f, output); err != nil {
				t.Fatal(err)
			}

			if err = f.Close(); err != nil {
				t.Fatal(err)
			}
		}
	}
}

func BenchmarkLeftBodyAlex(b *testing.B) {
	rawSkin := skin.GetDefaultSkin(true)

	for n := 0; n <= b.N; n++ {
		skin.RenderLeftBody(rawSkin, skin.Options{
			Scale:   defaultBenchmarkRenderScale,
			Overlay: true,
			Slim:    true,
		})
	}
}
