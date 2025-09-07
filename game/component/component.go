package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
)

type TargetData struct {
	math.Vec2
}

type VelocityData struct {
	math.Vec2
}

type TiltData struct {
	AbsMax float64
}

// var MessageComponent = donburi.NewComponentType[pubsub.Message]()
var (
	Sprite     = donburi.NewComponentType[ebiten.Image]()
	Target     = donburi.NewComponentType[TargetData]()
	Velocity   = donburi.NewComponentType[VelocityData]()
	Tilt       = donburi.NewComponentType[TiltData]()
	Alpha      = donburi.NewComponentType[float32]()
	Background = donburi.NewTag("background")
	Player     = donburi.NewTag("player")
)
