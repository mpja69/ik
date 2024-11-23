package main

import (
	"github.com/mpja69/ik/ik"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	WIDTH  = 600
	HEIGHT = 600
)

type Game struct {
	limb *ik.Limb
}

// Use the mouse pointer as target
func (g *Game) Update() error {
	x, y := ebiten.CursorPosition()
	tx := float64(x)
	ty := float64(y)

	g.limb.Update(ik.Point{X: tx, Y: ty})

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Clear()
	g.limb.Draw(screen)
}

func (g *Game) Layout(_, _ int) (int, int) {
	return WIDTH, HEIGHT
}

func main() {
	ebiten.SetWindowSize(WIDTH, HEIGHT)
	ebiten.SetWindowTitle("Inverse kinematics!")
	g := Game{limb: ik.LimbNew(WIDTH/2, HEIGHT/2, 5)}

	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}
