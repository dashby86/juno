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

	platform3 := NewPlatform(float64(screenWidth)/2-250, float64(screenHeight)/2+float64(junoHeight)+50, 500, 30, colornames.Red)
	platform1 := NewPlatform(200, 900, 500, 30, colornames.Red)
	platform2 := NewPlatform(800, 700, 500, 30, colornames.Black)

	j := juno.NewJuno(junoImg, float64(screenWidth), float64(screenHeight), 4)
	j.Speed = 4 // set Juno's speed to 10 pixels per frame

	// Start Juno on the first platform
	j.X = float64(platform1.X)
	j.Y = platform1.Y - float64(junoHeight)

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

	game.Platforms = []*Platform{platform1, platform2, platform3}

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

	// Check for collisions with platforms and apply gravity
	nextY := g.Juno.Y + g.Juno.VelY
	g.Juno.Grounded = false
	for _, platform := range g.Platforms {
		if nextY+float64(junoHeight) > platform.Y && nextY < platform.Y+platform.Height &&
			g.Juno.X+float64(junoWidth) > platform.X && g.Juno.X < platform.X+platform.Width {
			g.Juno.Grounded = true
			g.Juno.VelY = 0
			g.Juno.Y = platform.Y - float64(junoHeight)
			break
		}
	}

	if !g.Juno.Grounded {
		g.Juno.Y = nextY
	}

	// Prevent Juno from falling below the screen
	/**
	if g.Juno.Y > screenHeight {
		g.Juno.Y = screenHeight - float64(junoHeight)
		g.Juno.Grounded = true
		g.Juno.VelY = 0
	}

	*/

	/**
	x, y := g.Juno.GetPosition()

	g.Camera.PosX = x + float64(g.Juno.GetWidth())/2 - float64(screenWidth)/2

	if newY := y + float64(g.Juno.GetHeight())/2 - float64(screenHeight)/2; newY < screenHeight-float64(junoHeight) {
		g.Camera.PosY = newY
	} else {
		g.Camera.PosY = screenHeight - float64(junoHeight)
	}


	*/
	x, y := g.Juno.GetPosition()

	halfScreenWidth := float64(screenWidth) / 2
	halfScreenHeight := float64(screenHeight) / 2

	targetX := x + float64(g.Juno.GetWidth())/2 - halfScreenWidth
	targetY := y + float64(g.Juno.GetHeight())/2 - halfScreenHeight

	lerpFactor := 0.1 // You can adjust this value to change the camera's smoothness

	g.Camera.PosX = g.Camera.PosX + (targetX-g.Camera.PosX)*lerpFactor
	g.Camera.PosY = g.Camera.PosY + (targetY-g.Camera.PosY)*lerpFactor

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
	//w, h := g.Background.Size()
	//tileOffsetX := int(g.Camera.PosX) % w
	//tileOffsetY := int(g.Camera.PosY) % h

	/**
	for x := -tileOffsetX - w; x < screenWidth+w; x += w {
		for y := -tileOffsetY - h; y < screenHeight+h; y += h {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x)-g.Camera.PosX, float64(y)-g.Camera.PosY)
			screen.DrawImage(g.Background, op)
		}
	}

	*/

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-g.Camera.PosX, -g.Camera.PosY)
	screen.DrawImage(g.Background, op)

	// Draw the platforms
	for _, p := range g.Platforms {
		p.Draw(screen, g.Camera.PosX, g.Camera.PosY)
		//fmt.Printf("P:%v,%v\n", p.X, p.Y)

		//ebitenutil.DebugPrintAt(screen, fmt.Sprintf("P:%v,%v", p.X, p.Y), int(p.X-g.Camera.PosX), int(p.Y-g.Camera.PosY))

	}

	// Draw the enemies
	for _, e := range g.Enemies {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(e.Position.X-g.Camera.PosX, e.Position.Y-g.Camera.PosY)
		screen.DrawImage(e.Image, op)
	}

	x, y := g.Juno.GetPosition()
	//op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x-g.Camera.PosX, y-g.Camera.PosY)
	screen.DrawImage(g.Juno.Image.(*ebiten.Image), op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// Set the size of the screen based on the outsideWidth and outsideHeight parameters.
	// For example, you could set the screen size to be half the size of the outside area:
	screenWidth, screenHeight = outsideWidth/2, outsideHeight/2
	return screenWidth, screenHeight
}
