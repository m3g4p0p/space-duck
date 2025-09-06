package game

import (
	"sync"

	"m3g4p0p/spring/game/factory"
	"m3g4p0p/spring/game/system"

	"github.com/yohamta/donburi"
	ecslib "github.com/yohamta/donburi/ecs"
	eventslib "github.com/yohamta/donburi/features/events"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const scale = 2

type Game struct {
	*ecslib.ECS
	once sync.Once
}

func (g *Game) Update() error {
	g.ECS.Update()
	eventslib.ProcessAllEvents(g.World)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.ECS.Draw(screen)
	ebitenutil.DebugPrint(screen, "Hello, World!")
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

	factory.CreatePlayer(world)

	return game
}
