package game

import (
	"log/slog"
	"math/rand"
	"sync"

	"m3g4p0p/spring/game/component"
	"m3g4p0p/spring/game/factory"
	"m3g4p0p/spring/game/system"
	"m3g4p0p/spring/game/util"

	"github.com/yohamta/donburi"
	ecslib "github.com/yohamta/donburi/ecs"
	eventslib "github.com/yohamta/donburi/features/events"
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

type Game struct {
	*ecslib.ECS
	once sync.Once
}

func (g *Game) Update() error {
	g.ECS.Update()
	eventslib.ProcessAllEvents(g.World)
	slog.Info("hello world")
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.ECS.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth * scale, outsideHeight * scale
}

func New() *Game {
	world := donburi.NewWorld()
	ecs := ecslib.NewECS(world)
	game := &Game{ECS: ecs}

	ebiten.SetWindowTitle("Space Duck")
	ecs.AddSystem(system.NewTargetSystem().Update)
	ecs.AddSystem(system.NewTiltSystem().Update)
	ecs.AddSystem(system.NewPlayerSystem().Update)
	ecs.AddSystem(system.NewParticleSystem().Update)
	ecs.AddRenderer(LayerMain, system.NewRenderSystem(isBackground).Draw)
	ecs.AddRenderer(LayerMain, system.NewRenderSystem(filter.Not(isBackground)).Draw)
	ecs.AddRenderer(LayerUI, FlushLogs)

	factory.CreatePlayer(world)

	width, height := util.ClientSize()
	for range 20 {
		factory.CreateParticle(
			world,
			rand.Intn(width),
			rand.Intn(height),
			util.RandIntMN(2, 5),
		)
	}

	return game
}
