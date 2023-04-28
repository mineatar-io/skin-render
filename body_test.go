package skin_test

import (
	"fmt"
	"image/png"
	"os"
	"testing"

	"github.com/mineatar-io/skin-render"
)

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
