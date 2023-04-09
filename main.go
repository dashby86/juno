package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "image/png"
	"log"
)

type Game struct {
	bgImage   *ebiten.Image
	junoImage *ebiten.Image
	oniImage  *ebiten.Image
	junoPos   ebiten.GeoM
	oniPos    ebiten.GeoM
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.junoPos.Translate(-5, 0)
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.junoPos.Translate(5, 0)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.bgPos.X, g.bgPos.Y)
	screen.DrawImage(g.bgImage, op)

	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.junoPos.X, g.junoPos.Y)
	if g.isJunoFacingRight {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(-float64(g.junoImage.Bounds().Size().X), 0)
	}
	screen.DrawImage(g.junoImage, op)

	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.oniPos.X, g.oniPos.Y)
	if g.isOniFacingRight {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(-float64(g.oniImage.Bounds().Size().X), 0)
	}
	screen.DrawImage(g.oniImage, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	// Load the background image
	bgImage, _, err := ebitenutil.NewImageFromFile("assets/background.png")
	if err != nil {
		log.Fatal(err)
	}

	// Load the Juno image
	junoImage, _, err := ebitenutil.NewImageFromFile("assets/juno.png")
	if err != nil {
		log.Fatal(err)
	}

	// Load the Oni image
	oniImage, _, err := ebitenutil.NewImageFromFile("assets/oni.png")
	if err != nil {
		log.Fatal(err)
	}

	// Create the game object
	game := &Game{
		bgImage:   bgImage,
		junoImage: junoImage,
		oniImage:  oniImage,
		junoPos:   ebiten.GeoM{},
		oniPos:    ebiten.GeoM{},
	}

	// Set the initial positions of the characters
	game.junoPos.Translate(0, 200)
	game.oniPos.Translate(400, 200)

	// Run the game
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Juno")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
