package component

import (
	"github.com/charmbracelet/harmonica"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
)

type TiltData struct {
	AbsMax float64
}

type TextData struct {
	Text string
	Size float64
	Pos  math.Vec2
}

type SpringData struct {
	Pos, Vel, Eq float64
}

type PlayerData struct {
	Score int
}

var (
	Sprite     = donburi.NewComponentType[ebiten.Image]()
	Target     = donburi.NewComponentType[math.Vec2]()
	Velocity   = donburi.NewComponentType[math.Vec2]()
	Tilt       = donburi.NewComponentType[TiltData]()
	Text       = donburi.NewComponentType[TextData]()
	Spring     = donburi.NewComponentType[SpringData]()
	Projectile = donburi.NewComponentType[harmonica.Projectile]()
	Player     = donburi.NewComponentType[PlayerData]()
	Alpha      = donburi.NewComponentType[float32]()
	Background = donburi.NewTag("background")
)
