package structs

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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

func (p *Platform) Draw(screen *ebiten.Image, camX, camY float64) {
	x := p.X - camX + screenWidth/2
	y := p.Y - camY + screenHeight/2
	ebitenutil.DrawRect(screen, x, y, p.Width, p.Height, p.Color)
}
