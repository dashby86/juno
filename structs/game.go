package structs

import (
	"fmt"
	"github.com/dashby86/juno/juno"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/colornames"
)

const (
	junoWidth    = 64
	junoHeight   = 64
	screenWidth  = 1920
	screenHeight = 1080
)

// ... rest of the code

type Game struct {
	OniImage   *ebiten.Image
	Juno       *juno.Juno
	Background *Background
	Camera     *Camera
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

	borderTop := NewPlatform(0, 0, float64(screenWidth), 10, colornames.Black)
	borderBottom := NewPlatform(0, float64(screenHeight)-10, float64(screenWidth), 10, colornames.Black)
	borderLeft := NewPlatform(0, 0, 10, float64(screenHeight), colornames.Black)
	borderRight := NewPlatform(float64(screenWidth)-10, 0, 10, float64(screenHeight), colornames.Black)

	j := juno.NewJuno(junoImg, float64(screenWidth), float64(screenHeight), 4)
	j.Speed = 4 // set Juno's speed to 10 pixels per frame

	game := &Game{
		OniImage:   oniImg,
		Juno:       j,
		Background: bg,
		Camera:     &Camera{},
		Enemies:    make([]*Enemy, 0), // initialize the slice of enemies
	}
	game.Platforms = []*Platform{borderRight, borderBottom, borderLeft, borderTop}
	game.Camera.PosX = float64(screenWidth) / 2
	game.Camera.PosY = float64(screenHeight) / 2

	return game, nil
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		//g.Camera.PosX -= g.Juno.Speed
		g.Juno.MoveLeft()

	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		//g.Camera.PosX += g.Juno.Speed
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

	x, y := g.Juno.GetPosition()

	halfScreenWidth := float64(screenWidth) / 2
	halfScreenHeight := float64(screenHeight) / 2

	targetX := x + float64(g.Juno.GetWidth())/2 - halfScreenWidth
	targetY := y + float64(g.Juno.GetHeight())/2 - halfScreenHeight

	lerpFactor := 0.1 // You can adjust this value to change the camera's smoothness

	g.Camera.PosX = g.Camera.PosX + (targetX-g.Camera.PosX)*lerpFactor
	g.Camera.PosY = g.Camera.PosY + (targetY-g.Camera.PosY)*lerpFactor

	playerPos := Vec2{X: x, Y: y}

	g.checkCollision()

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
		p.Draw(screen, g.Camera.PosX, g.Camera.PosY)
	}

	// Draw the enemies
	for _, e := range g.Enemies {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(e.Position.X-g.Camera.PosX, e.Position.Y-g.Camera.PosY)
		screen.DrawImage(e.Image, op)
	}

	x, y := g.Juno.GetPosition()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x-g.Camera.PosX, y-g.Camera.PosY)
	screen.DrawImage(g.Juno.Image.(*ebiten.Image), op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// Set the size of the screen based on the outsideWidth and outsideHeight parameters.
	// For example, you could set the screen size to be half the size of the outside area:
	screenWidth, screenHeight = outsideWidth/2, outsideHeight/2
	return screenWidth, screenHeight
}

func (game *Game) checkCollision() {
	ju := game.Juno
	for _, platform := range game.Platforms {
		nextX := ju.X + ju.VelX
		nextY := ju.Y + ju.VelY

		// Check horizontal collision
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

		// Check vertical collision
		if ju.X+float64(junoWidth) > platform.X &&
			ju.X < platform.X+platform.Width &&
			nextY+float64(junoHeight) >= platform.Y &&
			nextY <= platform.Y+platform.Height {

			// Check if Juno is hitting the top of the platform
			if ju.VelY > 0 {
				ju.Y = platform.Y - float64(junoHeight)
				//ju.JumpCount = 0
			} else if ju.VelY < 0 { // Check if Juno is hitting the bottom of the platform
				ju.Y = platform.Y + platform.Height
			}
			ju.VelY = 0
		}
	}
}
