package skin

import (
	"image"
	"math"
)

// RenderBody renders a 3-dimensional image of the full body of a Minecraft player's skin.
func RenderBody(skin *image.NRGBA, opts Options) *image.NRGBA {
	scaleDouble := float64(opts.Scale)
	slimOffset := getSlimOffset(opts.Slim)
	isOldSkin := IsOldSkin(skin)

	var output *image.NRGBA = image.NewNRGBA(image.Rect(0, 0, 17*opts.Scale+int(math.Ceil(scaleDouble*0.32)), 39*opts.Scale))

	var (
		frontHead     *image.NRGBA = removeTransparency(extract(skin, HeadFront))
		topHead       *image.NRGBA = removeTransparency(extract(skin, HeadTop))
		rightHead     *image.NRGBA = removeTransparency(extract(skin, HeadRight))
		frontTorso    *image.NRGBA = removeTransparency(extract(skin, TorsoFront))
		frontLeftArm  *image.NRGBA = nil
		topLeftArm    *image.NRGBA = nil
		frontRightArm *image.NRGBA = removeTransparency(extract(skin, GetRightArmFront(opts.Slim)))
		topRightArm   *image.NRGBA = removeTransparency(extract(skin, GetRightArmTop(opts.Slim)))
		rightRightArm *image.NRGBA = removeTransparency(extract(skin, RightArmRight))
		frontLeftLeg  *image.NRGBA = nil
		frontRightLeg *image.NRGBA = removeTransparency(extract(skin, RightLegFront))
		rightRightLeg *image.NRGBA = removeTransparency(extract(skin, RightLegRight))
	)

	if isOldSkin {
		frontLeftArm = flipHorizontal(frontRightArm)
		topLeftArm = flipHorizontal(topRightArm)
		frontLeftLeg = flipHorizontal(frontRightLeg)
	} else {
		frontLeftArm = removeTransparency(extract(skin, GetLeftArmFront(opts.Slim)))
		topLeftArm = removeTransparency(extract(skin, GetLeftArmTop(opts.Slim)))
		frontLeftLeg = removeTransparency(extract(skin, LeftLegFront))
	}

	if opts.Overlay {
		overlaySkin := fixTransparency(skin)

		topHead = composite(topHead, extract(overlaySkin, HeadOverlayTop), 0, 0)
		frontHead = composite(frontHead, extract(overlaySkin, HeadOverlayFront), 0, 0)
		rightHead = composite(rightHead, extract(overlaySkin, HeadOverlayRight), 0, 0)

		if !isOldSkin {
			frontTorso = composite(frontTorso, extract(overlaySkin, TorsoOverlayFront), 0, 0)
			frontLeftArm = composite(frontLeftArm, extract(overlaySkin, GetLeftArmOverlayFront(opts.Slim)), 0, 0)
			topLeftArm = composite(topLeftArm, extract(overlaySkin, GetLeftArmOverlayTop(opts.Slim)), 0, 0)
			frontRightArm = composite(frontRightArm, extract(overlaySkin, GetRightArmOverlayFront(opts.Slim)), 0, 0)
			topRightArm = composite(topRightArm, extract(overlaySkin, GetRightArmOverlayTop(opts.Slim)), 0, 0)
			rightRightArm = composite(rightRightArm, extract(overlaySkin, RightArmOverlayRight), 0, 0)
			frontLeftLeg = composite(frontLeftLeg, extract(overlaySkin, LeftLegOverlayFront), 0, 0)
			frontRightLeg = composite(frontRightLeg, extract(overlaySkin, RightLegOverlayFront), 0, 0)
			rightRightLeg = composite(rightRightLeg, extract(overlaySkin, RightLegOverlayRight), 0, 0)
		}
	}

	// Right Side of Right Leg
	output = compositeTransform(output, scale(rightRightLeg, opts.Scale), sideMatrix, 4*scaleDouble, 23*scaleDouble)

	// Front of Right Leg
	output = compositeTransform(output, scale(frontRightLeg, opts.Scale), frontMatrix, 8*scaleDouble, 31*scaleDouble)

	// Front of Left Leg
	output = compositeTransform(output, scale(frontLeftLeg, opts.Scale), frontMatrix, 12*scaleDouble, 31*scaleDouble)

	// Front of Torso
	output = compositeTransform(output, scale(frontTorso, opts.Scale), frontMatrix, 8*scaleDouble, 19*scaleDouble)

	// Front of Right Arm
	output = compositeTransform(output, scale(frontRightArm, opts.Scale), frontMatrix, float64(4+slimOffset)*scaleDouble, 19*scaleDouble)

	// Front of Left Arm
	output = compositeTransform(output, scale(frontLeftArm, opts.Scale), frontMatrix, 16*scaleDouble, 19*scaleDouble)

	// Top of Left Arm
	output = compositeTransform(output, scale(rotate270(topLeftArm), opts.Scale), plantMatrix, 15*scaleDouble, float64(slimOffset-1)*scaleDouble)

	// Right Side of Right Arm
	output = compositeTransform(output, scale(rightRightArm, opts.Scale), sideMatrix, float64(slimOffset)*scaleDouble, float64(15-slimOffset)*scaleDouble)

	// Top of Right Arm
	output = compositeTransform(output, scale(rotate90(topRightArm), opts.Scale), plantMatrix, 15*scaleDouble, 11*scaleDouble)

	// Front of Head
	output = compositeTransform(output, scale(frontHead, opts.Scale), frontMatrix, 10*scaleDouble, 13*scaleDouble)

	// Top of Head
	output = compositeTransform(output, scale(topHead, opts.Scale), plantMatrix, 5*scaleDouble, -5*scaleDouble)

	// Right Side of Head
	output = compositeTransform(output, scale(rightHead, opts.Scale), sideMatrix, 2*scaleDouble, 3*scaleDouble)

	return output
}
