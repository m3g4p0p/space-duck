package system

import (
	"bytes"
	"math"

	"m3g4p0p/spring/game/component"

	"github.com/charmbracelet/harmonica"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	mathlib "github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
)

var fontSource *text.GoTextFaceSource

type UISystem struct {
	queryDraw   *donburi.Query
	queryUpdate *donburi.Query
	spring      harmonica.Spring
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
		queryDraw: donburi.NewQuery(filter.Contains(
			component.Text,
			transform.Transform,
		)),
		queryUpdate: donburi.NewQuery(filter.Contains(
			component.Text,
			component.Spring,
			transform.Transform,
		)),
		spring: harmonica.NewSpring(
			harmonica.FPS(ebiten.TPS()),
			7.5,
			1.5,
		),
	}
}

func (s *UISystem) Update(ecs *ecs.ECS) {
	s.queryUpdate.Each(ecs.World, func(e *donburi.Entry) {
		data := component.Spring.Get(e)
		data.Pos, data.Vel = s.spring.Update(data.Pos, data.Vel, data.Eq)
		transform.SetWorldScale(e, mathlib.NewVec2(data.Pos, data.Pos))

		if e.HasComponent(component.Alpha) {
			component.Alpha.SetValue(e, 1-float32(data.Pos/data.Eq))
		}

		if math.Abs(data.Eq-data.Pos)+math.Abs(data.Vel) < 0.01 {
			ecs.World.Remove(e.Entity())
		}
	})
}

func (s *UISystem) Draw(ecs *ecs.ECS, image *ebiten.Image) {
	s.queryDraw.Each(ecs.World, func(e *donburi.Entry) {
		data := component.Text.GetValue(e)
		drawOpts := ebiten.DrawImageOptions{}
		drawOpts.GeoM.Scale(transform.WorldScale(e).XY())
		drawOpts.GeoM.Translate(data.Pos.XY())
		drawOpts.ColorScale.Scale(0, 1, 1, 1)

		if e.HasComponent(component.Alpha) {
			alpha := component.Alpha.GetValue(e)
			drawOpts.ColorScale.ScaleAlpha(alpha)
		}

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
