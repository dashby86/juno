package structs

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

type Platform struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
	Color  color.Color
}

func NewPlatform(x, y, width, height float64, color color.Color) *Platform {
	return &Platform{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
		Color:  color,
	}
}

func (p *Platform) Draw(screen *ebiten.Image, cameraX, cameraY float64) {
	platformImage := ebiten.NewImage(int(p.Width), int(p.Height))
	platformImage.Fill(p.Color)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.X-cameraX+float64(screen.Bounds().Dx())/2, p.Y-cameraY+float64(screen.Bounds().Dy())/2)
	screen.DrawImage(platformImage, op)
}
