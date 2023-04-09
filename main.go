package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

type Game struct {
	junoImage *ebiten.Image
	oniImage  *ebiten.Image
	junoPos   ebiten.GeoM
	oniPos    ebiten.GeoM
	speed     float64
}

func (g *Game) Update(screen *ebiten.Image) error {
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.junoPos.Translate(-g.speed, 0)
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.junoPos.Translate(g.speed, 0)
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.junoPos.Translate(0, -g.speed)
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.junoPos.Translate(0, g.speed)
	}

	if ebiten.IsKeyPressed(ebiten.KeyJ) {
		g.oniPos.Translate(-g.speed, 0)
	}
	if ebiten.IsKeyPressed(ebiten.KeyL) {
		g.oniPos.Translate(g.speed, 0)
	}
	if ebiten.IsKeyPressed(ebiten.KeyI) {
		g.oniPos.Translate(0, -g.speed)
	}
	if ebiten.IsKeyPressed(ebiten.KeyK) {
		g.oniPos.Translate(0, g.speed)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	var junoImageOpts ebiten.DrawImageOptions
	junoImageOpts.GeoM = g.junoPos
	screen.DrawImage(g.junoImage, &junoImageOpts)

	var oniImageOpts ebiten.DrawImageOptions
	oniImageOpts.GeoM = g.oniPos
	screen.DrawImage(g.oniImage, &oniImageOpts)

	if g.junoPos.X < g.oniPos.X {
		var junoImageOpts ebiten.DrawImageOptions
		junoImageOpts.GeoM.Concat(&g.junoPos)
		junoImageOpts.GeoM.Scale(-1, 1)
		junoImageOpts.GeoM.Translate(64, 0)
		screen.DrawImage(g.junoImage, &junoImageOpts)
	} else {
		var junoImageOpts ebiten.DrawImageOptions
		junoImageOpts.GeoM.Concat(&g.junoPos)
		screen.DrawImage(g.junoImage, &junoImageOpts)
	}

	if g.oniPos.X < g.junoPos.X {
		var oniImageOpts ebiten.DrawImageOptions
		oniImageOpts.GeoM.Concat(&g.oniPos)
		oniImageOpts.GeoM.Scale(-1, 1)
		oniImageOpts.GeoM.Translate(64, 0)
		screen.DrawImage(g.oniImage, &oniImageOpts)
	} else {
		var oniImageOpts ebiten.DrawImageOptions
		oniImageOpts.GeoM.Concat(&g.oniPos)
		screen.DrawImage(g.oniImage, &oniImageOpts)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	// create game object
	g := &Game{}

	// initialize ebiten
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Juno")

	// load images
	junoImage, _, err := ebitenutil.NewImageFromFile("assets/juno.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	oniImage, _, err := ebitenutil.NewImageFromFile("assets/oni.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	// set image options
	var junoImageOpts ebiten.DrawImageOptions
	var oniImageOpts ebiten.DrawImageOptions

	// set positions
	g.junoPos = ebiten.GeoM{}
	g.oniPos = ebiten.GeoM{}
	g.oniPos.Translate(300, 0)

	// start game
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
