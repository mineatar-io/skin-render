package skin

import "image"

// RenderFace renders a 2-dimensional image of the face of a Minecraft player's skin.
func RenderFace(skin *image.NRGBA, opts Options) *image.NRGBA {
	output := removeTransparency(extract(skin, HeadFront))

	if opts.Overlay {
		overlaySkin := fixTransparency(skin)

		output = composite(output, extract(overlaySkin, HeadOverlayFront), 0, 0)
	}

	return scale(output, opts.Scale)
}
