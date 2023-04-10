package structs

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Camera struct {
	X, Y          float64
	Zoom          float64
	Speed         float64
	Width, Height int
}

func NewCamera(x, y float64, zoom float64, width, height int) *Camera {
	return &Camera{
		X:      x,
		Y:      y,
		Zoom:   zoom,
		Width:  width,
		Height: height,
	}
}

func (c *Camera) Move(x, y float64) {
	c.X += x
	c.Y += y
}

func (c *Camera) SetPosition(x, y float64) {
	c.X = x
	c.Y = y
}

func (c *Camera) SetZoom(zoom float64) {
	c.Zoom = zoom
}

func (c *Camera) View() ebiten.GeoM {
	geoM := ebiten.GeoM{}
	geoM.Translate(-c.X, -c.Y)
	geoM.Scale(c.Zoom, c.Zoom)
	return geoM
}

func (c *Camera) ScreenToWorld(x, y int) (float64, float64) {
	worldX := float64(x)/c.Zoom + c.X - float64(c.Width)/(2*c.Zoom)
	worldY := float64(y)/c.Zoom + c.Y - float64(c.Height)/(2*c.Zoom)
	return worldX, worldY
}

func (c *Camera) WorldToScreen(x, y float64) (int, int) {
	screenX := int((x-c.X+float64(c.Width)/(2*c.Zoom))*c.Zoom + 0.5)
	screenY := int((y-c.Y+float64(c.Height)/(2*c.Zoom))*c.Zoom + 0.5)
	return screenX, screenY
}
