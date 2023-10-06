package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
	"time"

	"github.com/jessevdk/go-flags"
	"github.com/mineatar-io/skin-render"
)

var (
	inputFile string
	opts      Options = Options{}
)

type Options struct {
	Type      string `short:"t" long:"type" description:"The type of image to render" default:"body"`
	Scale     int    `short:"s" long:"scale" description:"The scale of the rendered output image" default:"16"`
	NoOverlay bool   `short:"O" long:"no-overlay" description:"Disables the overlay layer of the resulting image"`
	Slim      bool   `short:"S" long:"slim" description:"Enable this option if the input skin image is slim"`
	Output    string `short:"o" long:"output" description:"The file to write the output image to" default:"output.png"`
	Verbose   bool   `short:"V" long:"verbose" description:"Prints extra debug information"`
}

func init() {
	args, err := flags.Parse(&opts)

	if err != nil {
		if flags.WroteHelp(err) {
			os.Exit(0)

			return
		}

		panic(err)
	}

	if len(args) < 1 {
		fmt.Println("missing input file argument")

		os.Exit(1)
	}

	inputFile = args[0]
}

func readInputFile() (*image.NRGBA, error) {
	f, err := os.Open(inputFile)

	if err != nil {
		return nil, err
	}

	img, err := png.Decode(f)

	if err != nil {
		return nil, err
	}

	if err = f.Close(); err != nil {
		return nil, err
	}

	output := image.NewNRGBA(img.Bounds())
	draw.Draw(output, img.Bounds(), img, image.Pt(0, 0), draw.Src)

	return output, nil
}

func writeOutputImage(fileName string, img image.Image) error {
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 0777)

	if err != nil {
		return err
	}

	if err = png.Encode(f, img); err != nil {
		return err
	}

	return f.Close()
}

func main() {
	if opts.Verbose {
		fmt.Printf(
			"Using options:\n - Input: %s\n - Output: %s\n - Scale: %d\n - Overlay: %t\n - Slim: %t\n",
			inputFile,
			opts.Output,
			opts.Scale,
			!opts.NoOverlay,
			opts.Slim,
		)
	}

	t := time.Now()

	skinImage, err := readInputFile()

	if err != nil {
		panic(err)
	}

	if opts.Verbose {
		fmt.Printf("Read input image (%s)\n", time.Since(t).Round(time.Microsecond))
	}

	var (
		result  *image.NRGBA
		options skin.Options = skin.Options{
			Scale:   opts.Scale,
			Overlay: !opts.NoOverlay,
			Slim:    opts.Slim,
		}
	)

	t = time.Now()

	switch opts.Type {
	case "face":
		{
			result = skin.RenderFace(skinImage, options)

			break
		}
	case "head":
		{
			result = skin.RenderHead(skinImage, options)

			break
		}
	case "body":
		{
			result = skin.RenderBody(skinImage, options)

			break
		}
	case "front":
		{
			result = skin.RenderFrontBody(skinImage, options)

			break
		}
	case "back":
		{
			result = skin.RenderBackBody(skinImage, options)

			break
		}
	case "left":
		{
			result = skin.RenderLeftBody(skinImage, options)

			break
		}
	case "right":
		{
			result = skin.RenderRightBody(skinImage, options)

			break
		}
	default:
		{
			fmt.Printf("unknown --type value: %s\n", opts.Type)

			os.Exit(1)
		}
	}

	if opts.Verbose {
		fmt.Printf("Generated result image (%s)\n", time.Since(t).Round(time.Microsecond))
	}

	t = time.Now()

	if err = writeOutputImage(opts.Output, result); err != nil {
		panic(err)
	}

	if opts.Verbose {
		fmt.Printf("Wrote output image (%s)\n", time.Since(t).Round(time.Microsecond))
	}
}
