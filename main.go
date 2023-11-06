package main

import (
	"main/game"

	ebiten "github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := game.NewGame()
	ebiten.RunGame(g)
}
