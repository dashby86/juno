package structs

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"math"
)

type Enemy struct {
	Image    *ebiten.Image
	Position Vec2
	Speed    float64
}

func NewEnemy(imagePath string, position Vec2, speed float64) (*Enemy, error) {
	img, _, err := ebitenutil.NewImageFromFile(imagePath)
	if err != nil {
		return nil, err
	}

	return &Enemy{
		Image:    img,
		Position: position,
		Speed:    speed,
	}, nil
}

func (e *Enemy) Update(playerPos Vec2) {
	// Move towards the player
	direction := playerPos.Sub(e.Position)
	direction.Normalize()
	direction.Mul(e.Speed)
	e.Position.Add(direction)
}

func (e *Enemy) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(e.Position.X, e.Position.Y)
	screen.DrawImage(e.Image, op)
}

type Vec2 struct {
	X, Y float64
}

func (v *Vec2) Add(other Vec2) {
	v.X += other.X
	v.Y += other.Y
}

func (v *Vec2) Sub(other Vec2) Vec2 {
	return Vec2{
		X: v.X - other.X,
		Y: v.Y - other.Y,
	}
}

func (v *Vec2) Mul(scalar float64) {
	v.X *= scalar
	v.Y *= scalar
}

func (v *Vec2) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vec2) Normalize() {
	magnitude := v.Magnitude()
	if magnitude != 0 {
		v.X /= magnitude
		v.Y /= magnitude
	}
}
