//go:build !js || !wasm

package util

import "github.com/hajimehoshi/ebiten/v2"

func ClientSize() (int, int) {
	return ebiten.WindowSize()
}
