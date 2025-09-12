package factory

import (
	"m3g4p0p/spring/game/component"
	"m3g4p0p/spring/game/util"

	"github.com/charmbracelet/harmonica"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
)

func CreateParticle(world donburi.World, x, y, size int) *donburi.Entry {
	entry := world.Entry(world.Create(
		component.Sprite,
		component.Alpha,
		component.Projectile,
		component.Background,
		transform.Transform,
	))

	component.Sprite.Set(entry, util.CreateCircle(size))
	transform.SetWorldPosition(entry, util.Vec2FromInt(x, y))

	component.Projectile.Set(entry, harmonica.NewProjectile(
		harmonica.FPS(ebiten.TPS()),
		harmonica.Point{},
		harmonica.Vector{Y: 100},
		harmonica.Vector{Y: 100},
	))

	return entry
}
