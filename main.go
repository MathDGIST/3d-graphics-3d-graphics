package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	size := 5
	for i := -size; i <= size; i++ {
		for j := -size; j <= size; j++ {
			screen.Set(300+i, 300+j, color.White)
		}
	}
}

func (g *Game) Layout(w, h int) (int, int) {
	return 600, 600
}

func main() {
	ebiten.RunGame(&Game{})
}
