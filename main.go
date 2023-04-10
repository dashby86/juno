package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
	"math"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

var g *game

type game struct {
	oniImage  *ebiten.Image
	junoImage *ebiten.Image
	oniPos    ebiten.GeoM
	junoPos   ebiten.GeoM
}

func (g *game) Update() error {
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

func (g *game) Draw(screen *ebiten.Image) {
	// Draw background
	bgImage, _, err := ebitenutil.NewImageFromFile("assets/background.png")
	if err != nil {
		log.Fatal(err)
	}
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(bgImage, op)

	// Draw Oni and Juno
	screen.DrawImage(g.oniImage, &ebiten.DrawImageOptions{
		GeoM: g.oniPos,
	})
	screen.DrawImage(g.junoImage, &ebiten.DrawImageOptions{
		GeoM: g.junoPos,
	})

	// Draw FPS counter
	fps := ebiten.CurrentFPS()
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.2f", fps))
}

func (g *game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
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
	//junoPos.Translate(screenWidth/4, screenHeight/2)

	junoPos.Translate(screenWidth/2, float64(screenHeight-junoImage.Bounds().Max.Y))

	g := &game{
		oniImage:  oniImage,
		junoImage: junoImage,
		oniPos:    ebiten.GeoM{},
		junoPos:   junoPos,
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetFullscreen(false)
	err = ebiten.RunGame(g)
	if err != nil {
		return
	}
}
