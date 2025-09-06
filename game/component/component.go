package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

// var MessageComponent = donburi.NewComponentType[pubsub.Message]()
var Sprite = donburi.NewComponentType[ebiten.Image]()
