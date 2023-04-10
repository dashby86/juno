package structs

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

type Game struct {
	oniImage  *ebiten.Image
	junoImage *ebiten.Image
	oniPos    ebiten.GeoM
	junoPos   ebiten.GeoM
	camera    *Camera // Pointer to camera struct
}

func (g *Game) Update() error {
	// Exit the game if the Escape key is pressed
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return fmt.Errorf("game is closed")
	}

	// Update Juno position
	junoSpeed := 4.0
	var junoPos ebiten.GeoM
	junoPos.Concat(g.junoPos)
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		junoPos.Translate(-junoSpeed, 0)
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		junoPos.Translate(junoSpeed, 0)
	}
	g.junoPos = junoPos

	// Update Oni position
	oniSpeed := 2.0
	var oniPos ebiten.GeoM
	oniPos.Concat(g.oniPos)

	// Move towards Juno
	dx := g.junoPos.Element(0, 2) - g.oniPos.Element(0, 2)
	dy := g.junoPos.Element(1, 2) - g.oniPos.Element(1, 2)
	angle := math.Atan2(dy, dx)
	oniPos.Translate(math.Cos(angle)*oniSpeed, math.Sin(angle)*oniSpeed)

	g.oniPos = oniPos

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Apply camera transform matrix
	screen.Set(ebiten.ComposeMatrix(
		ebiten.TranslateCoordinates(-g.camera.x, -g.camera.y),
		ebiten.Scale(g.camera.zoom, g.camera.zoom),
	))

	// Draw game objects
	oniImageOpts := &ebiten.DrawImageOptions{}
	oniImageOpts.GeoM = g.oniPos
	screen.DrawImage(g.oniImage, oniImageOpts)

	junoImageOpts := &ebiten.DrawImageOptions{}
	junoImageOpts.GeoM = g.junoPos
	screen.DrawImage(g.junoImage, junoImageOpts)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return main.screenWidth, main.screenHeight
}
