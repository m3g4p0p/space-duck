package factory

import (
	"m3g4p0p/spring/game/component"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
)

func CreateText(world donburi.World, text string, pos math.Vec2) *donburi.Entry {
	entry := world.Entry(world.Create(
		component.Text,
		transform.Transform,
	))

	component.Text.Set(entry, &component.TextData{
		Text: text,
		Size: 24,
	})

	return entry
}
