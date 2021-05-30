package Dino

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type Mode int

const (
	ModeTitle Mode = iota
	ModeGame
	ModeGameOver
	screenWidth  = 1024
	screenHeight = 512
	JUMP_VEL     = 8.5
)

// Game implements ebiten.Game interface.
type Game struct {
	mode Mode

	x16     int
	y16     float64
	vy16    float64
	gravity float64

	gameoverCount int
}

var (
	bgImg    *ebiten.Image
	carImg   *ebiten.Image
	horseImg *ebiten.Image
)

func jump() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeySpace)

}
func NewGame() *Game {
	g := &Game{}
	g.init()
	return g
}

func (g *Game) init() {
	g.x16 = 80
	g.y16 = 310
	g.vy16 = JUMP_VEL
	g.gravity = 4

}

func init() {
	var err error
	bgImg, _, err = ebitenutil.NewImageFromFile("bgImg.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	carImg, _, err = ebitenutil.NewImageFromFile("car2.png", ebiten.FilterDefault)

	if err != nil {
		log.Fatal(err)
	}
	horseImg, _, err = ebitenutil.NewImageFromFile("horse.png", ebiten.FilterDefault)

	if err != nil {
		log.Fatal(err)
	}

}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update(*ebiten.Image) error {
	// Write your game's logical update.

	if jump() {
		g.y16 -= g.vy16 * 4
		g.vy16 -= 0.8
	}
	if g.vy16 < -JUMP_VEL {
		g.vy16 = JUMP_VEL
	}

	fmt.Println("g.vy16", g.vy16)
	fmt.Println("y16", g.y16)

	/* if g.hit() {

		g.mode = ModeGameOver
		g.gameoverCount = 30
	} */
	/* case ModeGameOver:
	if g.gameoverCount > 0 {
		g.gameoverCount--
	} */
	/* if g.gameoverCount == 0 && jump() {
		g.init()
		g.mode = ModeTitle
	} */

	return nil

}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	// Write your game's rendering.
	screen.DrawImage(bgImg, nil)

	g.drawHorse(screen)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
func (g *Game) drawHorse(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	w, h := horseImg.Size()
	op.GeoM.Translate(float64(w)/3.0, float64(screenHeight-h)/1.3)
	op.GeoM.Translate(float64(g.x16/16.0), float64(g.y16/16.0))
	screen.DrawImage(horseImg, op)
}

func main() {

	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("DinoGame")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
