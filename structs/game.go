package structs

import (
	"fmt"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

type Game struct {
	OniImage  *ebiten.Image
	JunoImage *ebiten.Image
	OniPos    ebiten.GeoM
	JunoPos   ebiten.GeoM
	Camera    *Camera // Pointer to camera struct
}

func (g *Game) Update() error {
	// Exit the game if the Escape key is pressed
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return fmt.Errorf("game is closed")
	}

	// Update Juno position
	junoSpeed := 4.0
	var junoPos ebiten.GeoM
	junoPos.Concat(g.JunoPos)
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		junoPos.Translate(-junoSpeed, 0)
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		junoPos.Translate(junoSpeed, 0)
	}
	g.JunoPos = junoPos

	// Update Oni position
	oniSpeed := 2.0
	var oniPos ebiten.GeoM
	oniPos.Concat(g.OniPos)

	// Move towards Juno
	dx := g.JunoPos.Element(0, 2) - g.OniPos.Element(0, 2)
	dy := g.JunoPos.Element(1, 2) - g.OniPos.Element(1, 2)
	angle := math.Atan2(dy, dx)
	oniPos.Translate(math.Cos(angle)*oniSpeed, math.Sin(angle)*oniSpeed)

	g.OniPos = oniPos

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Apply camera transform matrix
	geoM := ebiten.GeoM{}
	geoM.Translate(-g.Camera.X, -g.Camera.Y)
	geoM.Scale(g.Camera.Zoom, g.Camera.Zoom)
	screen.DrawImage(g.OniImage, &ebiten.DrawImageOptions{
		GeoM: geoM,
	})

	// Draw game objects
	oniImageOpts := &ebiten.DrawImageOptions{}
	oniImageOpts.GeoM = g.OniPos
	screen.DrawImage(g.OniImage, oniImageOpts)

	junoImageOpts := &ebiten.DrawImageOptions{}
	junoImageOpts.GeoM = g.JunoPos
	screen.DrawImage(g.JunoImage, junoImageOpts)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
