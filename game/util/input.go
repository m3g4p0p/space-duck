package util

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi/features/math"
)

func TouchPosition() (math.Vec2, bool) {
	var x, y int
	var ok bool

	touchIDs := ebiten.AppendTouchIDs(nil)
	isTouch := len(touchIDs) > 0

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y = ebiten.CursorPosition()
		ok = true
	}

	if isTouch {
		x, y = ebiten.TouchPosition(touchIDs[0])
		ok = true
	}

	return math.NewVec2(float64(x), float64(y)), ok
}
