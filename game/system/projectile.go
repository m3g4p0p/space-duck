package system

import (
	"m3g4p0p/spring/game/component"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	mathlib "github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
)

type ProjectileSystem struct {
	query *donburi.Query
}

func NewProjectileSystem() ProjectileSystem {
	return ProjectileSystem{donburi.NewQuery(filter.Contains(
		component.Projectile,
		transform.Transform,
	))}
}

func (t ProjectileSystem) Update(ecs *ecs.ECS) {
	for entry := range t.query.Iter(ecs.World) {
		pos := component.Projectile.Get(entry).Update()
		transform.SetWorldPosition(entry, mathlib.NewVec2(pos.X, pos.Y))
	}
}
