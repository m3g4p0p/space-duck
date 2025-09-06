//go:build !js || !wasm

package util

import "github.com/hajimehoshi/ebiten/v2"

func GetClientSize() (int, int) {
	return ebiten.WindowSize()
}
