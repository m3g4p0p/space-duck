package factory

import (
	"m3g4p0p/spring/assets"
	"m3g4p0p/spring/game/component"
	"m3g4p0p/spring/game/util"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
)

var playerSprite = util.Must(util.LoadImage(assets.FS, "PNG/shipPink_manned.png"))

func CreatePlayer(world donburi.World) *donburi.Entry {
	entry := world.Entry(world.Create(
		component.Sprite,
		component.Player,
		component.Target,
		component.Velocity,
		component.Tilt,
		transform.Transform,
	))

	component.Sprite.Set(entry, playerSprite)
	transform.SetWorldScale(entry, math.NewVec2(0.5, 0.5))
	return entry
}
