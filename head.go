package skin

import (
	"image"
	"math"
)

// RenderHead renders a 3-dimensional image of the head of a Minecraft player's skin.
func RenderHead(skin *image.NRGBA, opts Options) *image.NRGBA {
	scaleDouble := float64(opts.Scale)

	var output *image.NRGBA = image.NewNRGBA(image.Rect(0, 0, 13*opts.Scale+int(math.Ceil(scaleDouble*0.86603)), 16*opts.Scale))

	var (
		frontHead *image.NRGBA = removeTransparency(extract(skin, HeadFront))
		topHead   *image.NRGBA = removeTransparency(extract(skin, HeadTop))
		rightHead *image.NRGBA = removeTransparency(extract(skin, HeadRight))
	)

	// Front Head
	output = compositeTransform(output, scale(frontHead, opts.Scale), frontMatrix, 8*scaleDouble, 12*scaleDouble-1)

	// Top Head
	output = compositeTransform(output, scale(topHead, opts.Scale), plantMatrix, 4*scaleDouble, -4*scaleDouble)

	// Right Head
	output = compositeTransform(output, scale(rightHead, opts.Scale), sideMatrix, 0, 4*scaleDouble)

	if opts.Overlay {
		overlaySkin := fixTransparency(skin)

		// Front Head Overlay
		output = compositeTransform(output, scale(extract(overlaySkin, HeadOverlayFront), opts.Scale), frontMatrix, 8*scaleDouble, 12*scaleDouble-1)

		// Top Head Overlay
		output = compositeTransform(output, scale(extract(overlaySkin, HeadOverlayTop), opts.Scale), plantMatrix, 4*scaleDouble, -4*scaleDouble)

		// Right Head Overlay
		output = compositeTransform(output, scale(extract(overlaySkin, HeadOverlayRight), opts.Scale), sideMatrix, 0, 4*scaleDouble)
	}

	return output
}
