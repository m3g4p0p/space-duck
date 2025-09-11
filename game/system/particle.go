package system

import (
	"m3g4p0p/spring/game/component"
	"m3g4p0p/spring/game/factory"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
)

type ParticleSystem struct {
	query *donburi.Query
}

func NewParticleSystem() ParticleSystem {
	return ParticleSystem{donburi.NewQuery(filter.Contains(
		component.Sprite,
		component.Projectile,
		transform.Transform,
	))}
}

func (t ParticleSystem) Update(ecs *ecs.ECS) {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		factory.CreateParticle(ecs.World, 10)
	}

	for entry := range t.query.Iter(ecs.World) {
		pos := component.Projectile.Get(entry).Update()
		transform.SetWorldPosition(entry, math.NewVec2(pos.X, pos.Y))
	}
}
