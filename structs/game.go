package structs

import (
	"fmt"
	"github.com/dashby86/juno/juno"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	OniImage   *ebiten.Image
	Juno       *Juno
	Background *ebiten.Image
	Camera     *Camera
	Enemies    []*Enemy // slice to store all enemies
}

func NewGame(oniImagePath, junoImagePath, bgImagePath string, screenWidth, screenHeight int) (*Game, error) {

	oniImg, _, err := ebitenutil.NewImageFromFile(oniImagePath)
	if err != nil {
		return nil, err
	}

	juno, err := NewJuno(junoImagePath, screenWidth, screenHeight)
	if err != nil {
		return nil, err
	}

	bgImg, _, err := ebitenutil.NewImageFromFile(bgImagePath)
	if err != nil {
		return nil, err
	}

	game := &Game{
		OniImage:   oniImg,
		Juno:       juno,
		Background: bgImg,
		Camera:     &Camera{},
		Enemies:    make([]*Enemy, 0), // initialize the slice of enemies
	}

	game.Camera.PosX = float64(screenWidth) / 2
	game.Camera.PosY = float64(screenHeight) / 2

	return game, nil
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.Juno.MoveLeft()
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.Juno.MoveRight()
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.Juno.MoveUp()
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.Juno.MoveDown()
	}

	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return fmt.Errorf("game is closed")
	}

	g.JunoPos = junoPos

	g.Camera.PosX = g.Juno.Position.X + float64(g.Juno.Width)/2
	g.Camera.PosY = g.Juno.Position.Y + float64(g.Juno.Height)/2

	playerPos := Vec2{X: g.Juno.Position.X, Y: g.Juno.Position.Y}

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
	op.GeoM.Translate(-g.Camera.PosX+float64(screen.Bounds().Max.X)/2, -g.Camera.PosY+float64(screen.Bounds().Max.Y)/2)
	screen.DrawImage(g.Background, op)

	// Draw the enemies
	for _, e := range g.Enemies {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(e.Position.X-g.Camera.PosX+float64(screen.Bounds().Max.X)/2, e.Position.Y-g.Camera.PosY+float64(screen.Bounds().Max.Y)/2)
		screen.DrawImage(e.Image, op)
	}

	// Draw Juno
	g.Juno.Draw(screen, g.Camera.PosX, g.Camera.PosY)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// Set the size of the screen based on the outsideWidth and outsideHeight parameters.
	// For example, you could set the screen size to be half the size of the outside area:
	screenWidth, screenHeight = outsideWidth/2, outsideHeight/2
	return screenWidth, screenHeight
}
