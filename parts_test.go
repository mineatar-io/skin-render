package skin_test

import (
	"image"
	"testing"

	"github.com/mineatar-io/skin-render"
)

var (
	headComponents []image.Rectangle = []image.Rectangle{
		skin.HeadTop,
		skin.HeadBottom,
		skin.HeadRight,
		skin.HeadFront,
		skin.HeadLeft,
		skin.HeadBack,
		skin.HeadOverlayTop,
		skin.HeadOverlayBottom,
		skin.HeadOverlayRight,
		skin.HeadOverlayFront,
		skin.HeadOverlayLeft,
		skin.HeadOverlayBack,
	}
	armRegularSideComponents []image.Rectangle = []image.Rectangle{
		skin.RightArmRight,
		skin.RightArmFrontRegular,
		skin.RightArmLeftRegular,
		skin.RightArmLeftSlim,
		skin.RightArmBackRegular,
		skin.LeftArmFrontRegular,
		skin.LeftArmLeftRegular,
		skin.LeftArmLeftSlim,
		skin.LeftArmBackRegular,
		skin.RightArmOverlayRight,
		skin.RightArmOverlayFrontRegular,
		skin.RightArmOverlayLeftRegular,
		skin.RightArmOverlayLeftSlim,
		skin.RightArmOverlayBackRegular,
		skin.LeftArmOverlayRight,
		skin.LeftArmOverlayFrontRegular,
		skin.LeftArmOverlayLeftRegular,
		skin.LeftArmOverlayLeftSlim,
		skin.LeftArmOverlayBackRegular,
	}
	armRegularTopComponents []image.Rectangle = []image.Rectangle{
		skin.RightArmTopRegular,
		skin.RightArmBottomRegular,
		skin.LeftArmTopRegular,
		skin.LeftArmBottomRegular,
	}
	armSlimSideComponents []image.Rectangle = []image.Rectangle{
		skin.RightArmFrontSlim,
		skin.RightArmBackSlim,
		skin.LeftArmFrontSlim,
		skin.LeftArmBackSlim,
		skin.RightArmOverlayFrontSlim,
		skin.RightArmOverlayBackSlim,
		skin.LeftArmOverlayFrontSlim,
		skin.LeftArmOverlayBackSlim,
	}
	armSlimTopComponents []image.Rectangle = []image.Rectangle{
		skin.RightArmTopSlim,
		skin.RightArmBottomSlim,
		skin.LeftArmTopSlim,
		skin.LeftArmBottomSlim,
	}
	legSideComponents []image.Rectangle = []image.Rectangle{
		skin.LeftLegRight,
		skin.LeftLegFront,
		skin.LeftLegLeft,
		skin.LeftLegBack,
		skin.LeftArmRight,
		skin.RightLegRight,
		skin.RightLegFront,
		skin.RightLegLeft,
		skin.RightLegBack,
		skin.LeftLegOverlayRight,
		skin.LeftLegOverlayFront,
		skin.LeftLegOverlayLeft,
		skin.LeftLegOverlayBack,
		skin.RightLegOverlayRight,
		skin.RightLegOverlayFront,
		skin.RightLegOverlayLeft,
		skin.RightLegOverlayBack,
	}
	torsoFrontComponents []image.Rectangle = []image.Rectangle{
		skin.TorsoFront,
		skin.TorsoBack,
		skin.TorsoOverlayFront,
		skin.TorsoOverlayBack,
	}
	torsoSideComponents []image.Rectangle = []image.Rectangle{
		skin.TorsoRight,
		skin.TorsoLeft,
		skin.TorsoOverlayRight,
		skin.TorsoOverlayLeft,
	}
	torsoTopComponents []image.Rectangle = []image.Rectangle{
		skin.TorsoTop,
		skin.TorsoBottom,
		skin.TorsoOverlayTop,
		skin.TorsoOverlayBottom,
	}
)

func TestHeadComponents(t *testing.T) {
	for k, c := range headComponents {
		if c.Dx() == 8 && c.Dy() == 8 {
			continue
		}

		t.Fatalf("head component %d has invalid dimensions: expected=(8,8) received=%s", k, c.Size())
	}
}

func TestRegularArmSideComponents(t *testing.T) {
	for k, c := range armRegularSideComponents {
		if c.Dx() == 4 && c.Dy() == 12 {
			continue
		}

		t.Fatalf("regular arm side component %d has invalid dimensions: expected=(4,12) received=%s", k, c.Size())
	}
}

func TestRegularArmTopComponents(t *testing.T) {
	for k, c := range armRegularTopComponents {
		if c.Dx() == 4 && c.Dy() == 4 {
			continue
		}

		t.Fatalf("regular arm top component %d has invalid dimensions: expected=(4,4) received=%s", k, c.Size())
	}
}

func TestSlimArmSideComponents(t *testing.T) {
	for k, c := range armSlimSideComponents {
		if c.Dx() == 3 && c.Dy() == 12 {
			continue
		}

		t.Fatalf("slim arm side component %d has invalid dimensions: expected=(3,12) received=%s", k, c.Size())
	}
}

func TestSlimArmTopComponents(t *testing.T) {
	for k, c := range armSlimTopComponents {
		if c.Dx() == 3 && c.Dy() == 4 {
			continue
		}

		t.Fatalf("slim arm top component %d has invalid dimensions: expected=(3,4) received=%s", k, c.Size())
	}
}

func TestLegComponents(t *testing.T) {
	for k, c := range legSideComponents {
		if c.Dx() == 4 && c.Dy() == 12 {
			continue
		}

		t.Fatalf("leg side component %d has invalid dimensions: expected=(4,12) received=%s", k, c.Size())
	}
}

func TestTorsoFrontComponents(t *testing.T) {
	for k, c := range torsoFrontComponents {
		if c.Dx() == 8 && c.Dy() == 12 {
			continue
		}

		t.Fatalf("torso front component %d has invalid dimensions: expected=(8,12) received=%s", k, c.Size())
	}
}

func TestTorsoSideComponents(t *testing.T) {
	for k, c := range torsoSideComponents {
		if c.Dx() == 4 && c.Dy() == 12 {
			continue
		}

		t.Fatalf("torso side component %d has invalid dimensions: expected=(4,12) received=%s", k, c.Size())
	}
}

func TestTorsoTopComponents(t *testing.T) {
	for k, c := range torsoTopComponents {
		if c.Dx() == 8 && c.Dy() == 4 {
			continue
		}

		t.Fatalf("torso top component %d has invalid dimensions: expected=(8,4) received=%s", k, c.Size())
	}
}
