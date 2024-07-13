package skin

import (
	"image"
	"math"
)

// RenderBody renders a 3-dimensional image of the full body of a Minecraft player's skin.
func RenderBody(img *image.NRGBA, opts Options) *image.NRGBA {
	if err := validateSkin(img); err != nil {
		panic(err)
	}

	var (
		skin          *image.NRGBA = convertToNRGBA(img)
		scaleDouble   float64      = float64(opts.Scale)
		slimOffset    int          = getSlimOffset(opts.Slim)
		isOldSkin     bool         = IsOldSkin(skin)
		output        *image.NRGBA = image.NewNRGBA(image.Rect(0, 0, 17*opts.Scale+int(math.Ceil(scaleDouble*0.32)), 39*opts.Scale))
		frontHead     *image.NRGBA = removeTransparency(extract(skin, HeadFront))
		topHead       *image.NRGBA = rotate90(flipHorizontal(removeTransparency(extract(skin, HeadTop))))
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

		composite(topHead, rotate90(flipHorizontal(extract(overlaySkin, HeadOverlayTop))), 0, 0)
		composite(frontHead, extract(overlaySkin, HeadOverlayFront), 0, 0)
		composite(rightHead, extract(overlaySkin, HeadOverlayRight), 0, 0)

		if !isOldSkin {
			composite(frontTorso, extract(overlaySkin, TorsoOverlayFront), 0, 0)
			composite(frontLeftArm, extract(overlaySkin, GetLeftArmOverlayFront(opts.Slim)), 0, 0)
			composite(topLeftArm, extract(overlaySkin, GetLeftArmOverlayTop(opts.Slim)), 0, 0)
			composite(frontRightArm, extract(overlaySkin, GetRightArmOverlayFront(opts.Slim)), 0, 0)
			composite(topRightArm, extract(overlaySkin, GetRightArmOverlayTop(opts.Slim)), 0, 0)
			composite(rightRightArm, extract(overlaySkin, RightArmOverlayRight), 0, 0)
			composite(frontLeftLeg, extract(overlaySkin, LeftLegOverlayFront), 0, 0)
			composite(frontRightLeg, extract(overlaySkin, RightLegOverlayFront), 0, 0)
			composite(rightRightLeg, extract(overlaySkin, RightLegOverlayRight), 0, 0)
		}
	}

	// Right Side of Right Leg
	compositeTransform(output, scale(rightRightLeg, opts.Scale), sideMatrix, 4*scaleDouble, 23*scaleDouble)

	// Front of Right Leg
	compositeTransform(output, scale(frontRightLeg, opts.Scale), frontMatrix, 8*scaleDouble, 31*scaleDouble)

	// Front of Left Leg
	compositeTransform(output, scale(frontLeftLeg, opts.Scale), frontMatrix, 12*scaleDouble, 31*scaleDouble)

	// Front of Torso
	compositeTransform(output, scale(frontTorso, opts.Scale), frontMatrix, 8*scaleDouble, 19*scaleDouble)

	// Front of Right Arm
	compositeTransform(output, scale(frontRightArm, opts.Scale), frontMatrix, float64(4+slimOffset)*scaleDouble, 19*scaleDouble)

	// Front of Left Arm
	compositeTransform(output, scale(frontLeftArm, opts.Scale), frontMatrix, 16*scaleDouble, 19*scaleDouble)

	// Top of Left Arm
	compositeTransform(output, scale(rotate270(topLeftArm), opts.Scale), plantMatrix, 15*scaleDouble, float64(slimOffset-1)*scaleDouble)

	// Right Side of Right Arm
	compositeTransform(output, scale(rightRightArm, opts.Scale), sideMatrix, float64(slimOffset)*scaleDouble, float64(15-slimOffset)*scaleDouble)

	// Top of Right Arm
	compositeTransform(output, scale(rotate90(topRightArm), opts.Scale), plantMatrix, 15*scaleDouble, 11*scaleDouble)

	// Front of Head
	compositeTransform(output, scale(frontHead, opts.Scale), frontMatrix, 10*scaleDouble, 13*scaleDouble)

	// Top of Head
	compositeTransform(output, scale(topHead, opts.Scale), plantMatrix, 5*scaleDouble, -5*scaleDouble)

	// Right Side of Head
	compositeTransform(output, scale(rightHead, opts.Scale), sideMatrix, 2*scaleDouble, 3*scaleDouble)

	if opts.Square {
		return squareAndCenter(output)
	}

	return output
}
