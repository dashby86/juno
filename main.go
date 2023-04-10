package main

import (
	"fmt"
	"github.com/dashby86/juno/structs"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

var g *structs.Game

func update(screen *ebiten.Image) error {
	if err := g.Update(); err != nil {
		return err
	}

	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return fmt.Errorf("game is closed")
	}

	return nil
}

func draw(screen *ebiten.Image) {
	g.Draw(screen)
}

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
		x:       screenWidth / 2,
		y:       screenHeight / 2,
		speed:   4.0,
		zoom:    1.0,
		minZoom: 0.5,
		maxZoom: 2.0,
	}

	g = &structs.Game{
		oniImage:  oniImage,
		junoImage: junoImage,
		oniPos:    ebiten.GeoM{},
		junoPos:   junoPos,
		camera:    cam, // Pass camera to game struct
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Juno")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
