package system

import (
	"m3g4p0p/spring/game/component"

	"github.com/charmbracelet/harmonica"
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
		projectile := component.Projectile.Get(entry)
		delta := projectileDelta(projectile)
		pos := transform.WorldPosition(entry)

		transform.SetWorldPosition(entry, pos.Add(delta))
	}
}

func projectileDelta(p *harmonica.Projectile) mathlib.Vec2 {
	current := p.Position()
	next := p.Update()

	return mathlib.NewVec2(next.X-current.X, next.Y-current.Y)
}
