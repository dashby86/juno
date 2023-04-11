package structs

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	OniImage   *ebiten.Image
	JunoImage  *ebiten.Image
	Background *ebiten.Image
	OniPos     ebiten.GeoM
	JunoPos    ebiten.GeoM
	Camera     *Camera
	Enemies    []*Enemy // slice to store all enemies
}

func NewGame(oniImagePath, junoImagePath, bgImagePath string, screenWidth, screenHeight int) (*Game, error) {

	oniImg, _, err := ebitenutil.NewImageFromFile(oniImagePath)
	if err != nil {
		return nil, err
	}

	junoImg, _, err := ebitenutil.NewImageFromFile(junoImagePath)
	if err != nil {
		return nil, err
	}

	bgImg, _, err := ebitenutil.NewImageFromFile(bgImagePath)
	if err != nil {
		return nil, err
	}

	game := &Game{
		OniImage:   oniImg,
		JunoImage:  junoImg,
		Background: bgImg,
		Camera:     &Camera{},
		Enemies:    make([]*Enemy, 0), // initialize the slice of enemies
	}

	game.JunoPos = ebiten.GeoM{}
	game.JunoPos.Translate(float64(screenWidth/2), float64(screenHeight-junoImg.Bounds().Max.Y))

	game.Camera.PosX = game.JunoPos.Element(0, 0) + float64(junoImg.Bounds().Max.X)/2
	game.Camera.PosY = game.JunoPos.Element(0, 1) + float64(junoImg.Bounds().Max.Y)/2

	return game, nil
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.JunoPos.Translate(-3, 0)
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.JunoPos.Translate(3, 0)
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.JunoPos.Translate(0, -3)
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.JunoPos.Translate(0, 3)
	}

	g.Camera.PosX = g.JunoPos.Element(0, 0) + float64(g.JunoImage.Bounds().Max.X)/2
	g.Camera.PosY = g.JunoPos.Element(0, 1) + float64(g.JunoImage.Bounds().Max.Y)/2

	playerPos := Vec2{X: g.JunoPos.Element(0, 0), Y: g.JunoPos.Element(0, 1)}

	// update all enemies
	for _, enemy := range g.Enemies {
		dir := playerPos.Sub(enemy.Position)
		enemy.Update(dir)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw the background
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(g.Background, op)

	// Draw the enemies
	for _, e := range g.Enemies {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(e.Position.X-g.Camera.PosX+float64(screen.Bounds().Max.X)/2, e.Position.Y-g.Camera.PosY+float64(screen.Bounds().Max.Y)/2)
		screen.DrawImage(e.Image, op)
	}

	// Draw Juno
	op = &ebiten.DrawImageOptions{}
	op.GeoM = g.JunoPos
	op.GeoM.Translate(-g.Camera.PosX+float64(screen.Bounds().Max.X)/2, -g.Camera.PosY+float64(screen.Bounds().Max.Y)/2)
	screen.DrawImage(g.JunoImage, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// Set the size of the screen based on the outsideWidth and outsideHeight parameters.
	// For example, you could set the screen size to be half the size of the outside area:
	screenWidth, screenHeight = outsideWidth/2, outsideHeight/2
	return screenWidth, screenHeight
}
