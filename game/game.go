package game

import (
	"sync"

	"m3g4p0p/spring/game/component"
	"m3g4p0p/spring/game/system"
	"m3g4p0p/spring/game/util"

	"github.com/yohamta/donburi"
	ecslib "github.com/yohamta/donburi/ecs"
	eventslib "github.com/yohamta/donburi/features/events"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

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
	return outsideWidth, outsideHeight
}

func New() *Game {
	world := donburi.NewWorld()
	ecs := ecslib.NewECS(world)
	game := &Game{ECS: ecs}

	ebiten.SetWindowTitle("Hello, World!")
	ecs.AddRenderer(ecslib.LayerDefault, system.NewRenderSystem().Draw)

	entry := world.Entry(world.Create(component.Sprite, transform.Transform))
	component.Sprite.Set(entry, util.CreateCircle(10))
	transform.Transform.SetValue(entry, transform.TransformData{
		LocalPosition: math.NewVec2(100, 100),
	})

	return game
}
