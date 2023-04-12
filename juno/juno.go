package juno

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"math/rand"
	"time"
)

type Juno struct {
	Image      image.Image
	X          float64
	Y          float64
	Speed      float64
	VelY       float64
	Grounded   bool
	JumpHeight float64
	Gravity    float64
}

func NewJuno(img image.Image, x, y, speed float64) *Juno {
	return &Juno{
		Image:      img,
		X:          x,
		Y:          y,
		Speed:      speed,
		Grounded:   false,
		JumpHeight: 15,
		Gravity:    0.5,
	}

}

func (j *Juno) Jump() {
	if j.Grounded {
		j.VelY = +j.JumpHeight
		j.Grounded = false
	}
}

func (j *Juno) Move(dx, dy float64) {
	j.X += dx
	j.Y += dy
}

func (j *Juno) Chase(enemyX, enemyY float64) {
	if enemyX < j.X {
		j.Move(-j.Speed, 0)
	} else if enemyX > j.X {
		j.Move(j.Speed, 0)
	}

	if enemyY < j.Y {
		j.Move(0, -j.Speed)
	} else if enemyY > j.Y {
		j.Move(0, j.Speed)
	}
}

func (j *Juno) RandomMove() {
	rand.Seed(time.Now().UnixNano())
	dx := rand.Float64()*j.Speed*2 - j.Speed
	dy := rand.Float64()*j.Speed*2 - j.Speed
	j.Move(dx, dy)
}

func (j *Juno) String() string {
	return fmt.Sprintf("Juno at (%f,%f)", j.X, j.Y)
}

// MoveLeft moves Juno to the left
func (j *Juno) MoveLeft() {
	j.X += j.Speed
}

// MoveRight moves Juno to the right
func (j *Juno) MoveRight() {
	j.X -= j.Speed
}

// MoveUp moves Juno up
func (j *Juno) MoveUp() {
	j.Y += j.Speed
}

// MoveDown moves Juno down
func (j *Juno) MoveDown() {
	j.Y -= j.Speed
}

func (j *Juno) GetPosition() (float64, float64) {
	return j.X, j.Y
}

func (j *Juno) GetWidth() int {
	if rgba, ok := j.Image.(*image.RGBA); ok {
		return rgba.Rect.Size().X
	} else {
		return j.Image.Bounds().Size().X
	}
}

func (j *Juno) GetHeight() int {
	if rgba, ok := j.Image.(*image.RGBA); ok {
		return rgba.Rect.Size().Y
	} else {
		return j.Image.Bounds().Size().Y
	}
}

func (j *Juno) Draw(screen *ebiten.Image, cameraX, cameraY float64) {
	op := &ebiten.DrawImageOptions{}
	ebitenImage := j.Image.(*ebiten.Image)
	x, y := j.GetPosition()
	op.GeoM.Translate(x-cameraX+float64(screen.Bounds().Max.X)/2, y-cameraY+float64(screen.Bounds().Max.Y)/2)
	screen.DrawImage(ebitenImage, op)
}
