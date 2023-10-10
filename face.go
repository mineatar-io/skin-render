package skin

import "image"

// RenderFace renders a 2-dimensional image of the face of a Minecraft player's skin.
func RenderFace(img *image.NRGBA, opts Options) *image.NRGBA {
	if err := validateSkin(img); err != nil {
		panic(err)
	}

	var (
		skin   *image.NRGBA = convertToNRGBA(img)
		output *image.NRGBA = removeTransparency(extract(skin, HeadFront))
	)

	if opts.Overlay {
		overlaySkin := fixTransparency(skin)

		composite(output, extract(overlaySkin, HeadOverlayFront), 0, 0)
	}

	return scale(output, opts.Scale)
}
