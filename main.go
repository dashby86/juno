package main

import (
	"github.com/hajimehoshi/ebiten"
	"log"
)

type Game struct {
	junoImage  *ebiten.Image
	oniImage   *ebiten.Image
	bgImage    *ebiten.Image
	junoPos    ebiten.GeoM
	oniPos     ebiten.GeoM
	screenSize struct {
		W, H int
	}
}

func NewGame() (*Game, error) {
	junoImage, _, err := ebitenutil.NewImageFromFile("assets/juno.png", ebiten.FilterDefault)
	if err != nil {
		return nil, err
	}
	oniImage, _, err := ebitenutil.NewImageFromFile("assets/oni.png", ebiten.FilterDefault)
	if err != nil {
		return nil, err
	}
	bgImage, _, err := ebitenutil.NewImageFromFile("assets/background.png", ebiten.FilterDefault)
	if err != nil {
		return nil, err
	}

	game := &Game{
		junoImage: junoImage,
		oniImage:  oniImage,
		bgImage:   bgImage,
		screenSize: struct {
			W, H int
		}{
			W: 640,
			H: 480,
		},
	}

	game.junoPos.Scale(2, 2)
	game.junoPos.Translate(0, float64(game.screenSize.H)/2)

	game.oniPos.Scale(2, 2)
	game.oniPos.Translate(float64(game.screenSize.W)-float64(game.oniImage.Bounds().Max.X)*2, float64(game.screenSize.H)/2)

	return game, nil
}

func (g *Game) Update(screen *ebiten.Image) error {
	return nil
}

func LeftFacing(gm *ebiten.GeoM, img *ebiten.Image) ebiten.DrawImageOptions {
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Scale(-1, 1)
	opts.GeoM.Concat(gm)
	opts.GeoM.Translate(-float64(img.Bounds().Size().X), 0)
	return opts
}

func RightFacing(gm *ebiten.GeoM, img *ebiten.Image) ebiten.DrawImageOptions {
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Concat(gm)
	return opts
}

func (g *Game) draw(screen *ebiten.Image) error {
	// Draw the background image.
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(g.bgImage, op)

	// Draw Juno and Oni, facing left or right depending on their respective velocities.
	if g.junoVelX < 0 {
		op := LeftFacing(&g.junoPos, g.junoImage)
		screen.DrawImage(g.junoImage, op)
	} else {
		op := RightFacing(&g.junoPos, g.junoImage)
		screen.DrawImage(g.junoImage, op)
	}

	if g.oniVelX < 0 {
		op := LeftFacing(&g.oniPos, g.oniImage)
		screen.DrawImage(g.oniImage, op)
	} else {
		op := RightFacing(&g.oniPos, g.oniImage)
		screen.DrawImage(g.oniImage, op)
	}

	return nil
}

func main() {
	game := NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
