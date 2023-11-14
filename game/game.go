package game

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	BoundaryPoints            [][3]float64
	HemisphereLines           [][][3]float64
	ProjectionVector          [3]float64
	ProjectedBoundaryPoints   [][2]float64
	ProjectedHemispherePoints [][][2]float64
	Width, Height             int
	Xbound, Ybound            [2]float64
	BoundaryXYs               [][2]int // xy coordinate for the pixel
	HemisphereXYs             [][][2]int
	N                         int // number of hemisphere lines
}

func (g *Game) Init() {
	radius := 1.0
	for i := 0; i < g.N; i++ {
		theta := 2 * math.Pi / float64(g.N) * float64(i)
		point := [3]float64{radius * math.Cos(theta), radius * math.Sin(theta), 0}
		g.BoundaryPoints = append(g.BoundaryPoints, point)
	}
	g.ProjectedBoundaryPoints = make([][2]float64, g.N)
	g.ProjectedHemispherePoints = make([][][2]float64, g.N)
	g.BoundaryXYs = make([][2]int, g.N)
	g.HemisphereXYs = make([][][2]int, g.N)
	for i := 0; i < g.N; i++ {
		g.ProjectedHemispherePoints[i] = make([][2]float64, g.N)
		g.HemisphereXYs[i] = make([][2]int, g.N)
	}
	for i := 0; i < g.N; i++ {
		hemispherePoints := [][3]float64{}
		theta := 2 * math.Pi / float64(g.N) * float64(i)
		for j := 0; j < g.N; j++ {
			phi := math.Pi / 2 / float64(g.N) * float64(j)
			hemispherePoints = append(hemispherePoints, [3]float64{
				radius * math.Sin(phi) * math.Cos(theta),
				radius * math.Sin(phi) * math.Sin(theta),
				radius * math.Cos(phi),
			})
		}
		g.HemisphereLines = append(g.HemisphereLines, hemispherePoints)
	}
}

func Projection(point [3]float64, projectionVector [3]float64) [2]float64 {
	basis := [2][3]float64{
		{0, 1, 0},
		{-1 / math.Sqrt(2), 0, 1 / math.Sqrt(2)},
	}
	return [2]float64{
		InnerProduct(point, basis[0]),
		InnerProduct(point, basis[1]),
	}
}

func InnerProduct(point, basis [3]float64) float64 {
	var inner float64
	for i := 0; i < 3; i++ {
		inner += point[i] * basis[i]
	}
	return inner
}

func (g *Game) Update() error {
	for i, point := range g.BoundaryPoints {
		g.ProjectedBoundaryPoints[i] = Projection(point, g.ProjectionVector)
	}
	for i, hline := range g.HemisphereLines {
		for j, point := range hline {
			g.ProjectedHemispherePoints[i][j] = Projection(point, g.ProjectionVector)
		}
	}
	for i, point := range g.ProjectedBoundaryPoints {
		g.BoundaryXYs[i] = [2]int{
			TranslateXcoord(point[0], g.Xbound, g.Width),
			TranslateYcoord(point[1], g.Ybound, g.Height),
		}
	}
	for i, line := range g.ProjectedHemispherePoints {
		for j, point := range line {
			g.HemisphereXYs[i][j] = [2]int{
				TranslateXcoord(point[0], g.Xbound, g.Width),
				TranslateYcoord(point[1], g.Ybound, g.Height),
			}
		}
	}
	return nil
}

func TranslateXcoord(x float64, xbound [2]float64, width int) int {
	return int((x - xbound[0]) / (xbound[1] - xbound[0]) * float64(width))
}

func TranslateYcoord(y float64, ybound [2]float64, height int) int {
	return int((ybound[1] - y) / (ybound[1] - ybound[0]) * float64(height))
}

func (g *Game) Draw(screen *ebiten.Image) {
	for i := 0; i < g.N; i++ {
		point1, point2 := g.BoundaryXYs[i], g.BoundaryXYs[(i+1)%g.N]
		line := LinePoints(point1, point2)
		for _, xy := range line {
			screen.Set(xy[0], xy[1], color.White)
		}
	}
	for _, line := range g.HemisphereXYs {
		for j := 0; j < g.N-1; j++ {
			point1, point2 := line[j], line[(j+1)%g.N]
			line := LinePoints(point1, point2)
			for _, xy := range line {
				screen.Set(xy[0], xy[1], color.White)
			}
		}
	}
}

func LinePoints(point1, point2 [2]int) [][2]int {
	x1, y1 := point1[0], point1[1]
	x2, y2 := point2[0], point2[1]
	w, h := x1-x2, y1-y2
	if w < 0 {
		w = -w
	}
	if h < 0 {
		h = -h
	}
	n := w
	if n < h {
		n = h
	}
	line := [][2]int{}
	for i := 0; i <= n; i++ {
		line = append(line, [2]int{
			x1 + int(float64(x2-x1)/float64(n)*float64(i)),
			y1 + int(float64(y2-y1)/float64(n)*float64(i)),
		})
	}
	return line
}

func (g *Game) Layout(w, h int) (int, int) {
	return 600, 600
}
