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

func CreatePlayer(world donburi.World, pos math.Vec2) *donburi.Entry {
	entry := world.Entry(world.Create(
		component.Sprite,
		component.Player,
		component.Target,
		component.Velocity,
		component.Tilt,
		transform.Transform,
	))

	tilt := component.TiltData{
		AbsMax: math.ToRadians(45),
	}

	component.Sprite.Set(entry, playerSprite)
	component.Tilt.SetValue(entry, tilt)
	component.Target.Set(entry, &component.TargetData{Vec2: pos})
	transform.SetWorldScale(entry, math.NewVec2(0.5, 0.5))
	transform.SetWorldPosition(entry, pos)

	return entry
}
