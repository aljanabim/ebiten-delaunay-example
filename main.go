package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/fogleman/delaunay"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	points  []delaunay.Point
	indices []int
}

func (g *Game) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		x, y := ebiten.CursorPosition()
		g.points = append(g.points, delaunay.Point{X: float64(x), Y: float64(y)})
		triangulation, err := delaunay.Triangulate(g.points)
		if err != nil {
			fmt.Println(err)
		}
		g.indices = triangulation.Triangles
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)

	for _, point := range g.points {
		vector.DrawFilledCircle(screen, float32(point.X), float32(point.Y), 5, color.White, true)
	}

	if len(g.indices) > 2 {
		for i := 0; i < len(g.indices)-2; i += 3 {
			vertices := g.indices[i : i+3]
			g.drawTriangle(screen, [3]int(vertices))

		}
	}
}

func (g *Game) drawTriangle(screen *ebiten.Image, v [3]int) {
	vx0, vy0 := g.points[v[0]].X, g.points[v[0]].Y
	vx1, vy1 := g.points[v[1]].X, g.points[v[1]].Y
	vx2, vy2 := g.points[v[2]].X, g.points[v[2]].Y

	vector.StrokeLine(screen, float32(vx0), float32(vy0), float32(vx1), float32(vy1), 1.5, color.RGBA{0, 255, 0, 1}, true)
	vector.StrokeLine(screen, float32(vx1), float32(vy1), float32(vx2), float32(vy2), 1.5, color.RGBA{0, 255, 0, 1}, true)
	vector.StrokeLine(screen, float32(vx2), float32(vy2), float32(vx0), float32(vy0), 1.5, color.RGBA{0, 255, 0, 1}, true)

}
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func main() {
	g := &Game{}
	ebiten.SetWindowSize(1000, 600)
	ebiten.SetWindowTitle("Test Delaunay Triangulation")
	if err := ebiten.RunGameWithOptions(g, nil); err != nil {
		log.Fatal(err)
	}

}
