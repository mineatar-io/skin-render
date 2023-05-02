package skin

import (
	"image"
)

// RenderBackBody renders a 2-dimensional image of the back of a Minecraft player's skin.
func RenderBackBody(skin *image.NRGBA, opts Options) *image.NRGBA {
	slimOffset := getSlimOffset(opts.Slim)
	isOldSkin := IsOldSkin(skin)

	var (
		backHead     *image.NRGBA = removeTransparency(extract(skin, HeadBack))
		backTorso    *image.NRGBA = removeTransparency(extract(skin, TorsoBack))
		backLeftArm  *image.NRGBA = nil
		backRightArm *image.NRGBA = removeTransparency(extract(skin, GetRightArmBack(opts.Slim)))
		backLeftLeg  *image.NRGBA = nil
		backRightLeg *image.NRGBA = removeTransparency(extract(skin, RightLegBack))
	)

	if isOldSkin {
		backLeftArm = flipHorizontal(backRightArm)
		backLeftLeg = flipHorizontal(backRightLeg)
	} else {
		backLeftArm = removeTransparency(extract(skin, GetLeftArmBack(opts.Slim)))
		backLeftLeg = removeTransparency(extract(skin, LeftLegBack))
	}

	if opts.Overlay {
		overlaySkin := fixTransparency(skin)

		backHead = composite(backHead, extract(overlaySkin, HeadOverlayBack), 0, 0)

		if !isOldSkin {
			backTorso = composite(backTorso, extract(overlaySkin, TorsoOverlayBack), 0, 0)
			backLeftArm = composite(backLeftArm, extract(overlaySkin, GetLeftArmOverlayBack(opts.Slim)), 0, 0)
			backRightArm = composite(backRightArm, extract(overlaySkin, GetRightArmOverlayBack(opts.Slim)), 0, 0)
			backLeftLeg = composite(backLeftLeg, extract(overlaySkin, LeftLegOverlayBack), 0, 0)
			backRightLeg = composite(backRightLeg, extract(overlaySkin, RightLegOverlayBack), 0, 0)
		}
	}

	var output *image.NRGBA = image.NewNRGBA(image.Rect(0, 0, 16, 32))

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
