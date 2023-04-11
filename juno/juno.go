package juno

import (
	"fmt"
	"image"
	"math/rand"
	"time"
)

type Juno struct {
	Image image.Image
	X     int
	Y     int
}

func NewJuno(img image.Image, x, y int) *Juno {
	return &Juno{
		Image: img,
		X:     x,
		Y:     y,
	}
}

func (j *Juno) Move(dx, dy int) {
	j.X += dx
	j.Y += dy
}

func (j *Juno) Chase(enemyX, enemyY int) {
	if enemyX < j.X {
		j.Move(-1, 0)
	} else if enemyX > j.X {
		j.Move(1, 0)
	}

	if enemyY < j.Y {
		j.Move(0, -1)
	} else if enemyY > j.Y {
		j.Move(0, 1)
	}
}

func (j *Juno) RandomMove() {
	rand.Seed(time.Now().UnixNano())
	dx := rand.Intn(3) - 1
	dy := rand.Intn(3) - 1
	j.Move(dx, dy)
}

func (j *Juno) String() string {
	return fmt.Sprintf("Juno at (%d,%d)", j.X, j.Y)
}
