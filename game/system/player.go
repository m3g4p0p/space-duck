package system

import (
	"m3g4p0p/spring/game/component"
	"m3g4p0p/spring/game/util"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
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
	pos, ok := util.ActivePosition()

	if !ok {
		return
	}

	for entry := range t.query.Iter(ecs.World) {
		component.Target.Set(entry, &component.TargetData{
			Vec2: pos,
		})
	}
}
