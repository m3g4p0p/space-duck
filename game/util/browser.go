//go:build js && wasm

package util

import (
	"syscall/js"
)

var (
	clientWidth, clientHeight = getDimensions()
	window                    = js.Global().Get("window")
)

func getDimensions() (int, int) {
	width := window.Get("innerWidth").Int()
	height := window.Get("innerHeight").Int()

	return width, height
}

func init() {
	window.Call("addEventListener", "resize", js.FuncOf(func(
		this js.Value,
		args []js.Value,
	) any {
		clientWidth, clientHeight = getDimensions()
		return nil
	}))
}

func ClientSize() (int, int) {
	return clientWidth, clientHeight
}
