package main

import (
	"main/game"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := &game.Game{
		Width:            600,
		Height:           600,
		Xbound:           [2]float64{-2, 2},
		Ybound:           [2]float64{-2, 2},
		N:                100,
		ProjectionVector: [3]float64{1 / math.Sqrt(2), 0, 1 / math.Sqrt(2)},
	}
	g.Init()
	ebiten.RunGame(g)
}
