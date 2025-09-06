package main

import (
	"log"

	"m3g4p0p/spring/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	if err := ebiten.RunGame(game.New()); err != nil {
		log.Fatal(err)
	}
}
