package skin

import (
	"image"
)

// RenderRightBody renders a 2-dimensional image of the right side of a Minecraft player's skin.
func RenderRightBody(skin *image.NRGBA, opts Options) *image.NRGBA {
	var (
		rightHead     *image.NRGBA = removeTransparency(extract(skin, HeadRight))
		rightRightArm *image.NRGBA = removeTransparency(extract(skin, RightArmRight))
		rightRightLeg *image.NRGBA = removeTransparency(extract(skin, RightLegRight))
	)

	if opts.Overlay {
		overlaySkin := fixTransparency(skin)

		rightHead = composite(rightHead, extract(overlaySkin, HeadOverlayRight), 0, 0)

		if !IsOldSkin(skin) {
			rightRightArm = composite(rightRightArm, extract(overlaySkin, RightArmOverlayRight), 0, 0)
			rightRightLeg = composite(rightRightLeg, extract(overlaySkin, RightLegOverlayRight), 0, 0)
		}
	}

	var output *image.NRGBA = image.NewNRGBA(image.Rect(0, 0, 16, 32))

	// Right Head
	output = composite(output, rightHead, 4, 0)

	// Right Arm
	output = composite(output, rightRightArm, 6, 8)

	// Right Leg
	output = composite(output, rightRightLeg, 6, 20)

	return scale(output, opts.Scale)
}
