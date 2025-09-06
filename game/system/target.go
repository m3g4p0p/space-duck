package system

import (
	"m3g4p0p/spring/game/component"

	"github.com/charmbracelet/harmonica"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
)

type TargetSystem struct {
	query  *donburi.Query
	spring harmonica.Spring
}

func NewTargetSystem() TargetSystem {
	return TargetSystem{
		donburi.NewQuery(filter.Contains(
			component.Target,
			transform.Transform,
		)),
		harmonica.NewSpring(harmonica.FPS(ebiten.TPS()), 5.0, 0.2),
	}
}

func (t TargetSystem) Update(ecs *ecs.ECS) {
	for entry := range t.query.Iter(ecs.World) {
		target := component.Target.Get(entry)
		pos := transform.WorldPosition(entry)
		vel := component.Velocity.Get(entry)
		pos.X, vel.X = t.spring.Update(pos.X, vel.X, target.X)
		pos.Y, vel.Y = t.spring.Update(pos.Y, vel.Y, target.Y)

		transform.SetWorldPosition(entry, pos)
	}
}
