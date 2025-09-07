package factory

import (
	"m3g4p0p/spring/game/component"
	"m3g4p0p/spring/game/util"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
)

func CreateParticle(world donburi.World, size int) *donburi.Entry {
	entry := world.Entry(world.Create(
		component.Sprite,
		component.Alpha,
		transform.Transform,
	))

	component.Sprite.Set(entry, util.CreateCircle(size))

	return entry
}
