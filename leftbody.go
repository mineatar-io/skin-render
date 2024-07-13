package skin

import (
	"image"
)

// RenderLeftBody renders a 2-dimensional image of the left side of a Minecraft player's skin.
func RenderLeftBody(img *image.NRGBA, opts Options) *image.NRGBA {
	if err := validateSkin(img); err != nil {
		panic(err)
	}

	var (
		skin        *image.NRGBA = convertToNRGBA(img)
		isOldSkin   bool         = IsOldSkin(skin)
		output      *image.NRGBA = image.NewNRGBA(image.Rect(0, 0, 16, 32))
		leftHead    *image.NRGBA = removeTransparency(extract(skin, HeadLeft))
		leftLeftArm *image.NRGBA = nil
		leftLeftLeg *image.NRGBA = nil
	)

	if isOldSkin {
		leftLeftArm = flipHorizontal(removeTransparency(extract(skin, GetRightArmLeft(false))))
		leftLeftLeg = flipHorizontal(removeTransparency(extract(skin, RightLegLeft)))
	} else {
		leftLeftArm = removeTransparency(extract(skin, GetLeftArmLeft(opts.Slim)))
		leftLeftLeg = removeTransparency(extract(skin, LeftLegLeft))
	}

	if opts.Overlay {
		overlaySkin := fixTransparency(skin)

		composite(leftHead, extract(overlaySkin, HeadOverlayLeft), 0, 0)

		if !isOldSkin {
			composite(leftLeftArm, extract(overlaySkin, GetLeftArmOverlayLeft(opts.Slim)), 0, 0)
			composite(leftLeftLeg, extract(overlaySkin, LeftLegOverlayLeft), 0, 0)
		}
	}

	// Left Head
	composite(output, leftHead, 4, 0)

	// Left Arm
	composite(output, leftLeftArm, 6, 8)

	// Left Leg
	composite(output, leftLeftLeg, 6, 20)

	if opts.Square {
		output = squareAndCenter(output)
	}

	return scale(output, opts.Scale)
}
