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

	ebiten.SetWindowTitle("Hello, World!")
	ecs.AddSystem(system.NewTargetSystem().Update)
	ecs.AddSystem(system.NewPlayerSystem().Update)
	ecs.AddRenderer(ecslib.LayerDefault, system.NewRenderSystem().Draw)
	ecs.AddRenderer(ecslib.LayerDefault, FlushLogs)

	factory.CreatePlayer(world)

	return game
}
