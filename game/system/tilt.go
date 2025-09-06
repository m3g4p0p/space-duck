package system

import (
	"m3g4p0p/spring/game/component"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
)

type TiltSystem struct {
	query *donburi.Query
}

func NewTiltSystem() TiltSystem {
	return TiltSystem{
		donburi.NewQuery(filter.Contains(
			component.Tilt,
			component.Velocity,
			transform.Transform,
		)),
	}
}

func (t TiltSystem) Update(ecs *ecs.ECS) {
	for entry := range t.query.Iter(ecs.World) {
		vel := component.Velocity.Get(entry)
		maxTilt := component.Tilt.Get(entry).AbsMax
		angle := max(-maxTilt, min(maxTilt, vel.X/1000))

		transform.SetWorldRotation(entry, angle)
	}
}
