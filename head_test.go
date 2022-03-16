package skin_test

import (
	"image/png"
	"os"
	"testing"

	"github.com/mineatar-io/skin-render"
)

func TestHeadSteve(t *testing.T) {
	rawSkin := skin.GetDefaultSkin(false)

	output := skin.RenderHead(rawSkin, skin.Options{
		Scale:   16,
		Overlay: true,
		Slim:    false,
	})

	f, err := os.OpenFile("head_steve_test.png", os.O_CREATE|os.O_RDWR, 0777)

	if err != nil {
		t.Fatal(err)
	}

	defer f.Close()

	if err = png.Encode(f, output); err != nil {
		t.Fatal(err)
	}
}

func TestHeadAlex(t *testing.T) {
	rawSkin := skin.GetDefaultSkin(true)

	output := skin.RenderHead(rawSkin, skin.Options{
		Scale:   16,
		Overlay: true,
		Slim:    true,
	})

	f, err := os.OpenFile("head_alex_test.png", os.O_CREATE|os.O_RDWR, 0777)

	if err != nil {
		t.Fatal(err)
	}

	defer f.Close()

	if err = png.Encode(f, output); err != nil {
		t.Fatal(err)
	}
}
