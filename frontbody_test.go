package skin_test

import (
	"fmt"
	"image/png"
	"os"
	"testing"

	"github.com/mineatar-io/skin-render"
)

func TestFrontBodySteve(t *testing.T) {
	rawSkin := skin.GetDefaultSkin(false)

	for i := 0; i <= 8; i++ {
		scale := 1 << i

		output := skin.RenderFrontBody(rawSkin, skin.Options{
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
			f, err := os.OpenFile(fmt.Sprintf("frontbody_steve_test_%d.png", scale), os.O_CREATE|os.O_RDWR, 0777)

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

func BenchmarkFrontBodySteve(b *testing.B) {
	rawSkin := skin.GetDefaultSkin(false)

	for n := 0; n <= b.N; n++ {
		skin.RenderFrontBody(rawSkin, skin.Options{
			Scale:   defaultBenchmarkRenderScale,
			Overlay: true,
			Slim:    false,
			Square:  false,
		})
	}
}

func TestFrontBodyAlex(t *testing.T) {
	rawSkin := skin.GetDefaultSkin(true)

	for i := 0; i <= 8; i++ {
		scale := 1 << i

		output := skin.RenderFrontBody(rawSkin, skin.Options{
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
			f, err := os.OpenFile(fmt.Sprintf("frontbody_alex_test_%d.png", scale), os.O_CREATE|os.O_RDWR, 0777)

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

func BenchmarkFrontBodyAlex(b *testing.B) {
	rawSkin := skin.GetDefaultSkin(true)

	for n := 0; n <= b.N; n++ {
		skin.RenderFrontBody(rawSkin, skin.Options{
			Scale:   defaultBenchmarkRenderScale,
			Overlay: true,
			Slim:    true,
			Square:  false,
		})
	}
}

func TestFrontBodySteveSquare(t *testing.T) {
	rawSkin := skin.GetDefaultSkin(false)

	for i := 0; i <= 8; i++ {
		scale := 1 << i

		output := skin.RenderFrontBody(rawSkin, skin.Options{
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
			f, err := os.OpenFile(fmt.Sprintf("frontbody_steve_test_%d_square.png", scale), os.O_CREATE|os.O_RDWR, 0777)

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

func BenchmarkFrontBodySteveSquare(b *testing.B) {
	rawSkin := skin.GetDefaultSkin(false)

	for n := 0; n <= b.N; n++ {
		skin.RenderFrontBody(rawSkin, skin.Options{
			Scale:   defaultBenchmarkRenderScale,
			Overlay: true,
			Slim:    false,
			Square:  true,
		})
	}
}

func TestFrontBodyAlexSquare(t *testing.T) {
	rawSkin := skin.GetDefaultSkin(true)

	for i := 0; i <= 8; i++ {
		scale := 1 << i

		output := skin.RenderFrontBody(rawSkin, skin.Options{
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
			f, err := os.OpenFile(fmt.Sprintf("frontbody_alex_test_%d_square.png", scale), os.O_CREATE|os.O_RDWR, 0777)

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

func BenchmarkFrontBodyAlexSquare(b *testing.B) {
	rawSkin := skin.GetDefaultSkin(true)

	for n := 0; n <= b.N; n++ {
		skin.RenderFrontBody(rawSkin, skin.Options{
			Scale:   defaultBenchmarkRenderScale,
			Overlay: true,
			Slim:    true,
			Square:  true,
		})
	}
}
