package main

import (
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

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.bgImage, nil)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Concat(&g.junoPos)
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		op.GeoM = LeftFacing(&g.junoPos, g.junoImage)
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		op.GeoM = RightFacing(&g.junoPos, g.junoImage)
	}
	screen.DrawImage(g.junoImage, op)
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Concat(&g.oniPos)
	screen.DrawImage(g.oniImage, op)
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
