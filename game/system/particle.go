package system

import (
	"log/slog"
	"math"
	"math/rand/v2"

	"m3g4p0p/spring/game/component"
	"m3g4p0p/spring/game/factory"
	"m3g4p0p/spring/game/util"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	mathlib "github.com/yohamta/donburi/features/math"
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
		pos := component.Projectile.Get(entry).Update()

		if height < int(pos.Y) {
			ecs.World.Remove(entry.Entity())
			return
		}

		transform.SetWorldPosition(entry, mathlib.NewVec2(pos.X, pos.Y))
		alpha := 0.5 + math.Sin(pos.Y/10)/2
		component.Alpha.SetValue(entry, float32(alpha))
	}
}
