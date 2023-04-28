package skin

import (
	"image"
)

// RenderLeftBody renders a 2-dimensional image of the left side of a Minecraft player's skin.
func RenderLeftBody(skin image.Image, opts Options) image.Image {
	slimOffset := getSlimOffset(opts.Slim)

	var (
		leftHead    image.Image = removeTransparency(extract(skin, 16, 8, 8, 8))
		leftLeftArm image.Image = nil
		leftLeftLeg image.Image = nil
	)

	if IsOldSkin(skin) {
		leftLeftArm = flipHorizontal(removeTransparency(extract(skin, 40, 20, 4, 12)))
		leftLeftLeg = flipHorizontal(removeTransparency(extract(skin, 0, 20, 4, 12)))
	} else {
		leftLeftArm = removeTransparency(extract(skin, 40-slimOffset, 52, 4, 12))
		leftLeftLeg = removeTransparency(extract(skin, 24, 52, 4, 12))

		if opts.Overlay {
			overlaySkin := fixTransparency(skin)

			leftHead = composite(leftHead, extract(overlaySkin, 48, 8, 8, 8), 0, 0)
			leftLeftArm = composite(leftLeftArm, extract(overlaySkin, 56-slimOffset, 52, 4, 12), 0, 0)
			leftLeftLeg = composite(leftLeftLeg, extract(overlaySkin, 8, 52, 4, 12), 0, 0)
		}
	}

	var output image.Image = image.NewNRGBA(image.Rect(0, 0, 8, 32))

	// Left Head
	output = composite(output, leftHead, 0, 0)

	// Left Arm
	output = composite(output, leftLeftArm, 2, 8)

	// Left Leg
	output = composite(output, leftLeftLeg, 2, 20)

	return scale(output, opts.Scale)
}
