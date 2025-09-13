package system

import (
	"bytes"
	"log"

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

	text.Draw(image, "hello", &text.GoTextFace{Source: f, Size: 30}, &text.DrawOptions{})
}

func NewUISystem() *UISystem {
	return &UISystem{}
}
