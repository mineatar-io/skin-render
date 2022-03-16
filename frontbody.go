package skin

import (
	"image"
)

func RenderFrontBody(skin *image.NRGBA, opts Options) *image.NRGBA {
	slimOffset := getSlimOffset(opts.Slim)

	var (
		frontHead  *image.NRGBA = removeTransparency(extract(skin, 8, 8, 8, 8))
		frontTorso *image.NRGBA = removeTransparency(extract(skin, 20, 20, 8, 12))
		leftArm    *image.NRGBA = nil
		rightArm   *image.NRGBA = removeTransparency(extract(skin, 44, 20, 4-slimOffset, 12))
		leftLeg    *image.NRGBA = nil
		rightLeg   *image.NRGBA = removeTransparency(extract(skin, 4, 20, 4, 12))
	)

	if IsOldSkin(skin) {
		leftArm = flipHorizontal(rightArm)
		leftLeg = flipHorizontal(rightLeg)
	} else {
		leftArm = removeTransparency(extract(skin, 36, 52, 4-slimOffset, 12))
		leftLeg = removeTransparency(extract(skin, 20, 52, 4, 12))

		if opts.Overlay {
			overlaySkin := fixTransparency(skin)

			frontHead = composite(frontHead, extract(overlaySkin, 40, 8, 8, 8), 0, 0)
			frontTorso = composite(frontTorso, extract(overlaySkin, 20, 36, 8, 12), 0, 0)
			leftArm = composite(leftArm, extract(overlaySkin, 52, 52, 4-slimOffset, 64), 0, 0)
			rightArm = composite(rightArm, extract(overlaySkin, 44, 36, 4-slimOffset, 48), 0, 0)
			leftLeg = composite(leftLeg, extract(overlaySkin, 4, 52, 4, 12), 0, 0)
			rightLeg = composite(rightLeg, extract(overlaySkin, 4, 36, 4, 12), 0, 0)
		}
	}

	output := image.NewNRGBA(image.Rect(0, 0, 16, 32))

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
