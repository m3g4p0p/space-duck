package system

import (
	"m3g4p0p/spring/game/component"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
)

type TargetSystem struct {
	query *donburi.Query
}

func NewTargetSystem() TargetSystem {
	return TargetSystem{donburi.NewQuery(filter.Contains(
		component.Target,
		transform.Transform,
	))}
}

func (t TargetSystem) Update(ecs *ecs.ECS) {
	for entry := range t.query.Iter(ecs.World) {
		target := component.Target.Get(entry)
		transform.SetWorldPosition(entry, target.Vec2)
	}
}
