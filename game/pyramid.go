package game

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const DeltaTheta = 1. / 360 * (2 * math.Pi)
const ScreenSize = 600
const Pad = 10

var Connect = [6][2]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}, {2, 3}}

type Pyramid struct {
	Vertices [4]Vector
}

func NewGame() *Pyramid {
	return &Pyramid{
		Vertices: [4]Vector{
			{2 * math.Sqrt(2) / 3., 0, -1 / 3.},
			{-math.Sqrt(2) / 3., math.Sqrt(6) / 3, -1 / 3.},
			{-math.Sqrt(2) / 3., -math.Sqrt(6) / 3, -1 / 3.},
			{0, 0, 1}},
	}
}

func (g *Pyramid) Update() error {
	keyPressed := false
	var m Matrix
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		m = LongitudeRotate(DeltaTheta)
		keyPressed = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		m = LongitudeRotate(-DeltaTheta)
		keyPressed = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		m = LatitudeRotate(DeltaTheta)
		keyPressed = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		m = LatitudeRotate(-DeltaTheta)
		keyPressed = true
	}
	if keyPressed {
		for i, v := range g.Vertices {
			g.Vertices[i] = m.MulVec(v)
		}
	}
	return nil
}

func (g *Pyramid) Draw(screen *ebiten.Image) {
	for _, c := range Connect {
		var xy [2][2]int
		for i := 0; i < 2; i++ {
			for j := 0; j < 2; j++ {
				xy[i][j] = Pad + int((g.Vertices[c[i]][j]+1)/2*(ScreenSize-2*Pad))
			}
		}
		step := Max(xy[0][0]-xy[1][0], xy[0][1]-xy[1][1])
		for k := 0; k <= step; k++ {
			var p [2]int
			for j := 0; j < 2; j++ {
				p[j] = xy[0][j] + int(float64(k)*float64(xy[1][j]-xy[0][j])/float64(step))
			}
			screen.Set(p[0], p[1], color.RGBA{255, 255, 255, 255})
		}
	}
}

func (g *Pyramid) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenSize, ScreenSize
}

func Max(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	if a < b {
		return b
	}
	return a
}
