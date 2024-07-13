package skin_test

import (
	"fmt"
	"image/png"
	"os"
	"testing"

	"github.com/mineatar-io/skin-render"
)

func TestRightBodySteve(t *testing.T) {
	rawSkin := skin.GetDefaultSkin(false)

	for i := 0; i <= 8; i++ {
		scale := 1 << i

		output := skin.RenderRightBody(rawSkin, skin.Options{
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
			f, err := os.OpenFile(fmt.Sprintf("rightbody_steve_test_%d.png", scale), os.O_CREATE|os.O_RDWR, 0777)

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

func BenchmarkRightBodySteve(b *testing.B) {
	rawSkin := skin.GetDefaultSkin(false)

	for n := 0; n <= b.N; n++ {
		skin.RenderRightBody(rawSkin, skin.Options{
			Scale:   defaultBenchmarkRenderScale,
			Overlay: true,
			Slim:    false,
			Square:  false,
		})
	}
}

func TestRightBodyAlex(t *testing.T) {
	rawSkin := skin.GetDefaultSkin(true)

	for i := 0; i <= 8; i++ {
		scale := 1 << i

		output := skin.RenderRightBody(rawSkin, skin.Options{
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
			f, err := os.OpenFile(fmt.Sprintf("rightbody_alex_test_%d.png", scale), os.O_CREATE|os.O_RDWR, 0777)

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

func BenchmarkRightBodyAlex(b *testing.B) {
	rawSkin := skin.GetDefaultSkin(true)

	for n := 0; n <= b.N; n++ {
		skin.RenderRightBody(rawSkin, skin.Options{
			Scale:   defaultBenchmarkRenderScale,
			Overlay: true,
			Slim:    true,
			Square:  false,
		})
	}
}

func TestRightBodySteveSquare(t *testing.T) {
	rawSkin := skin.GetDefaultSkin(false)

	for i := 0; i <= 8; i++ {
		scale := 1 << i

		output := skin.RenderRightBody(rawSkin, skin.Options{
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
			f, err := os.OpenFile(fmt.Sprintf("rightbody_steve_test_%d_square.png", scale), os.O_CREATE|os.O_RDWR, 0777)

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

func BenchmarkRightBodySteveSquare(b *testing.B) {
	rawSkin := skin.GetDefaultSkin(false)

	for n := 0; n <= b.N; n++ {
		skin.RenderRightBody(rawSkin, skin.Options{
			Scale:   defaultBenchmarkRenderScale,
			Overlay: true,
			Slim:    false,
			Square:  true,
		})
	}
}

func TestRightBodyAlexSquare(t *testing.T) {
	rawSkin := skin.GetDefaultSkin(true)

	for i := 0; i <= 8; i++ {
		scale := 1 << i

		output := skin.RenderRightBody(rawSkin, skin.Options{
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
			f, err := os.OpenFile(fmt.Sprintf("rightbody_alex_test_%d_square.png", scale), os.O_CREATE|os.O_RDWR, 0777)

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

func BenchmarkRightBodyAlexSquare(b *testing.B) {
	rawSkin := skin.GetDefaultSkin(true)

	for n := 0; n <= b.N; n++ {
		skin.RenderRightBody(rawSkin, skin.Options{
			Scale:   defaultBenchmarkRenderScale,
			Overlay: true,
			Slim:    true,
			Square:  true,
		})
	}
}
