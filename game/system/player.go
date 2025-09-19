package system

import (
	"m3g4p0p/spring/game/component"
	"m3g4p0p/spring/game/events"
	"m3g4p0p/spring/game/util"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
)

type PlayerSystem struct {
	query     *donburi.Query
	isTouched bool
}

func NewPlayerSystem() *PlayerSystem {
	return &PlayerSystem{
		query: donburi.NewQuery(filter.Contains(
			component.Player,
			component.Target,
		)),
	}
}

func (t *PlayerSystem) Update(ecs *ecs.ECS) {
	pos, ok := util.TouchPosition()

	for entry := range t.query.Iter(ecs.World) {
		if ok && !t.isTouched {
			component.Target.SetValue(entry, pos)

			data := component.Player.Get(entry)
			data.Score++
			events.ScoreEvent.Publish(ecs.World, data.Score)
		}

		t.isTouched = ok
	}
}
