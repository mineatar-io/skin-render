package skin

import (
	"image"
)

// RenderRightBody renders a 2-dimensional image of the right side of a Minecraft player's skin.
func RenderRightBody(img *image.NRGBA, opts Options) *image.NRGBA {
	if err := validateSkin(img); err != nil {
		panic(err)
	}

	var (
		skin          *image.NRGBA = convertToNRGBA(img)
		output        *image.NRGBA = image.NewNRGBA(image.Rect(0, 0, 16, 32))
		rightHead     *image.NRGBA = removeTransparency(extract(skin, HeadRight))
		rightRightArm *image.NRGBA = removeTransparency(extract(skin, RightArmRight))
		rightRightLeg *image.NRGBA = removeTransparency(extract(skin, RightLegRight))
	)

	if opts.Overlay {
		overlaySkin := fixTransparency(skin)

		composite(rightHead, extract(overlaySkin, HeadOverlayRight), 0, 0)

		if !IsOldSkin(skin) {
			composite(rightRightArm, extract(overlaySkin, RightArmOverlayRight), 0, 0)
			composite(rightRightLeg, extract(overlaySkin, RightLegOverlayRight), 0, 0)
		}
	}

	// Right Head
	composite(output, rightHead, 4, 0)

	// Right Arm
	composite(output, rightRightArm, 6, 8)

	// Right Leg
	composite(output, rightRightLeg, 6, 20)

	return scale(output, opts.Scale)
}
