package structs

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	OniImage   *ebiten.Image
	JunoImage  *ebiten.Image
	Background *ebiten.Image
	OniPos     ebiten.GeoM
	JunoPos    ebiten.GeoM
	Camera     *Camera
}

const (
	oniX  = 200
	oniY  = 200
	junoX = 100
	junoY = 100
)

func NewGame() (*Game, error) {
	oniImage, _, err := ebitenutil.NewImageFromFile("oni.png")
	if err != nil {
		return nil, err
	}
	junoImage, _, err := ebitenutil.NewImageFromFile("juno.png")
	if err != nil {
		return nil, err
	}
	bgImage, _, err := ebitenutil.NewImageFromFile("background.png")
	if err != nil {
		return nil, err
	}
	oniPos := ebiten.GeoM{}
	oniPos.Translate(oniX, oniY)
	junoPos := ebiten.GeoM{}
	junoPos.Translate(junoX, junoY)
	game := &Game{
		OniImage:   oniImage,
		JunoImage:  junoImage,
		Background: bgImage,
		OniPos:     oniPos,
		JunoPos:    junoPos,
		Camera:     &Camera{X: 0, Y: 0, Zoom: 1},
	}

	return game, nil
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.JunoPos.Translate(-10, 0)
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.JunoPos.Translate(10, 0)
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.JunoPos.Translate(0, -10)
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.JunoPos.Translate(0, 10)
	}
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		g.Camera.ZoomIn()
	}
	if ebiten.IsKeyPressed(ebiten.KeyE) {
		g.Camera.ZoomOut()
	}
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return fmt.Errorf("game is closed")
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	w, h := g.Background.Size()

	// Draw the background
	bgOp := &ebiten.DrawImageOptions{}
	bgOp.GeoM.Concat(g.Camera.GetMatrix())
	bgOp.GeoM.Translate(float64(w)/2, float64(h)/2)
	bgOp.GeoM.Translate(-g.Camera.Zoom*float64(w)/2, -g.Camera.Zoom*float64(h)/2)
	screen.DrawImage(g.Background, bgOp)

	// Draw Juno
	junoOp := &ebiten.DrawImageOptions{}
	junoOp.GeoM.Concat(g.Camera.GetMatrix())
	junoOp.GeoM.Concat(g.JunoPos)
	junoOp.GeoM.Translate(-float64(g.JunoImage.Bounds().Dx())/2, -float64(g.JunoImage.Bounds().Dy())/2)
	screen.DrawImage(g.JunoImage, junoOp)

	// Draw Oni
	oniOp := &ebiten.DrawImageOptions{}
	oniOp.GeoM.Concat(g.Camera.GetMatrix())
	oniOp.GeoM.Concat(g.OniPos)
	oniOp.GeoM.Translate(-float64(g.OniImage.Bounds().Dx())/2, -float64(g.OniImage.Bounds().Dy())/2)
	screen.DrawImage(g.OniImage, oniOp)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// Set the size of the screen based on the outsideWidth and outsideHeight parameters.
	// For example, you could set the screen size to be half the size of the outside area:
	screenWidth, screenHeight = outsideWidth/2, outsideHeight/2
	return screenWidth, screenHeight
}
