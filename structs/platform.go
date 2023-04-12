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

func NewPlatform(x, y, width, height float64, col color.Color) *Platform {
	//r, g, b, _ := col.RGBA()
	return &Platform{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
		Color:  color.Black,
		//Color:  color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: 0x33},
	}
}

func (p *Platform) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.X, p.Y)
	tmpImg := ebiten.NewImage(int(p.Width), int(p.Height))
	tmpImg.Fill(p.Color)
	screen.DrawImage(tmpImg, op)
}
