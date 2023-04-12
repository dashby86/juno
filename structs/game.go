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
	Background *ebiten.Image
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

	j := juno.NewJuno(junoImg, float64(screenWidth), float64(screenHeight), 4)
	j.Speed = 4 // set Juno's speed to 10 pixels per frame

	bgImg, _, err := ebitenutil.NewImageFromFile(bgImagePath)
	if err != nil {
		return nil, err
	}

	game := &Game{
		OniImage:   oniImg,
		Juno:       j,
		Background: bgImg,
		Camera:     &Camera{},
		Enemies:    make([]*Enemy, 0), // initialize the slice of enemies
	}

	platform3 := NewPlatform(float64(screenWidth)/2-250, float64(screenHeight)/2-50, 500, 30, colornames.Black)
	platform1 := NewPlatform(200, 900, 500, 30, colornames.Black)
	platform2 := NewPlatform(800, 700, 500, 30, colornames.Black)
	game.Platforms = []*Platform{platform1, platform2, platform3}

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

	// Handle jump input
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.Juno.Jump()
	}

	// Apply gravity
	if !g.Juno.Grounded {
		g.Juno.VelY -= g.Juno.Gravity
	}

	// Check for collisions with platforms and apply gravity
	g.Juno.Grounded = false
	for _, platform := range g.Platforms {
		if g.Juno.Y+float64(junoHeight) >= platform.Y && g.Juno.Y+float64(junoHeight) <= platform.Y+platform.Height &&
			g.Juno.X+float64(junoWidth) >= platform.X && g.Juno.X <= platform.X+platform.Width {
			g.Juno.Grounded = true
			g.Juno.VelY = 0
			g.Juno.Y = platform.Y - float64(junoHeight)
		}
	}

	// Update Juno's vertical position
	g.Juno.Y += g.Juno.VelY

	// Prevent Juno from falling below the screen
	if g.Juno.Y > screenHeight {
		g.Juno.Y = screenHeight - float64(junoHeight)
		g.Juno.Grounded = true
		g.Juno.VelY = 0
	}

	x, y := g.Juno.GetPosition()

	g.Camera.PosX = x + float64(g.Juno.GetWidth())/2
	g.Camera.PosY = y + float64(g.Juno.GetHeight())/2

	playerPos := Vec2{X: x, Y: y}

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
	w, h := g.Background.Size()
	for x := -w + int(g.Camera.PosX)%w - w; x < screen.Bounds().Max.X+w; x += w {
		for y := -h + int(g.Camera.PosY)%h - h; y < screen.Bounds().Max.Y+h; y += h {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x), float64(y))
			screen.DrawImage(g.Background, op)
		}
	}

	// Draw the platforms
	for _, p := range g.Platforms {
		p.Draw(screen, g.Camera.PosX, g.Camera.PosY)
	}

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
