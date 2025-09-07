package system

import (
	"m3g4p0p/spring/game/component"
	"m3g4p0p/spring/game/util"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
)

type RenderSystem struct {
	query *donburi.Query
}

func (s RenderSystem) Draw(ecs *ecs.ECS, image *ebiten.Image) {
	s.query.Each(ecs.World, func(e *donburi.Entry) {
		var angle float64
		sprite := component.Sprite.Get(e)
		op := &ebiten.DrawImageOptions{}

		op.GeoM.Translate(util.ImageSize(sprite).DivScalar(-2).XY())
		op.GeoM.Rotate(transform.WorldRotation(e) + angle)
		op.GeoM.Scale(transform.WorldScale(e).XY())
		op.GeoM.Translate(transform.WorldPosition(e).XY())

		aplha := util.GetValue(e, component.Alpha, 1)
		op.ColorScale.ScaleAlpha(aplha)

		image.DrawImage(sprite, op)
	})
}

func NewRenderSystem(components ...donburi.IComponentType) RenderSystem {
	components = append(components, component.Sprite, transform.Transform)

	return RenderSystem{
		query: donburi.NewQuery(filter.Contains(components...)),
	}
}
