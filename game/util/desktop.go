//go:build !js || !wasm

package util

import "github.com/hajimehoshi/ebiten/v2"

var IsBrowser = false

func ClientSize() (int, int) {
	return ebiten.WindowSize()
}
