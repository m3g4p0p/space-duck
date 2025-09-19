package factory

import (
	"image/color"

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
		component.Color,
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

	component.Color.Set(entry, &component.ColorData{
		Color: color.RGBA{G: 255, B: 255},
	})

	return entry
}
