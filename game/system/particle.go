package system

import (
	"log/slog"
	"math/rand/v2"

	"m3g4p0p/spring/game/component"
	"m3g4p0p/spring/game/factory"
	"m3g4p0p/spring/game/util"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
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
	slog.Info("particles", slog.Int("count", t.query.Count(ecs.World)))
	width, height := util.ClientSize()

	if rand.Float64() < 10.0/float64(ebiten.TPS()) {
		factory.CreateParticle(
			ecs.World,
			rand.IntN(width),
			0,
			util.RandIntMN(2, 5),
		)
	}

	for entry := range t.query.Iter(ecs.World) {
		pos := transform.WorldPosition(entry)

		if height < int(pos.Y) {
			ecs.World.Remove(entry.Entity())
			return
		}

		alpha := 1 - pos.Y/float64(height)
		component.Alpha.SetValue(entry, float32(alpha))
	}
}
