package skin

import (
	"image"
	"math"
)

// RenderHead renders a 3-dimensional image of the head of a Minecraft player's skin.
func RenderHead(skin *image.NRGBA, opts Options) *image.NRGBA {
	scaleDouble := float64(opts.Scale)

	var output *image.NRGBA = image.NewNRGBA(image.Rect(0, 0, 13*opts.Scale+int(math.Floor(scaleDouble*0.855)), 16*opts.Scale))

	// Front Head
	output = compositeTransform(output, scale(removeTransparency(extract(skin, HeadFront)), opts.Scale), frontMatrix, 8*scaleDouble, 12*scaleDouble)

	// Top Head
	output = compositeTransform(output, scale(removeTransparency(extract(skin, HeadTop)), opts.Scale), plantMatrix, 4*scaleDouble, -4*scaleDouble)

	// Right Head
	output = compositeTransform(output, scale(removeTransparency(extract(skin, HeadRight)), opts.Scale), sideMatrix, 0, 4*scaleDouble)

	if opts.Overlay {
		overlaySkin := fixTransparency(skin)

		// Front Head Overlay
		output = compositeTransform(output, scale(extract(overlaySkin, HeadOverlayFront), opts.Scale), frontMatrix, 8*scaleDouble, 12*scaleDouble)

		// Top Head Overlay
		output = compositeTransform(output, scale(extract(overlaySkin, HeadOverlayTop), opts.Scale), plantMatrix, 4*scaleDouble, -4*scaleDouble)

		// Right Head Overlay
		output = compositeTransform(output, scale(extract(overlaySkin, HeadOverlayRight), opts.Scale), sideMatrix, 0, 4*scaleDouble)
	}

	return output
}
