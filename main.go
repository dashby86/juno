package main

import (
	"github.com/dashby86/juno/structs"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

const (
	screenWidth  = 1280
	screenHeight = 960
)

func main() {
	junoImage, _, err := ebitenutil.NewImageFromFile("assets/juno.png")
	if err != nil {
		log.Fatal(err)
	}

	junoPos := ebiten.GeoM{}
	junoPos.Translate(screenWidth/2, float64(screenHeight-junoImage.Bounds().Max.Y))

	g, err := structs.NewGame("assets/oni.png", "assets/juno.png", "assets/background.png", screenWidth, screenHeight)
	if err != nil {
		log.Fatal(err)
	}

	enemies := make([]*structs.Enemy, 0)
	enemy1, err := structs.NewEnemy("assets/oni.png", structs.Vec2{X: 100, Y: 100}, 20)
	if err != nil {
		log.Fatal(err)
	}
	enemies = append(enemies, enemy1)
	//enemy2, err := structs.NewEnemy("assets/enemy.png", structs.Vec2{X: 200, Y: 200})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//enemies = append(enemies, enemy2)

	g.Enemies = enemies

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Juno")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
