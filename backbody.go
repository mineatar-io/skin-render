package skin

import (
	"image"
)

// RenderBackBody renders a 2-dimensional image of the back of a Minecraft player's skin.
func RenderBackBody(img image.Image, opts Options) *image.NRGBA {
	if err := validateSkin(img); err != nil {
		panic(err)
	}

	var (
		skin         *image.NRGBA = convertToNRGBA(img)
		slimOffset   int          = getSlimOffset(opts.Slim)
		isOldSkin    bool         = IsOldSkin(skin)
		output       *image.NRGBA = image.NewNRGBA(image.Rect(0, 0, 16, 32))
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

		composite(backHead, extract(overlaySkin, HeadOverlayBack), 0, 0)

		if !isOldSkin {
			composite(backTorso, extract(overlaySkin, TorsoOverlayBack), 0, 0)
			composite(backLeftArm, extract(overlaySkin, GetLeftArmOverlayBack(opts.Slim)), 0, 0)
			composite(backRightArm, extract(overlaySkin, GetRightArmOverlayBack(opts.Slim)), 0, 0)
			composite(backLeftLeg, extract(overlaySkin, LeftLegOverlayBack), 0, 0)
			composite(backRightLeg, extract(overlaySkin, RightLegOverlayBack), 0, 0)
		}
	}

	// Face
	composite(output, backHead, 4, 0)

	// Torso
	composite(output, backTorso, 4, 8)

	// Left Arm
	composite(output, backLeftArm, slimOffset, 8)

	// Right Arm
	composite(output, backRightArm, 12, 8)

	// Left Leg
	composite(output, backLeftLeg, 4, 20)

	// Right Leg
	composite(output, backRightLeg, 8, 20)

	return scale(output, opts.Scale)
}
