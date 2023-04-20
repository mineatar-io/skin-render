package skin

import "image"

// RenderFace renders a 2-dimensional image of the face of a Minecraft player's skin.
func RenderFace(skin *image.NRGBA, opts Options) *image.NRGBA {
	output := removeTransparency(extract(skin, 8, 8, 8, 8))

	if opts.Overlay && !IsOldSkin(skin) {
		output = composite(output, extract(skin, 40, 8, 8, 8), 0, 0)
	}

	return scale(output, opts.Scale)
}
