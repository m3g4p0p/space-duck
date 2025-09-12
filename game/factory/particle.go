package factory

import (
	"math/rand"

	"m3g4p0p/spring/game/component"
	"m3g4p0p/spring/game/util"

	"github.com/charmbracelet/harmonica"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
)

func CreateParticle(world donburi.World, size int) *donburi.Entry {
	entry := world.Entry(world.Create(
		component.Sprite,
		component.Alpha,
		component.Projectile,
		transform.Transform,
	))

	component.Sprite.Set(entry, util.CreateCircle(size))
	component.Alpha.SetValue(entry, 1)

	width, _ := util.ClientSize()

	component.Projectile.Set(entry, harmonica.NewProjectile(
		harmonica.FPS(ebiten.TPS()),
		harmonica.Point{X: float64(width) * rand.Float64()},
		harmonica.Vector{Y: 100},
		harmonica.Vector{Y: 100},
	))

	return entry
}
