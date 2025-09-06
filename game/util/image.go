package util

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/yohamta/donburi/features/math"
)

func CreateCircle(size int) *ebiten.Image {
	img := ebiten.NewImage(size, size)
	vector.DrawFilledCircle(
		img,
		float32(size)/2,
		float32(size)/2,
		float32(size)/2,
		color.White,
		true,
	)
	return img
}

func ImageSize(image *ebiten.Image) math.Vec2 {
	return math.NewVec2(
		float64(image.Bounds().Dx()),
		float64(image.Bounds().Dy()),
	)
}
