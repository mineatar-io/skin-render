package skin_test

import (
	"fmt"
	"image/png"
	"os"
	"testing"

	"github.com/mineatar-io/skin-render"
)

func TestLeftBodySteve(t *testing.T) {
	rawSkin := skin.GetDefaultSkin(false)

	for i := 0; i <= 8; i++ {
		scale := 1 << i

		output := skin.RenderLeftBody(rawSkin, skin.Options{
			Scale:   scale,
			Overlay: true,
			Slim:    false,
			Square:  false,
		})

		if output.Bounds().Dx() < 1 {
			t.Fatalf("result image width is %d pixels\n", output.Bounds().Dx())
		}

		if output.Bounds().Dy() < 1 {
			t.Fatalf("result image height is %d pixels\n", output.Bounds().Dy())
		}

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
			Square:  false,
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
			Square:  false,
		})

		if output.Bounds().Dx() < 1 {
			t.Fatalf("result image width is %d pixels\n", output.Bounds().Dx())
		}

		if output.Bounds().Dy() < 1 {
			t.Fatalf("result image height is %d pixels\n", output.Bounds().Dy())
		}

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
			Square:  false,
		})
	}
}

func TestLeftBodySteveSquare(t *testing.T) {
	rawSkin := skin.GetDefaultSkin(false)

	for i := 0; i <= 8; i++ {
		scale := 1 << i

		output := skin.RenderLeftBody(rawSkin, skin.Options{
			Scale:   scale,
			Overlay: true,
			Slim:    false,
			Square:  true,
		})

		if output.Bounds().Dx() < 1 {
			t.Fatalf("result image width is %d pixels\n", output.Bounds().Dx())
		}

		if output.Bounds().Dy() < 1 {
			t.Fatalf("result image height is %d pixels\n", output.Bounds().Dy())
		}

		if output.Bounds().Size().X != output.Bounds().Size().Y {
			t.Fatalf("result image is not square (%s)\n", output.Bounds().Size())
		}

		if writeRenders {
			f, err := os.OpenFile(fmt.Sprintf("leftbody_steve_test_%d_square.png", scale), os.O_CREATE|os.O_RDWR, 0777)

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

func BenchmarkLeftBodySteveSquare(b *testing.B) {
	rawSkin := skin.GetDefaultSkin(false)

	for n := 0; n <= b.N; n++ {
		skin.RenderLeftBody(rawSkin, skin.Options{
			Scale:   defaultBenchmarkRenderScale,
			Overlay: true,
			Slim:    false,
			Square:  true,
		})
	}
}

func TestLeftBodyAlexSquare(t *testing.T) {
	rawSkin := skin.GetDefaultSkin(true)

	for i := 0; i <= 8; i++ {
		scale := 1 << i

		output := skin.RenderLeftBody(rawSkin, skin.Options{
			Scale:   scale,
			Overlay: true,
			Slim:    true,
			Square:  true,
		})

		if output.Bounds().Dx() < 1 {
			t.Fatalf("result image width is %d pixels\n", output.Bounds().Dx())
		}

		if output.Bounds().Dy() < 1 {
			t.Fatalf("result image height is %d pixels\n", output.Bounds().Dy())
		}

		if output.Bounds().Size().X != output.Bounds().Size().Y {
			t.Fatalf("result image is not square (%s)\n", output.Bounds().Size())
		}

		if writeRenders {
			f, err := os.OpenFile(fmt.Sprintf("leftbody_alex_test_%d_square.png", scale), os.O_CREATE|os.O_RDWR, 0777)

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

func BenchmarkLeftBodyAlexSquare(b *testing.B) {
	rawSkin := skin.GetDefaultSkin(true)

	for n := 0; n <= b.N; n++ {
		skin.RenderLeftBody(rawSkin, skin.Options{
			Scale:   defaultBenchmarkRenderScale,
			Overlay: true,
			Slim:    true,
			Square:  true,
		})
	}
}
