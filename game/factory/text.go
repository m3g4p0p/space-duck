package factory

import (
	"m3g4p0p/spring/game/component"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
)

func CreateFCT(
	world donburi.World,
	text string,
	size float64,
	pos math.Vec2,
) *donburi.Entry {
	entry := world.Entry(world.Create(
		component.Text,
		component.Alpha,
		component.Spring,
		transform.Transform,
	))

	component.Text.Set(entry, &component.TextData{
		Text: text,
		Size: size,
		Pos:  pos,
	})

	component.Spring.Set(entry, &component.SpringData{
		Eq: 2,
	})

	return entry
}
