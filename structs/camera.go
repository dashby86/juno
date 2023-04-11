package structs

import "github.com/hajimehoshi/ebiten/v2"

type Camera struct {
	X       float64
	Y       float64
	Speed   float64
	Zoom    float64
	MinZoom float64
	MaxZoom float64
}

func (c *Camera) Move(dx, dy float64) {
	c.X += dx * c.Speed
	c.Y += dy * c.Speed
}

func (c *Camera) ZoomIn() {
	c.Zoom *= 1.1
	if c.Zoom > c.MaxZoom {
		c.Zoom = c.MaxZoom
	}
}

func (c *Camera) ZoomOut() {
	c.Zoom /= 1.1
	if c.Zoom < c.MinZoom {
		c.Zoom = c.MinZoom
	}
}

func (c *Camera) Apply(screen *ebiten.Image) {
	w, h := screen.Size()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-c.X, -c.Y)
	op.GeoM.Translate(float64(w)/2, float64(h)/2)
	op.GeoM.Scale(c.Zoom, c.Zoom)
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	screen.DrawImage(screen, op)
}
