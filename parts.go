package skin

import "image"

var (
	// HeadTop is the top side of the head
	HeadTop image.Rectangle = image.Rect(8, 0, 16, 8)
	// HeadBottom is the bottom side of the head
	HeadBottom image.Rectangle = image.Rect(16, 0, 24, 8)
	// HeadRight is the right side of the head
	HeadRight image.Rectangle = image.Rect(0, 8, 8, 16)
	// HeadFront is the front side of the head
	HeadFront image.Rectangle = image.Rect(8, 8, 16, 16)
	// HeadLeft is the left side of the head
	HeadLeft image.Rectangle = image.Rect(16, 8, 24, 16)
	// HeadBack is the back side of the head
	HeadBack image.Rectangle = image.Rect(24, 8, 32, 16)
	// HeadOverlayTop is the top side of the head overlay
	HeadOverlayTop image.Rectangle = image.Rect(40, 0, 48, 8)
	// HeadOverlayBottom is the bottom side of the head overlay
	HeadOverlayBottom image.Rectangle = image.Rect(48, 0, 56, 8)
	// HeadOverlayRight is the right side of the head overlay
	HeadOverlayRight image.Rectangle = image.Rect(32, 8, 40, 16)
	// HeadOverlayFront is the front side of the head overlay
	HeadOverlayFront image.Rectangle = image.Rect(40, 8, 48, 16)
	// HeadOverlayLeft is the left side of the head overlay
	HeadOverlayLeft image.Rectangle = image.Rect(48, 8, 56, 16)
	// HeadOverlayBack is the back side of the head overlay
	HeadOverlayBack image.Rectangle = image.Rect(56, 8, 64, 16)
	// RightLegTop is the top side of the right leg
	RightLegTop image.Rectangle = image.Rect(4, 16, 8, 20)
	// RightLegBottom is the bottom side of the right leg
	RightLegBottom image.Rectangle = image.Rect(8, 16, 12, 20)
	// RightLegRight is the right side of the right leg
	RightLegRight image.Rectangle = image.Rect(0, 20, 4, 32)
	// RightLegFront is the front side of the right leg
	RightLegFront image.Rectangle = image.Rect(4, 20, 8, 32)
	// RightLegLeft is the left side of the right leg
	RightLegLeft image.Rectangle = image.Rect(8, 20, 12, 32)
	// RightLegBack is the back side of the right leg
	RightLegBack image.Rectangle = image.Rect(12, 20, 16, 32)
	// TorsoTop is the top side of the torso
	TorsoTop image.Rectangle = image.Rect(20, 16, 28, 20)
	// TorsoBottom is the bottom side of the torso
	TorsoBottom image.Rectangle = image.Rect(28, 16, 36, 20)
	// TorsoRight is the right side of the torso
	TorsoRight image.Rectangle = image.Rect(16, 20, 20, 32)
	// TorsoFront is the front side of the torso
	TorsoFront image.Rectangle = image.Rect(20, 20, 28, 32)
	// TorsoLeft is the left side of the torso
	TorsoLeft image.Rectangle = image.Rect(28, 20, 32, 32)
	// TorsoBack is the back side of the torso
	TorsoBack image.Rectangle = image.Rect(32, 20, 40, 32)
	// RightArmTopRegular is the top side of the right arm for regular skin models
	RightArmTopRegular image.Rectangle = image.Rect(44, 16, 48, 20)
	// RightArmTopSlim is the top side of the right arm for slim skin models
	RightArmTopSlim image.Rectangle = image.Rect(44, 16, 47, 20)
	// RightArmBottomRegular is the bottom side of the right arm for regular skin models
	RightArmBottomRegular image.Rectangle = image.Rect(48, 16, 52, 20)
	// RightArmBottomSlim is the bottom side of the right arm for slim skin models
	RightArmBottomSlim image.Rectangle = image.Rect(47, 16, 50, 20)
	// RightArmRight is the right side of the right arm
	RightArmRight image.Rectangle = image.Rect(40, 20, 44, 32)
	// RightArmFrontRegular is the front side of the right arm for regular skin models
	RightArmFrontRegular image.Rectangle = image.Rect(44, 20, 48, 32)
	// RightArmFrontSlim is the front side of the right arm for slim skin models
	RightArmFrontSlim image.Rectangle = image.Rect(44, 20, 47, 32)
	// RightArmLeftRegular is the left side of the right arm for regular skin models
	RightArmLeftRegular image.Rectangle = image.Rect(48, 20, 52, 32)
	// RightArmLeftSlim is the left side of the right arm for slim skin models
	RightArmLeftSlim image.Rectangle = image.Rect(47, 20, 51, 32)
	// RightArmBackRegular is the back side of the right arm for regular skin models
	RightArmBackRegular image.Rectangle = image.Rect(52, 20, 56, 32)
	// RightArmBackSlim is the back side of the right arm for slim skin models
	RightArmBackSlim image.Rectangle = image.Rect(51, 20, 54, 32)
	// LeftLegTop is the top side of the left leg
	LeftLegTop image.Rectangle = image.Rect(20, 48, 24, 52)
	// LeftLegBottom is the bottom side of the left leg
	LeftLegBottom image.Rectangle = image.Rect(24, 48, 28, 52)
	// LeftLegRight is the right side of the left leg
	LeftLegRight image.Rectangle = image.Rect(16, 52, 20, 64)
	// LeftLegFront is the front side of the left leg
	LeftLegFront image.Rectangle = image.Rect(20, 52, 24, 64)
	// LeftLegLeft is the left side of the left leg
	LeftLegLeft image.Rectangle = image.Rect(24, 52, 28, 64)
	// LeftLegBack is the back side of the left leg
	LeftLegBack image.Rectangle = image.Rect(28, 52, 32, 64)
	// LeftArmTopRegular is the top side of the left arm for regular skin models
	LeftArmTopRegular image.Rectangle = image.Rect(36, 48, 40, 52)
	// LeftArmTopSlim is the top side of the left arm for slim skin models
	LeftArmTopSlim image.Rectangle = image.Rect(36, 48, 39, 52)
	// LeftArmBottomRegular is the bottom side of the left arm for regular skin models
	LeftArmBottomRegular image.Rectangle = image.Rect(40, 48, 44, 52)
	// LeftArmBottomSlim is the bottom side of the left arm for slim skin models
	LeftArmBottomSlim image.Rectangle = image.Rect(39, 48, 42, 52)
	// LeftArmRight is the right side of the left arm
	LeftArmRight image.Rectangle = image.Rect(32, 52, 36, 64)
	// LeftArmFrontRegular is the front side of the left arm for regular skin models
	LeftArmFrontRegular image.Rectangle = image.Rect(36, 52, 40, 64)
	// LeftArmFrontSlim is the front side of the left arm for slim skin models
	LeftArmFrontSlim image.Rectangle = image.Rect(36, 52, 39, 64)
	// LeftArmLeftRegular is the left side of the left arm for regular skin models
	LeftArmLeftRegular image.Rectangle = image.Rect(40, 52, 44, 64)
	// LeftArmLeftSlim is the left side of the left arm for slim skin models
	LeftArmLeftSlim image.Rectangle = image.Rect(39, 52, 43, 64)
	// LeftArmBackRegular is the back side of the left arm for regular skin models
	LeftArmBackRegular image.Rectangle = image.Rect(44, 52, 48, 64)
	// LeftArmBackSlim is the back side of the left arm for slim skin models
	LeftArmBackSlim image.Rectangle = image.Rect(43, 52, 46, 64)
	// RightLegOverlayTop is the top side of the right leg overlay
	RightLegOverlayTop image.Rectangle = image.Rect(4, 48, 8, 36)
	// RightLegOverlayBottom is the bottom side of the right leg overlay
	RightLegOverlayBottom image.Rectangle = image.Rect(8, 48, 12, 36)
	// RightLegOverlayRight is the right side of the right leg overlay
	RightLegOverlayRight image.Rectangle = image.Rect(0, 36, 4, 48)
	// RightLegOverlayFront is the front side of the right leg overlay
	RightLegOverlayFront image.Rectangle = image.Rect(4, 36, 8, 48)
	// RightLegOverlayLeft is the left side of the right leg overlay
	RightLegOverlayLeft image.Rectangle = image.Rect(8, 36, 12, 48)
	// RightLegOverlayBack is the back side of the right leg overlay
	RightLegOverlayBack image.Rectangle = image.Rect(12, 36, 16, 48)
	// TorsoOverlayTop is the top side of the torso overlay
	TorsoOverlayTop image.Rectangle = image.Rect(20, 48, 28, 36)
	// TorsoOverlayBottom is the bottom side of the torso overlay
	TorsoOverlayBottom image.Rectangle = image.Rect(28, 48, 36, 36)
	// TorsoOverlayRight is the right side of the torso overlay
	TorsoOverlayRight image.Rectangle = image.Rect(16, 36, 20, 48)
	// TorsoOverlayFront is the front side of the torso overlay
	TorsoOverlayFront image.Rectangle = image.Rect(20, 36, 28, 48)
	// TorsoOverlayLeft is the left side of the torso overlay
	TorsoOverlayLeft image.Rectangle = image.Rect(28, 36, 32, 48)
	// TorsoOverlayBack is the back side of the torso overlay
	TorsoOverlayBack image.Rectangle = image.Rect(32, 36, 40, 48)
	// RightArmOverlayTopRegular is the top side of the right arm overlay for regular skin models
	RightArmOverlayTopRegular image.Rectangle = image.Rect(44, 48, 48, 36)
	// RightArmOverlayTopSlim is the top side of the right arm overlay for slim skin models
	RightArmOverlayTopSlim image.Rectangle = image.Rect(44, 48, 47, 36)
	// RightArmOverlayBottomRegular is the bottom side of the right arm overlay for regular skin models
	RightArmOverlayBottomRegular image.Rectangle = image.Rect(48, 48, 52, 36)
	// RightArmOverlayBottomSlim is the bottom side of the right arm overlay for slim skin models
	RightArmOverlayBottomSlim image.Rectangle = image.Rect(47, 48, 50, 36)
	// RightArmOverlayRight is the right side of the right arm overlay
	RightArmOverlayRight image.Rectangle = image.Rect(40, 36, 44, 48)
	// RightArmOverlayFrontRegular is the front side of the right arm overlay for regular skin models
	RightArmOverlayFrontRegular image.Rectangle = image.Rect(44, 36, 48, 48)
	// RightArmOverlayFrontSlim is the front side of the right arm overlay for slim skin models
	RightArmOverlayFrontSlim image.Rectangle = image.Rect(44, 36, 47, 48)
	// RightArmOverlayLeftRegular is the left side of the right arm overlay for regular skin models
	RightArmOverlayLeftRegular image.Rectangle = image.Rect(48, 36, 52, 48)
	// RightArmOverlayLeftSlim is the left side of the right arm overlay for slim skin models
	RightArmOverlayLeftSlim image.Rectangle = image.Rect(47, 36, 51, 48)
	// RightArmOverlayBackRegular is the back side of the right arm overlay for regular skin models
	RightArmOverlayBackRegular image.Rectangle = image.Rect(52, 36, 56, 48)
	// RightArmOverlayBackSlim is the back side of the right arm overlay for slim skin models
	RightArmOverlayBackSlim image.Rectangle = image.Rect(51, 36, 54, 48)
	// LeftLegOverlayTop is the top side of the left leg overlay
	LeftLegOverlayTop image.Rectangle = image.Rect(4, 48, 8, 52)
	// LeftLegOverlayBottom is the bottom side of the left leg overlay
	LeftLegOverlayBottom image.Rectangle = image.Rect(8, 48, 12, 52)
	// LeftLegOverlayRight is the right side of the left leg overlay
	LeftLegOverlayRight image.Rectangle = image.Rect(0, 52, 4, 64)
	// LeftLegOverlayFront is the front side of the left leg overlay
	LeftLegOverlayFront image.Rectangle = image.Rect(4, 52, 8, 64)
	// LeftLegOverlayLeft is the left side of the left leg overlay
	LeftLegOverlayLeft image.Rectangle = image.Rect(8, 52, 12, 64)
	// LeftLegOverlayBack is the back side of the left leg overlay
	LeftLegOverlayBack image.Rectangle = image.Rect(12, 52, 16, 64)
	// LeftArmOverlayTopRegular is the top side of the left arm overlay for regular skin models
	LeftArmOverlayTopRegular image.Rectangle = image.Rect(52, 48, 56, 52)
	// LeftArmOverlayTopSlim is the top side of the left arm overlay for slim skin models
	LeftArmOverlayTopSlim image.Rectangle = image.Rect(52, 48, 55, 52)
	// LeftArmOverlayBottomRegular is the bottom side of the left arm overlay for regular skin models
	LeftArmOverlayBottomRegular image.Rectangle = image.Rect(56, 48, 60, 52)
	// LeftArmOverlayBottomSlim is the bottom side of the left arm overlay for slim skin models
	LeftArmOverlayBottomSlim image.Rectangle = image.Rect(55, 48, 58, 52)
	// LeftArmOverlayRight is the right side of the left arm overlay
	LeftArmOverlayRight image.Rectangle = image.Rect(48, 52, 52, 64)
	// LeftArmOverlayFrontRegular is the front side of the left arm overlay for regular skin models
	LeftArmOverlayFrontRegular image.Rectangle = image.Rect(52, 52, 56, 64)
	// LeftArmOverlayFrontSlim is the front side of the left arm overlay for slim skin models
	LeftArmOverlayFrontSlim image.Rectangle = image.Rect(52, 52, 55, 64)
	// LeftArmOverlayLeftRegular is the left side of the left arm overlay for regular skin models
	LeftArmOverlayLeftRegular image.Rectangle = image.Rect(56, 52, 60, 64)
	// LeftArmOverlayLeftSlim is the left side of the left arm overlay for slim skin models
	LeftArmOverlayLeftSlim image.Rectangle = image.Rect(55, 52, 59, 64)
	// LeftArmOverlayBackRegular is the back side of the left arm overlay for regular skin models
	LeftArmOverlayBackRegular image.Rectangle = image.Rect(60, 52, 64, 64)
	// LeftArmOverlayBackSlim is the back side of the left arm overlay for slim skin models
	LeftArmOverlayBackSlim image.Rectangle = image.Rect(59, 52, 62, 64)
)

// GetRightArmTop returns the top of a right arm based on if the skin is slim or not
func GetRightArmTop(slim bool) image.Rectangle {
	if slim {
		return RightArmTopSlim
	}

	return RightArmTopRegular
}

// GetRightArmBottom returns the bottom of a right arm based on if the skin is slim or not
func GetRightArmBottom(slim bool) image.Rectangle {
	if slim {
		return RightArmBottomSlim
	}

	return RightArmBottomRegular
}

// GetRightArmFront returns the front of a right arm based on if the skin is slim or not
func GetRightArmFront(slim bool) image.Rectangle {
	if slim {
		return RightArmFrontSlim
	}

	return RightArmFrontRegular
}

// GetRightArmLeft returns the left of a right arm based on if the skin is slim or not
func GetRightArmLeft(slim bool) image.Rectangle {
	if slim {
		return RightArmLeftSlim
	}

	return RightArmLeftRegular
}

// GetRightArmBack returns the back of a right arm based on if the skin is slim or not
func GetRightArmBack(slim bool) image.Rectangle {
	if slim {
		return RightArmBackSlim
	}

	return RightArmBackRegular
}

// GetLeftArmTop returns the top of a left arm based on if the skin is slim or not
func GetLeftArmTop(slim bool) image.Rectangle {
	if slim {
		return LeftArmTopSlim
	}

	return LeftArmTopRegular
}

// GetLeftArmBottom returns the bottom of a left arm based on if the skin is slim or not
func GetLeftArmBottom(slim bool) image.Rectangle {
	if slim {
		return LeftArmBottomSlim
	}

	return LeftArmBottomRegular
}

// GetLeftArmFront returns the front of a left arm based on if the skin is slim or not
func GetLeftArmFront(slim bool) image.Rectangle {
	if slim {
		return LeftArmFrontSlim
	}

	return LeftArmFrontRegular
}

// GetLeftArmLeft returns the left of a left arm based on if the skin is slim or not
func GetLeftArmLeft(slim bool) image.Rectangle {
	if slim {
		return LeftArmLeftSlim
	}

	return LeftArmLeftRegular
}

// GetLeftArmBack returns the back of a left arm based on if the skin is slim or not
func GetLeftArmBack(slim bool) image.Rectangle {
	if slim {
		return LeftArmBackSlim
	}

	return LeftArmBackRegular
}

// GetRightArmOverlayTop returns the top of a right arm overlay based on if the skin is slim or not
func GetRightArmOverlayTop(slim bool) image.Rectangle {
	if slim {
		return RightArmOverlayTopSlim
	}

	return RightArmOverlayTopRegular
}

// GetRightArmOverlayBottom returns the bottom of a right arm overlay based on if the skin is slim or not
func GetRightArmOverlayBottom(slim bool) image.Rectangle {
	if slim {
		return RightArmOverlayBottomSlim
	}

	return RightArmOverlayBottomRegular
}

// GetRightArmOverlayFront returns the front of a right arm overlay based on if the skin is slim or not
func GetRightArmOverlayFront(slim bool) image.Rectangle {
	if slim {
		return RightArmOverlayFrontSlim
	}

	return RightArmOverlayFrontRegular
}

// GetRightArmOverlayLeft returns the left of a right arm overlay based on if the skin is slim or not
func GetRightArmOverlayLeft(slim bool) image.Rectangle {
	if slim {
		return RightArmOverlayLeftSlim
	}

	return RightArmOverlayLeftRegular
}

// GetRightArmOverlayBack returns the back of a right arm overlay based on if the skin is slim or not
func GetRightArmOverlayBack(slim bool) image.Rectangle {
	if slim {
		return RightArmOverlayBackSlim
	}

	return RightArmOverlayBackRegular
}

// GetLeftArmOverlayTop returns the top of a left arm overlay based on if the skin is slim or not
func GetLeftArmOverlayTop(slim bool) image.Rectangle {
	if slim {
		return LeftArmOverlayTopSlim
	}

	return LeftArmOverlayTopRegular
}

// GetLeftArmOverlayBottom returns the bottom of a left arm overlay based on if the skin is slim or not
func GetLeftArmOverlayBottom(slim bool) image.Rectangle {
	if slim {
		return LeftArmOverlayBottomSlim
	}

	return LeftArmOverlayBottomRegular
}

// GetLeftArmOverlayFront returns the front of a left arm overlay based on if the skin is slim or not
func GetLeftArmOverlayFront(slim bool) image.Rectangle {
	if slim {
		return LeftArmOverlayFrontSlim
	}

	return LeftArmOverlayFrontRegular
}

// GetLeftArmOverlayLeft returns the left of a left arm overlay based on if the skin is slim or not
func GetLeftArmOverlayLeft(slim bool) image.Rectangle {
	if slim {
		return LeftArmOverlayLeftSlim
	}

	return LeftArmOverlayLeftRegular
}

// GetLeftArmOverlayBack returns the back of a left arm overlay based on if the skin is slim or not
func GetLeftArmOverlayBack(slim bool) image.Rectangle {
	if slim {
		return LeftArmOverlayBackSlim
	}

	return LeftArmOverlayBackRegular
}
