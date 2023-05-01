package skin

import (
	"image"
)

// RenderLeftBody renders a 2-dimensional image of the left side of a Minecraft player's skin.
func RenderLeftBody(skin *image.NRGBA, opts Options) *image.NRGBA {
	var (
		leftHead    *image.NRGBA = removeTransparency(extract(skin, HeadLeft))
		leftLeftArm *image.NRGBA = nil
		leftLeftLeg *image.NRGBA = nil
	)

	if IsOldSkin(skin) {
		leftLeftArm = flipHorizontal(removeTransparency(extract(skin, GetLeftArmLeft(false))))
		leftLeftLeg = flipHorizontal(removeTransparency(extract(skin, LeftLegLeft)))
	} else {
		leftLeftArm = removeTransparency(extract(skin, GetLeftArmLeft(opts.Slim)))
		leftLeftLeg = removeTransparency(extract(skin, LeftLegLeft))

		if opts.Overlay {
			overlaySkin := fixTransparency(skin)

			leftHead = composite(leftHead, extract(overlaySkin, HeadOverlayLeft), 0, 0)
			leftLeftArm = composite(leftLeftArm, extract(overlaySkin, GetLeftArmOverlayLeft(opts.Slim)), 0, 0)
			leftLeftLeg = composite(leftLeftLeg, extract(overlaySkin, LeftLegOverlayLeft), 0, 0)
		}
	}

	var output *image.NRGBA = image.NewNRGBA(image.Rect(0, 0, 8, 32))

	// Left Head
	output = composite(output, leftHead, 0, 0)

	// Left Arm
	output = composite(output, leftLeftArm, 2, 8)

	// Left Leg
	output = composite(output, leftLeftLeg, 2, 20)

	return scale(output, opts.Scale)
}
