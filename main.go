package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2/text"
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

type Game struct {
	bgImage   *ebiten.Image
	junoImage *ebiten.Image
	oniImage  *ebiten.Image
	junoPos   ebiten.GeoM
	oniPos    ebiten.GeoM
}

func NewGame() (*Game, error) {
	bg, err := ebiten.NewImageFromFile("./assets/background.png")
	if err != nil {
		return nil, err
	}

	juno, err := ebiten.NewImageFromFile("./assets/juno.png")
	if err != nil {
		return nil, err
	}

	oni, err := ebiten.NewImageFromFile("./assets/oni.png")
	if err != nil {
		return nil, err
	}

	g := &Game{
		bgImage:   bg,
		junoImage: juno,
		oniImage:  oni,
		junoPos:   ebiten.GeoM{},
		oniPos:    ebiten.GeoM{},
	}
	g.junoPos.Translate(50, 50)
	g.oniPos.Translate(300, 50)

	return g, nil
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.junoPos.Translate(-5, 0)
		g.junoImageOpts = LeftFacing(g)
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.junoPos.Translate(5, 0)
		g.junoImageOpts = RightFacing(g)
	}

	return nil
}

func (g *Game) draw(screen *ebiten.Image) {
	// Draw the background image
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(2, 2)
	screen.DrawImage(g.bgImage, op)

	// Draw Juno
	junoOp := LeftFacing(&g.junoPos, g.junoImage)
	junoOp.GeoM.Scale(2, 2)
	screen.DrawImage(g.junoImage, &junoOp)

	// Draw the oni
	oniOp := RightFacing(&g.oniPos, g.oniImage)
	oniOp.GeoM.Scale(2, 2)
	screen.DrawImage(g.oniImage, &oniOp)

	// Draw the score
	scoreText := fmt.Sprintf("Score: %d", g.score)
	text.Draw(screen, scoreText, mplusBigFont, screenWidth/2, 60, color.Black)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func LeftFacing(g *Game) ebiten.DrawImageOptions {
	return ebiten.DrawImageOptions{
		Image: g.junoImage,
		GeoM:  g.junoPos,
	}
}

func RightFacing(g *Game) ebiten.DrawImageOptions {
	return ebiten.DrawImageOptions{
		Image: g.junoImage,
		GeoM:  g.junoPos,
		FX:    ebiten.FlipHorizontal,
	}
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Juno")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
