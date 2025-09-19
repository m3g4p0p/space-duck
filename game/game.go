package game

import (
	"fmt"
	"math/rand"

	"m3g4p0p/spring/game/component"
	"m3g4p0p/spring/game/events"
	"m3g4p0p/spring/game/factory"
	"m3g4p0p/spring/game/scene"
	"m3g4p0p/spring/game/system"
	"m3g4p0p/spring/game/util"

	"github.com/yohamta/donburi"
	ecslib "github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"

	"github.com/hajimehoshi/ebiten/v2"
)

const scale = 1

const (
	LayerBackground ecslib.LayerID = iota
	LayerMain
	LayerUI
)

var isBackground = filter.Contains(component.Background)

func New() ebiten.Game {
	manager := scene.NewManager()

	manager.Register("main", func() *ecslib.ECS {
		world := donburi.NewWorld()
		ecs := ecslib.NewECS(world)
		ui := system.NewUISystem()

		ebiten.SetWindowTitle("Space Duck")

		ecs.AddSystem(system.NewTargetSystem().Update)
		ecs.AddSystem(system.NewTiltSystem().Update)
		ecs.AddSystem(system.NewPlayerSystem().Update)
		ecs.AddSystem(system.NewParticleSystem().Update)
		ecs.AddSystem(system.NewProjectileSystem().Update)
		ecs.AddSystem(ui.Update)

		ecs.AddRenderer(LayerMain, system.NewRenderSystem(isBackground).Draw)
		ecs.AddRenderer(LayerMain, system.NewRenderSystem(filter.Not(isBackground)).Draw)
		ecs.AddRenderer(LayerUI, FlushLogs)
		ecs.AddRenderer(LayerUI, ui.Draw)

		if !util.IsBrowser {
			ebiten.SetWindowSize(400, 800)
		}

		width, height := util.ClientSize()
		for range 20 {
			factory.CreateParticle(
				world,
				rand.Intn(width),
				rand.Intn(height),
				util.RandIntMN(2, 5),
			)
		}

		factory.CreatePlayer(world, util.Vec2FromInt(width/2, height/2))

		events.ScoreEvent.Subscribe(world, func(w donburi.World, event int) {
			size := util.ClientSizeVec2()
			size.X /= 2
			size.Y -= 60
			factory.CreateText(w, fmt.Sprint(event), 30, size)
		})

		return ecs
	})

	manager.Goto("main")
	return manager
}
