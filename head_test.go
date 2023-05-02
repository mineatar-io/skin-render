package skin_test

import (
	"fmt"
	"image/png"
	"os"
	"testing"

	"github.com/mineatar-io/skin-render"
)

func TestHeadSteve(t *testing.T) {
	rawSkin := skin.GetDefaultSkin(false)

	for i := 0; i <= 8; i++ {
		scale := 1 << i

		output := skin.RenderHead(rawSkin, skin.Options{
			Scale:   scale,
			Overlay: true,
			Slim:    false,
		})

		if writeRenders {
			f, err := os.OpenFile(fmt.Sprintf("head_steve_test_%d.png", scale), os.O_CREATE|os.O_RDWR, 0777)

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

func BenchmarkHeadSteve(b *testing.B) {
	rawSkin := skin.GetDefaultSkin(false)

	for n := 0; n <= b.N; n++ {
		skin.RenderHead(rawSkin, skin.Options{
			Scale:   defaultBenchmarkRenderScale,
			Overlay: true,
			Slim:    false,
		})
	}
}

func TestHeadAlex(t *testing.T) {
	rawSkin := skin.GetDefaultSkin(true)

	for i := 0; i <= 8; i++ {
		scale := 1 << i

		output := skin.RenderHead(rawSkin, skin.Options{
			Scale:   scale,
			Overlay: true,
			Slim:    true,
		})

		if writeRenders {
			f, err := os.OpenFile(fmt.Sprintf("head_alex_test_%d.png", scale), os.O_CREATE|os.O_RDWR, 0777)

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

func BenchmarkHeadAlex(b *testing.B) {
	rawSkin := skin.GetDefaultSkin(true)

	for n := 0; n <= b.N; n++ {
		skin.RenderHead(rawSkin, skin.Options{
			Scale:   defaultBenchmarkRenderScale,
			Overlay: true,
			Slim:    true,
		})
	}
}
