package system

import (
	"bytes"

	"m3g4p0p/spring/game/util"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/yohamta/donburi/ecs"
)

var fontSource *text.GoTextFaceSource

type UISystem struct{}

func init() {
	var err error

	fontSource, err = text.NewGoTextFaceSource(bytes.NewReader(fonts.PressStart2P_ttf))
	if err != nil {
		panic(err)
	}
}

func (s *UISystem) Draw(ecs *ecs.ECS, image *ebiten.Image) {
	drawOpts := ebiten.DrawImageOptions{}
	center := util.CLientSizeVec2().DivScalar(2)
	drawOpts.GeoM.Translate(center.XY())

	text.Draw(
		image,
		"hello",
		&text.GoTextFace{Source: fontSource, Size: 30},
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
