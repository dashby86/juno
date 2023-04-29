package structs

import (
	"fmt"
	"github.com/dashby86/juno/juno"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/colornames"
)

const (
	junoWidth  = 64
	junoHeight = 64
)

// ... rest of the code

type Game struct {
	OniImage   *ebiten.Image
	Juno       *juno.Juno
	Background *Background
	Enemies    []*Enemy
	Platforms  []*Platform
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

	bg := NewBackground(bgImagePath)

	borderTop := NewPlatform(0, 0, float64(screenWidth), 10, colornames.Red)
	borderBottom := NewPlatform(0, float64(screenHeight)-10, float64(screenWidth), 10, colornames.Green)
	borderLeft := NewPlatform(0, 0, 10, float64(screenHeight), colornames.Blue)
	borderRight := NewPlatform(float64(screenWidth)-10, 0, 10, float64(screenHeight), colornames.Yellow)

	j := juno.NewJuno(junoImg, 100, 100, 10)
	j.Speed = 10 // set Juno's speed to 10 pixels per frame

	game := &Game{
		OniImage:   oniImg,
		Juno:       j,
		Background: bg,
		Enemies:    make([]*Enemy, 0), // initialize the slice of enemies
	}
	game.Platforms = []*Platform{borderRight, borderBottom, borderLeft, borderTop}

	return game, nil
}

func (g *Game) Update() error {
	g.Juno.Grounded = false
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.Juno.MoveLeft()
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.Juno.MoveRight()
	}

	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return fmt.Errorf("game is closed")
	}

	// Handle jump input
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.Juno.Jump()
	}

	// Apply gravity
	if !g.Juno.Grounded {
		g.Juno.VelY += g.Juno.Gravity
	} else {
		g.Juno.VelY = 0
	}

	//fmt.Printf("Juno: {X: %f, Y: %f, VelX: %f, VelY: %f}\n", g.Juno.X, g.Juno.Y, g.Juno.VelX, g.Juno.VelY)

	g.checkCollision()
	x, y := g.Juno.GetPosition()

	playerPos := Vec2{X: x, Y: y}

	g.Juno.X += g.Juno.VelX
	g.Juno.Y += g.Juno.VelY

	// update all enemies
	for _, enemy := range g.Enemies {
		dir := playerPos.Sub(enemy.Position)
		enemy.Update(dir)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	g.Background.Draw(screen)

	// Draw the platforms
	for _, p := range g.Platforms {
		p.Draw(screen)
	}

	// Draw the enemies
	for _, e := range g.Enemies {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(e.Position.X, e.Position.Y)
		screen.DrawImage(e.Image, op)
	}
	/**
	x, y := g.Juno.GetPosition()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	screen.DrawImage(g.Juno.Image.(*ebiten.Image), op)

	*/
	g.Juno.Draw(screen)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	screenWidth, screenHeight = ebiten.ScreenSizeInFullscreen()
	return screenWidth, screenHeight
}

func (game *Game) checkCollision() {
	ju := game.Juno
	grounded := false

	// Check horizontal collision
	for _, platform := range game.Platforms {
		nextX := ju.X + ju.VelX

		if nextX+float64(junoWidth) >= platform.X &&
			nextX <= platform.X+platform.Width &&
			ju.Y+float64(junoHeight) > platform.Y &&
			ju.Y < platform.Y+platform.Height {

			// Check if Juno is hitting the left side of the platform
			if ju.VelX > 0 {
				ju.X = platform.X - float64(junoWidth)
			} else if ju.VelX < 0 { // Check if Juno is hitting the right side of the platform
				ju.X = platform.X + platform.Width
			}
			ju.VelX = 0
		}
	}

	// Check vertical collision
	for _, platform := range game.Platforms {
		nextY := ju.Y + ju.VelY

		if ju.X+float64(junoWidth) > platform.X &&
			ju.X < platform.X+platform.Width &&
			nextY+float64(junoHeight) >= platform.Y &&
			nextY <= platform.Y+platform.Height {

			// Check if Juno is hitting the top of the platform
			if ju.VelY > 0 {
				ju.Y = platform.Y - float64(junoHeight)
				grounded = true
			} else if ju.VelY < 0 { // Check if Juno is hitting the bottom of the platform
				ju.Y = platform.Y + platform.Height
			}
			ju.VelY = 0
		}
	}

	ju.Grounded = grounded
}
