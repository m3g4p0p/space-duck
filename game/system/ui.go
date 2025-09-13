package system

import (
	"bytes"
	"log"

	"m3g4p0p/spring/game/util"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/yohamta/donburi/ecs"
)

type UISystem struct{}

func (s *UISystem) Draw(ecs *ecs.ECS, image *ebiten.Image) {
	f, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.PressStart2P_ttf))
	if err != nil {
		log.Fatal(err)
	}

	drawOpts := ebiten.DrawImageOptions{}
	drawOpts.GeoM.Translate(util.CLientSizeVec2().DivScalar(2).XY())

	text.Draw(
		image,
		"hello",
		&text.GoTextFace{Source: f, Size: 30},
		&text.DrawOptions{
			DrawImageOptions: drawOpts,
			LayoutOptions: text.LayoutOptions{
				PrimaryAlign:   text.AlignCenter,
				SecondaryAlign: text.AlignCenter,
			},
		},
	)
}

func NewUISystem() *UISystem {
	return &UISystem{}
}
