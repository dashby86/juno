package structs

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	OniImage   *ebiten.Image
	JunoImage  *ebiten.Image
	Background *ebiten.Image
	OniPos     ebiten.GeoM
	JunoPos    ebiten.GeoM
	Camera     *Camera
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.Camera.Move(-1, 0)
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.Camera.Move(1, 0)
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.Camera.Move(0, -1)
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.Camera.Move(0, 1)
	}
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		g.Camera.ZoomIn()
	}
	if ebiten.IsKeyPressed(ebiten.KeyE) {
		g.Camera.ZoomOut()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	w, h := g.Background.Size()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(g.Camera.Zoom, g.Camera.Zoom)
	op.GeoM.Translate(-g.Camera.X, -g.Camera.Y)
	op.GeoM.Translate(float64(w)/2, float64(h)/2)
	op.GeoM.Translate(-g.Camera.Zoom*float64(w)/2, -g.Camera.Zoom*float64(h)/2)
	screen.DrawImage(g.Background, op)

	op = &ebiten.DrawImageOptions{}
	op.GeoM.Concat(g.JunoPos)
	screen.DrawImage(g.JunoImage, op)

	op = &ebiten.DrawImageOptions{}
	op.GeoM.Concat(g.OniPos)
	screen.DrawImage(g.OniImage, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// Set the size of the screen based on the outsideWidth and outsideHeight parameters.
	// For example, you could set the screen size to be half the size of the outside area:
	screenWidth, screenHeight = outsideWidth/2, outsideHeight/2
	return screenWidth, screenHeight
}
