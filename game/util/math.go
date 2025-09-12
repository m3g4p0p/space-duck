package util

import "github.com/yohamta/donburi/features/math"

func Vec2FromInt(x, y int) math.Vec2 {
	return math.NewVec2(float64(x), float64(y))
}
