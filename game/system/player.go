package system

import (
	"m3g4p0p/spring/game/component"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/filter"
)

type PlayerSystem struct {
	query *donburi.Query
}

func NewPlayerSystem() PlayerSystem {
	return PlayerSystem{donburi.NewQuery(filter.Contains(
		component.Player,
		component.Target,
	))}
}

func (t PlayerSystem) Update(ecs *ecs.ECS) {
	if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		return
	}

	x, y := ebiten.CursorPosition()

	for entry := range t.query.Iter(ecs.World) {
		component.Target.Set(entry, &component.TargetData{
			Vec2: math.NewVec2(float64(x), float64(y)),
		})
	}
}
