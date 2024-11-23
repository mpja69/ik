package main

import (
	"ik"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	WIDTH  = 600
	HEIGHT = 600
)

type Game struct {
	limb *limb
}

// Use the mouse pointer as target
func (g *Game) Update() error {
	x, y := ebiten.CursorPosition()
	tx := float64(x)
	ty := float64(y)

	g.limb.update(point{tx, ty})

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Clear()
	g.limb.draw(screen)
}

func (g *Game) Layout(_, _ int) (int, int) {
	return WIDTH, HEIGHT
}

func main() {
	ebiten.SetWindowSize(WIDTH, HEIGHT)
	ebiten.SetWindowTitle("Inverse kinematics!")
	g := Game{limb: limbNew(WIDTH/2, HEIGHT/2, 5)}

	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}
