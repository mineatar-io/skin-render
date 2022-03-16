package skin_test

import (
	"image/png"
	"os"
	"testing"

	"github.com/mineatar-io/skin-render"
)

func TestFrontBodySteve(t *testing.T) {
	rawSkin := skin.GetDefaultSkin(false)

	output := skin.RenderFrontBody(rawSkin, skin.Options{
		Scale:   16,
		Overlay: true,
		Slim:    false,
	})

	f, err := os.OpenFile("frontbody_steve_test.png", os.O_CREATE|os.O_RDWR, 0777)

	if err != nil {
		t.Fatal(err)
	}

	defer f.Close()

	if err = png.Encode(f, output); err != nil {
		t.Fatal(err)
	}
}

func TestFrontBodyAlex(t *testing.T) {
	rawSkin := skin.GetDefaultSkin(true)

	output := skin.RenderFrontBody(rawSkin, skin.Options{
		Scale:   16,
		Overlay: true,
		Slim:    true,
	})

	f, err := os.OpenFile("frontbody_alex_test.png", os.O_CREATE|os.O_RDWR, 0777)

	if err != nil {
		t.Fatal(err)
	}

	defer f.Close()

	if err = png.Encode(f, output); err != nil {
		t.Fatal(err)
	}
}
