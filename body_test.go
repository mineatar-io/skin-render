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
	f, err := os.Open("test.png")

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

	img := image.NewNRGBA(rawImg.Bounds())
	draw.Draw(img, rawImg.Bounds(), rawImg, image.ZP, draw.Src)

	output := skin.RenderHead(img, skin.Options{
		Scale:   8,
		Overlay: true,
		Slim:    false,
	})

	f, err = os.OpenFile("test-body.png", os.O_CREATE|os.O_RDWR, 0777)

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

func TestFullBodySteve(t *testing.T) {
	rawSkin := skin.GetDefaultSkin(false)

	for i := 0; i <= 8; i++ {
		scale := 1 << i

		output := skin.RenderBody(rawSkin, skin.Options{
			Scale:   scale,
			Overlay: true,
			Slim:    false,
		})

		f, err := os.OpenFile(fmt.Sprintf("fullbody_steve_test_%d.png", scale), os.O_CREATE|os.O_RDWR, 0777)

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

func TestFullBodyAlex(t *testing.T) {
	rawSkin := skin.GetDefaultSkin(true)

	for i := 0; i <= 8; i++ {
		scale := 1 << i

		output := skin.RenderBody(rawSkin, skin.Options{
			Scale:   scale,
			Overlay: true,
			Slim:    true,
		})

		f, err := os.OpenFile(fmt.Sprintf("fullbody_alex_test_%d.png", scale), os.O_CREATE|os.O_RDWR, 0777)

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
