package skin

import (
	"image"
)

// RenderBackBody renders a 2-dimensional image of the back of a Minecraft player's skin.
func RenderBackBody(skin *image.NRGBA, opts Options) *image.NRGBA {
	slimOffset := getSlimOffset(opts.Slim)

	var (
		backHead     *image.NRGBA = removeTransparency(extract(skin, 24, 8, 8, 8))
		backTorso    *image.NRGBA = removeTransparency(extract(skin, 32, 20, 8, 12))
		backLeftArm  *image.NRGBA = nil
		backRightArm *image.NRGBA = removeTransparency(extract(skin, 52-slimOffset, 20, 4-slimOffset, 12))
		backLeftLeg  *image.NRGBA = nil
		backRightLeg *image.NRGBA = removeTransparency(extract(skin, 12, 20, 4, 12))
	)

	if IsOldSkin(skin) {
		backLeftArm = flipHorizontal(backRightArm)
		backLeftLeg = flipHorizontal(backRightLeg)
	} else {
		backLeftArm = removeTransparency(extract(skin, 44-slimOffset, 52, 4-slimOffset, 12))
		backLeftLeg = removeTransparency(extract(skin, 28, 52, 4, 12))

		if opts.Overlay {
			overlaySkin := fixTransparency(skin)

			backHead = composite(backHead, extract(overlaySkin, 56, 8, 8, 8), 0, 0)
			backTorso = composite(backTorso, extract(overlaySkin, 32, 36, 8, 12), 0, 0)
			backLeftArm = composite(backLeftArm, extract(overlaySkin, 60-slimOffset, 52, 4-slimOffset, 64), 0, 0)
			backRightArm = composite(backRightArm, extract(overlaySkin, 52-slimOffset, 36, 4-slimOffset, 48), 0, 0)
			backLeftLeg = composite(backLeftLeg, extract(overlaySkin, 12, 52, 8, 64), 0, 0)
			backRightLeg = composite(backRightLeg, extract(overlaySkin, 12, 36, 8, 48), 0, 0)
		}
	}

	output := image.NewNRGBA(image.Rect(0, 0, 16, 32))

	// Face
	output = composite(output, backHead, 4, 0)

	// Torso
	output = composite(output, backTorso, 4, 8)

	// Left Arm
	output = composite(output, backLeftArm, slimOffset, 8)

	// Right Arm
	output = composite(output, backRightArm, 12, 8)

	// Left Leg
	output = composite(output, backLeftLeg, 4, 20)

	// Right Leg
	output = composite(output, backRightLeg, 8, 20)

	return scale(output, opts.Scale)
}
