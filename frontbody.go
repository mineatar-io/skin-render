package skin

import (
	"image"
)

// RenderFrontBody renders a 2-dimensional image of the front of a Minecraft player's skin.
func RenderFrontBody(skin *image.NRGBA, opts Options) *image.NRGBA {
	slimOffset := getSlimOffset(opts.Slim)

	var (
		frontHead  *image.NRGBA = removeTransparency(extract(skin, HeadFront))
		frontTorso *image.NRGBA = removeTransparency(extract(skin, TorsoFront))
		leftArm    *image.NRGBA = nil
		rightArm   *image.NRGBA = removeTransparency(extract(skin, GetRightArmFront(opts.Slim)))
		leftLeg    *image.NRGBA = nil
		rightLeg   *image.NRGBA = removeTransparency(extract(skin, RightLegFront))
	)

	if IsOldSkin(skin) {
		leftArm = flipHorizontal(rightArm)
		leftLeg = flipHorizontal(rightLeg)
	} else {
		leftArm = removeTransparency(extract(skin, GetLeftArmFront(opts.Slim)))
		leftLeg = removeTransparency(extract(skin, LeftLegFront))

		if opts.Overlay {
			overlaySkin := fixTransparency(skin)

			frontHead = composite(frontHead, extract(overlaySkin, HeadOverlayFront), 0, 0)
			frontTorso = composite(frontTorso, extract(overlaySkin, TorsoOverlayFront), 0, 0)
			leftArm = composite(leftArm, extract(overlaySkin, GetLeftArmOverlayFront(opts.Slim)), 0, 0)
			rightArm = composite(rightArm, extract(overlaySkin, GetRightArmOverlayFront(opts.Slim)), 0, 0)
			leftLeg = composite(leftLeg, extract(overlaySkin, LeftLegOverlayFront), 0, 0)
			rightLeg = composite(rightLeg, extract(overlaySkin, RightLegOverlayFront), 0, 0)
		}
	}

	var output *image.NRGBA = image.NewNRGBA(image.Rect(0, 0, 16, 32))

	// Face
	output = composite(output, frontHead, 4, 0)

	// Torso
	output = composite(output, frontTorso, 4, 8)

	// Left Arm
	output = composite(output, leftArm, 12, 8)

	// Right Arm
	output = composite(output, rightArm, slimOffset, 8)

	// Left Leg
	output = composite(output, leftLeg, 8, 20)

	// Right Leg
	output = composite(output, rightLeg, 4, 20)

	return scale(output, opts.Scale)
}
