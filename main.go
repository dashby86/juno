package main

import (
	"log"

	"github.com/dashby86/juno/structs"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 1920
	screenHeight = 1080
)

func main() {
	/**
	oniImage, _, err := ebitenutil.NewImageFromFile("assets/oni.png")
	if err != nil {
		log.Fatal(err)
	}
	backgroundImage, _, err := ebitenutil.NewImageFromFile("assets/background.png")
	if err != nil {
		log.Fatal(err)
	}

	*/

	junoImage, _, err := ebitenutil.NewImageFromFile("assets/juno.png")
	if err != nil {
		log.Fatal(err)
	}
	// Create and configure background

	junoPos := ebiten.GeoM{}
	junoPos.Translate(screenWidth/2, float64(screenHeight-junoImage.Bounds().Max.Y))

	// Create and configure camera
	/**
	cam := &structs.Camera{
		X:       screenWidth / 2,
		Y:       screenHeight / 2,
		Speed:   4.0,
		Zoom:    1.0,
		MinZoom: 0.5,
		MaxZoom: 2.0,
	}

	*/

	g, err := structs.NewGame("assets/background.png", "assets/juno.png", "assets/background.png", screenWidth, screenHeight)
	if err != nil {
		log.Fatal(err)
	}

	enemies := make([]*structs.Enemy, 0)
	enemy1, err := structs.NewEnemy("assets/oni.png", structs.Vec2{X: 100, Y: 100}, 2)
	enemy2, err := structs.NewEnemy("assets/oni.png", structs.Vec2{X: 100, Y: 100}, 2)
	if err != nil {
		log.Fatal(err)
	}
	enemies = append(enemies, enemy1)
	enemies = append(enemies, enemy2)
	//enemy2, err := structs.NewEnemy("assets/enemy.png", structs.Vec2{X: 200, Y: 200})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//enemies = append(enemies, enemy2)

	g.Enemies = enemies

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("Juno")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
