package system

import (
	"bytes"

	"m3g4p0p/spring/game/component"

	"github.com/charmbracelet/harmonica"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
)

var fontSource *text.GoTextFaceSource

type UISystem struct {
	query  *donburi.Query
	spring harmonica.Spring
}

func init() {
	var err error

	fontSource, err = text.NewGoTextFaceSource(bytes.NewReader(fonts.PressStart2P_ttf))
	if err != nil {
		panic(err)
	}
}

func NewUISystem() *UISystem {
	return &UISystem{
		query: donburi.NewQuery(filter.Contains(
			component.Text,
			transform.Transform,
		)),
		spring: harmonica.NewSpring(
			harmonica.FPS(ebiten.TPS()),
			5.0,
			0.5,
		),
	}
}

func (s *UISystem) Update(ecs *ecs.ECS) {}

func (s *UISystem) Draw(ecs *ecs.ECS, image *ebiten.Image) {
	s.query.Each(ecs.World, func(e *donburi.Entry) {
		data := component.Text.GetValue(e)
		drawOpts := ebiten.DrawImageOptions{}
		drawOpts.GeoM.Translate(data.Pos.XY())

		text.Draw(
			image,
			data.Text,
			&text.GoTextFace{
				Source: fontSource,
				Size:   data.Size,
			},
			&text.DrawOptions{
				DrawImageOptions: drawOpts,
				LayoutOptions: text.LayoutOptions{
					PrimaryAlign:   text.AlignCenter,
					SecondaryAlign: text.AlignCenter,
				},
			},
		)
	})
}
