//go:build ignore

package main

import (
	"log"
	"m3g4p0p/spring/game/scene"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	ecslib "github.com/yohamta/donburi/ecs"
)

func introScene() *ecs.ECS {
	world := donburi.NewWorld()
	ecs := ecslib.NewECS(world)

	ecs.AddRenderer(ecslib.LayerDefault, func(ecs *ecslib.ECS, image *ebiten.Image) {
		ebitenutil.DebugPrint(image, "hello")
	})

	return ecs
}

func mainScene() *ecs.ECS {
	world := donburi.NewWorld()
	ecs := ecslib.NewECS(world)
	return ecs
}

func main() {
	game := scene.NewManager(
		scene.WithScene("intro", introScene()),
		scene.WithScene("main", mainScene()),
		scene.WithInitial("intro"),
	)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
