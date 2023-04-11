package main

import (
	"log"

	"github.com/dashby86/juno/structs"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

func main() {
	oniImage, _, err := ebitenutil.NewImageFromFile("assets/oni.png")
	if err != nil {
		log.Fatal(err)
	}

	junoImage, _, err := ebitenutil.NewImageFromFile("assets/juno.png")
	if err != nil {
		log.Fatal(err)
	}

	junoPos := ebiten.GeoM{}
	junoPos.Translate(screenWidth/2, float64(screenHeight-junoImage.Bounds().Max.Y))

	// Create and configure camera
	cam := &structs.Camera{
		X:       screenWidth / 2,
		Y:       screenHeight / 2,
		Speed:   4.0,
		Zoom:    1.0,
		MinZoom: 0.5,
		MaxZoom: 2.0,
	}

	// Create and configure background
	backgroundImage, _, err := ebitenutil.NewImageFromFile("assets/background.png")
	if err != nil {
		log.Fatal(err)
	}

	g, err := structs.NewGame(oniImage, junoImage, "assets/background.png", screenWidth, screenHeight)
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Juno")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
