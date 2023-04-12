package structs

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

type Background struct {
	Image *ebiten.Image
}

func NewBackground(imagePath string) *Background {
	img, _, err := ebitenutil.NewImageFromFile(imagePath)
	if err != nil {
		log.Fatal(err)
	}

	return &Background{
		Image: img,
	}
}

func (b *Background) Draw(screen *ebiten.Image) {
	screenWidth, screenHeight := screen.Size()
	bgWidth, bgHeight := b.Image.Size()
	scaleX, scaleY := float64(screenWidth)/float64(bgWidth), float64(screenHeight)/float64(bgHeight)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scaleX, scaleY)
	screen.DrawImage(b.Image, op)
}
