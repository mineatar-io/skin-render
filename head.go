package skin

import "image"

// RenderHead renders a 3-dimensional image of the head of a Minecraft player's skin.
func RenderHead(skin *image.NRGBA, opts Options) *image.NRGBA {
	scaleDouble := float64(opts.Scale)
	output := image.NewNRGBA(image.Rect(0, 0, 16*opts.Scale, 19*opts.Scale-int(scaleDouble/2.0)-1))

	var (
		frontHead *image.NRGBA = removeTransparency(extract(skin, 8, 8, 8, 8))
		topHead   *image.NRGBA = removeTransparency(extract(skin, 8, 0, 8, 8))
		rightHead *image.NRGBA = removeTransparency(extract(skin, 0, 8, 8, 8))
	)

	if opts.Overlay && !IsOldSkin(skin) {
		overlaySkin := fixTransparency(skin)

		frontHead = composite(frontHead, extract(overlaySkin, 40, 8, 8, 8), 0, 0)
		topHead = composite(topHead, extract(overlaySkin, 40, 0, 8, 8), 0, 0)
		rightHead = composite(rightHead, extract(overlaySkin, 32, 8, 8, 8), 0, 0)
	}

	// Front Head
	output = compositeTransform(output, scale(frontHead, opts.Scale), transformForward, 8*scaleDouble, 12*scaleDouble-1)

	// Top Head
	output = compositeTransform(output, scale(topHead, opts.Scale), transformUp, -4*scaleDouble, 4*scaleDouble)

	// Right Head
	output = compositeTransform(output, scale(rightHead, opts.Scale), transformRight, 0, 4*scaleDouble)

	return output
}
