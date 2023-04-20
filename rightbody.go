package skin

import (
	"image"
)

// RenderRightBody renders a 2-dimensional image of the right side of a Minecraft player's skin.
func RenderRightBody(skin *image.NRGBA, opts Options) *image.NRGBA {
	var (
		rightHead     *image.NRGBA = removeTransparency(extract(skin, 0, 8, 8, 8))
		rightRightArm *image.NRGBA = removeTransparency(extract(skin, 40, 20, 4, 12))
		rightRightLeg *image.NRGBA = removeTransparency(extract(skin, 0, 20, 4, 12))
	)

	if opts.Overlay && !IsOldSkin(skin) {
		overlaySkin := fixTransparency(skin)

		rightHead = composite(rightHead, extract(overlaySkin, 32, 8, 8, 8), 0, 0)
		rightRightArm = composite(rightRightArm, extract(overlaySkin, 40, 36, 4, 12), 0, 0)
		rightRightLeg = composite(rightRightLeg, extract(overlaySkin, 0, 36, 4, 12), 0, 0)
	}

	output := image.NewNRGBA(image.Rect(0, 0, 8, 32))

	// Right Head
	output = composite(output, rightHead, 0, 0)

	// Right Arm
	output = composite(output, rightRightArm, 2, 8)

	// Right Leg
	output = composite(output, rightRightLeg, 2, 20)

	return scale(output, opts.Scale)
}
