package game

import (
	"log/slog"
	"sync"

	"m3g4p0p/spring/game/factory"
	"m3g4p0p/spring/game/system"

	"github.com/yohamta/donburi"
	ecslib "github.com/yohamta/donburi/ecs"
	eventslib "github.com/yohamta/donburi/features/events"

	"github.com/hajimehoshi/ebiten/v2"
)

const scale = 1

const (
	LayerBackground ecslib.LayerID = iota
	LayerMain
	LayerUI
)

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
	ecs.AddRenderer(LayerMain, system.NewRenderSystem().Draw)
	ecs.AddRenderer(LayerUI, FlushLogs)

	factory.CreatePlayer(world)

	return game
}
