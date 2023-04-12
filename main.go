package main

import (
	"log"

	"github.com/dashby86/juno/structs"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

/**
const (
	screenWidth  = 1920
	screenHeight = 1080
)

*/

func main() {
	junoImage, _, err := ebitenutil.NewImageFromFile("assets/juno.png")
	if err != nil {
		log.Fatal(err)
	}
	// Create and configure background

	screenWidth, screenHeight := ebiten.ScreenSizeInFullscreen()
	junoPos := ebiten.GeoM{}
	junoPos.Translate(float64(screenWidth/2), float64(screenHeight-junoImage.Bounds().Max.Y))

	g, err := structs.NewGame("assets/background.png", "assets/juno.png", "assets/level2.png", screenWidth, screenHeight)
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

	//ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("Juno")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
