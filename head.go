package skin

import (
	"image"
	"math"
)

// RenderHead renders a 3-dimensional image of the head of a Minecraft player's skin.
func RenderHead(skin *image.NRGBA, opts Options) *image.NRGBA {
	scaleDouble := float64(opts.Scale)

	var (
		frontHead *image.NRGBA = removeTransparency(extract(skin, HeadFront))
		topHead   *image.NRGBA = removeTransparency(extract(skin, HeadTop))
		rightHead *image.NRGBA = removeTransparency(extract(skin, HeadRight))
	)

	if opts.Overlay {
		overlaySkin := fixTransparency(skin)

		frontHead = composite(frontHead, extract(overlaySkin, HeadOverlayFront), 0, 0)
		topHead = composite(topHead, extract(overlaySkin, HeadOverlayTop), 0, 0)
		rightHead = composite(rightHead, extract(overlaySkin, HeadOverlayRight), 0, 0)
	}

	var output *image.NRGBA = image.NewNRGBA(image.Rect(0, 0, 13*opts.Scale+int(math.Floor(scaleDouble*0.855)), 16*opts.Scale))

	// Front Head
	output = compositeTransform(output, scale(frontHead, opts.Scale), frontMatrix, 8*scaleDouble, 12*scaleDouble)

	// Top Head
	output = compositeTransform(output, scale(topHead, opts.Scale), plantMatrix, 4*scaleDouble, -4*scaleDouble)

	// Right Head
	output = compositeTransform(output, scale(rightHead, opts.Scale), sideMatrix, 0, 4*scaleDouble)

	return output
}
