package skin

import (
	"image"
)

// RenderFrontBody renders a 2-dimensional image of the front of a Minecraft player's skin.
func RenderFrontBody(img *image.NRGBA, opts Options) *image.NRGBA {
	if err := validateSkin(img); err != nil {
		panic(err)
	}

	var (
		skin       *image.NRGBA = convertToNRGBA(img)
		slimOffset int          = getSlimOffset(opts.Slim)
		isOldSkin  bool         = IsOldSkin(skin)
		output     *image.NRGBA = image.NewNRGBA(image.Rect(0, 0, 16, 32))
		frontHead  *image.NRGBA = removeTransparency(extract(skin, HeadFront))
		frontTorso *image.NRGBA = removeTransparency(extract(skin, TorsoFront))
		leftArm    *image.NRGBA = nil
		rightArm   *image.NRGBA = removeTransparency(extract(skin, GetRightArmFront(opts.Slim)))
		leftLeg    *image.NRGBA = nil
		rightLeg   *image.NRGBA = removeTransparency(extract(skin, RightLegFront))
	)

	if isOldSkin {
		leftArm = flipHorizontal(rightArm)
		leftLeg = flipHorizontal(rightLeg)
	} else {
		leftArm = removeTransparency(extract(skin, GetLeftArmFront(opts.Slim)))
		leftLeg = removeTransparency(extract(skin, LeftLegFront))
	}

	if opts.Overlay {
		overlaySkin := fixTransparency(skin)

		composite(frontHead, extract(overlaySkin, HeadOverlayFront), 0, 0)

		if !isOldSkin {
			composite(frontTorso, extract(overlaySkin, TorsoOverlayFront), 0, 0)
			composite(leftArm, extract(overlaySkin, GetLeftArmOverlayFront(opts.Slim)), 0, 0)
			composite(rightArm, extract(overlaySkin, GetRightArmOverlayFront(opts.Slim)), 0, 0)
			composite(leftLeg, extract(overlaySkin, LeftLegOverlayFront), 0, 0)
			composite(rightLeg, extract(overlaySkin, RightLegOverlayFront), 0, 0)
		}
	}

	// Face
	composite(output, frontHead, 4, 0)

	// Torso
	composite(output, frontTorso, 4, 8)

	// Left Arm
	composite(output, leftArm, 12, 8)

	// Right Arm
	composite(output, rightArm, slimOffset, 8)

	// Left Leg
	composite(output, leftLeg, 8, 20)

	// Right Leg
	composite(output, rightLeg, 4, 20)

	if opts.Square {
		output = squareAndCenter(output)
	}

	return scale(output, opts.Scale)
}
