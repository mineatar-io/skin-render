package skin

import "image"

func RenderFace(skin *image.NRGBA, opts Options) *image.NRGBA {
	output := removeTransparency(extract(skin, 8, 8, 8, 8))

	if opts.Overlay && !IsOldSkin(skin) {
		output = composite(output, extract(skin, 40, 8, 8, 8), 0, 0)
	}

	return scale(output, opts.Scale)
}
