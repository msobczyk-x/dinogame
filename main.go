package main

import (
	"image/color"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
)

const (
	screenHeight = 600
	screenWidth  = 1100
	y_pos        = 310
	GRAVITY      = 9.81
	JV           = 35
)

var (
	RUNNING  *ebiten.Image
	RUNNING2 *ebiten.Image
	JUMPING  *ebiten.Image
	DUCKING  *ebiten.Image
	DUCKING2 *ebiten.Image

	SMALL_CACTUS *ebiten.Image
	LARGE_CACTUS *ebiten.Image

	BIRD  *ebiten.Image
	BIRD2 *ebiten.Image
	CLOUD *ebiten.Image
	BG    *ebiten.Image
)

func init() {

	var err error
	RUNNING, _, err = ebitenutil.NewImageFromFile("assets/Dino/DinoRun1.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	RUNNING2, _, err = ebitenutil.NewImageFromFile("assets/Dino/DinoRun2.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	JUMPING, _, err = ebitenutil.NewImageFromFile("assets/Dino/DinoJump.png", ebiten.FilterDefault)

	if err != nil {
		log.Fatal(err)
	}
	DUCKING, _, err = ebitenutil.NewImageFromFile("assets/Dino/DinoDuck1.png", ebiten.FilterDefault)

	if err != nil {
		log.Fatal(err)
	}
	DUCKING2, _, err = ebitenutil.NewImageFromFile("assets/Dino/DinoDuck2.png", ebiten.FilterDefault)

	if err != nil {
		log.Fatal(err)
	}
	SMALL_CACTUS, _, err = ebitenutil.NewImageFromFile("assets/Cactus/SmallCactus1.png", ebiten.FilterDefault)

	if err != nil {
		log.Fatal(err)
	}
	LARGE_CACTUS, _, err = ebitenutil.NewImageFromFile("assets/Cactus/LargeCactus1.png", ebiten.FilterDefault)

	if err != nil {
		log.Fatal(err)
	}
	BIRD, _, err = ebitenutil.NewImageFromFile("assets/Bird/Bird1.png", ebiten.FilterDefault)

	if err != nil {
		log.Fatal(err)
	}
	BIRD2, _, err = ebitenutil.NewImageFromFile("assets/Bird/Bird2.png", ebiten.FilterDefault)

	if err != nil {
		log.Fatal(err)
	}
	CLOUD, _, err = ebitenutil.NewImageFromFile("assets/Other/Cloud.png", ebiten.FilterDefault)

	if err != nil {
		log.Fatal(err)
	}
	BG, _, err = ebitenutil.NewImageFromFile("assets/Other/Track.png", ebiten.FilterDefault)

	if err != nil {
		log.Fatal(err)
	}
}

func NewGame() *Game {
	g := &Game{}
	g.init()

	return g
}

func (g *Game) init() {
	g.dino.X_POS = 80
	g.dino.Y_POS = 310
	g.dino.Y_POS_DUCK = 340
	g.dino.JUMP_VEL = 240
	g.dino.DUCK_STATE = false
	g.dino.RUN_STATE = true
	g.dino.JUMP_STATE = false
	g.dino.STEP_INDEX = 0
	g.dino.IMAGE = RUNNING
	g.dino.IMAGE2 = RUNNING2

	g.Bird.X_POS = 1700
	g.Bird.Y_POS = 300
	g.Bird.IMAGE = BIRD
	g.Bird.IMAGE2 = BIRD2
	g.Bird.X_VELOCITY = 120

	g.Cloud.X_POS = 1600
	g.Cloud.Y_POS = 550
	g.Cloud.IMAGE = CLOUD
	g.Cloud.X_VELOCITY = 60

	g.Cloud2.X_POS = 1700
	g.Cloud2.Y_POS = 700
	g.Cloud2.IMAGE = CLOUD
	g.Cloud2.X_VELOCITY = 90

	g.Cloud3.X_POS = 1200
	g.Cloud3.Y_POS = 630
	g.Cloud3.IMAGE = CLOUD
	g.Cloud3.X_VELOCITY = 40

	g.Cactus.X_POS = 1400
	g.Cactus.Y_POS = 330
	g.Cactus.IMAGE = SMALL_CACTUS
	g.Cactus.X_VELOCITY = 240

}
func mod(a, b int) int {
	m := a % b
	if a < 0 && b < 0 {
		m -= b
	}
	if a < 0 && b > 0 {
		m += b
	}
	return m
}

type Game struct {
	dino   Dinosaur
	Bird   entity
	Cloud  entity
	Cloud2 entity
	Cloud3 entity
	Cactus entity
}

func (g *Game) Update(screen *ebiten.Image) error {
	//poruszanie sie ptakow i chmur, kaktusa
	if g.Cloud.X_POS < 0 {
		g.Cloud.X_POS = 1600
	}
	g.Cloud.X_POS -= g.Cloud.X_VELOCITY * 0.4

	if g.Cloud2.X_POS < 0 {
		g.Cloud2.X_POS = 1700
	}
	g.Cloud2.X_POS -= g.Cloud2.X_VELOCITY * 0.4

	if g.Cloud3.X_POS < 0 {
		g.Cloud3.X_POS = 1200
	}
	g.Cloud3.X_POS -= g.Cloud3.X_VELOCITY * 0.4

	if g.Bird.X_POS < 0 {
		g.Bird.X_POS = 1600
	}
	g.Bird.X_POS -= g.Bird.X_VELOCITY * 0.4

	if g.Cactus.X_POS < 0 {
		g.Cactus.X_POS = 2000
	}
	g.Cactus.X_POS -= g.Cactus.X_VELOCITY * 0.4

	//reakcje na klawisze dla dinozaura
	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		g.dino.JUMP_STATE = true
		g.dino.DUCK_STATE = false
		g.dino.RUN_STATE = false
	} else if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.dino.JUMP_STATE = false
		g.dino.DUCK_STATE = true
		g.dino.RUN_STATE = false
	} else {
		g.dino.JUMP_STATE = false
		g.dino.DUCK_STATE = false
		g.dino.RUN_STATE = true
	}

	//fmt.Println("Duck state: ", g.dino.DUCK_STATE)
	//fmt.Println("Run state: ", g.dino.RUN_STATE)
	//fmt.Println("Jump state: ", g.dino.JUMP_STATE)
	//fmt.Println("jumpvel", g.dino.JUMP_VEL)
	if g.dino.Y_POS > 310 {

		if g.dino.JUMP_VEL > 0 { // a może
			g.dino.JUMP_VEL -= 40 * 2 // terminal velocity // nie no git ale interval za szybki w sensie -= albo przemnozyc przez cos +1 gravity trzeba dodac
			// g.dino.Jump_vel += 4
			//fmt.Println("TEN IF Pierwszy")

		}

		g.dino.JUMP_STATE = false
	} else {
		g.dino.JUMP_VEL = 0
		g.dino.Y_POS = 310
	}

	if g.dino.JUMP_STATE {
		g.dino.IMAGE = RUNNING
		g.dino.IMAGE2 = JUMPING
		g.dino.JUMP_VEL += 75 // nie no bo jak jest + to idzie w dol += to w dol a jak -= to w gore # niemożliwe no odpal NO I DZIAŁA
		g.dino.Y_POS += g.dino.JUMP_VEL
		g.dino.JUMP_VEL += 20
		//fmt.Println("TEN IF DRUGI")

	}
	g.dino.Y_POS += g.dino.JUMP_VEL // niech jego lokalizacja Y się aktualizuje cały czas, zamieniamy tylko jumpvel
	if g.dino.DUCK_STATE {
		g.dino.IMAGE = DUCKING
		g.dino.IMAGE2 = DUCKING2
		g.dino.STEP_INDEX++
		time.Sleep(100 * time.Millisecond)
		//fmt.Println(g.dino.STEP_INDEX)

	}

	if g.dino.RUN_STATE {
		g.dino.IMAGE = RUNNING
		g.dino.IMAGE2 = RUNNING2
		g.dino.STEP_INDEX++
		time.Sleep(100 * time.Millisecond)
		//fmt.Println(g.dino.STEP_INDEX)
	}

	if g.dino.STEP_INDEX >= 10 {
		g.dino.STEP_INDEX = 0
	}
	return nil
}
func (g *Game) Collision() bool {

	return false
}
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0xff, 0xff, 0xff, 0xea})
	g.DrawBG(screen)
	g.DrawDinoRun(screen)
	g.DrawBird(screen)
	g.DrawCloud(screen)
	g.DrawCloud2(screen)
	g.DrawCloud3(screen)
	g.DrawCactus(screen)

}
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) DrawCloud(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	w, h := g.Cloud.IMAGE.Size()
	op.GeoM.Translate(-float64(w)/2.0, -float64(h)/2.0)
	op.GeoM.Translate(float64(g.Cloud.X_POS+float64(w)), float64(screenHeight-g.Cloud.Y_POS+float64(h+128)))
	screen.DrawImage(g.Cloud.IMAGE, op)

}

func (g *Game) DrawCloud2(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	w, h := g.Cloud2.IMAGE.Size()
	op.GeoM.Translate(-float64(w)/2.0, -float64(h)/2.0)
	op.GeoM.Translate(float64(g.Cloud2.X_POS+float64(w)), float64(screenHeight-g.Cloud2.Y_POS+float64(h+128)))
	screen.DrawImage(g.Cloud2.IMAGE, op)

}

func (g *Game) DrawCloud3(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	w, h := g.Cloud3.IMAGE.Size()
	op.GeoM.Translate(-float64(w)/2.0, -float64(h)/2.0)
	op.GeoM.Translate(float64(g.Cloud3.X_POS+float64(w)), float64(screenHeight-g.Cloud3.Y_POS+float64(h+128)))
	screen.DrawImage(g.Cloud3.IMAGE, op)

}

func (g *Game) DrawCactus(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	w, h := g.Cactus.IMAGE.Size()
	op.GeoM.Translate(-float64(w)/2.0, -float64(h)/2.0)
	op.GeoM.Translate(float64(g.Cactus.X_POS+float64(w)), float64(screenHeight-g.Cactus.Y_POS+float64(h+128)))
	screen.DrawImage(g.Cactus.IMAGE, op)

}

func (g *Game) DrawBird(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	w, h := g.Bird.IMAGE.Size()
	op.GeoM.Translate(-float64(w)/2.0, -float64(h)/2.0)
	op.GeoM.Translate(float64(g.Bird.X_POS+float64(w)), float64(screenHeight-g.Bird.Y_POS+float64(h+128)))
	op.GeoM.Scale(0.7, 0.7)
	if mod(g.dino.STEP_INDEX, 2) == 0 {
		screen.DrawImage(g.Bird.IMAGE, op)
	} else {
		screen.DrawImage(g.Bird.IMAGE2, op)

	}

}
func (g *Game) DrawDinoRun(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	w, h := g.dino.IMAGE.Size()
	if g.dino.IMAGE == DUCKING {
		op.GeoM.Translate(-float64(w)/2.0, -float64(h)/2.0)
	} else if g.dino.IMAGE == RUNNING {
		op.GeoM.Translate(-float64(w), -float64(h))
	}

	op.GeoM.Translate(float64(g.dino.X_POS+float64(w)), float64(screenHeight-g.dino.Y_POS+float64(h+128)))
	if mod(g.dino.STEP_INDEX, 2) == 0 {
		screen.DrawImage(g.dino.IMAGE, op)
		//fmt.Println("Width:", g.dino.X_POS)
		//fmt.Println("Height:", g.dino.Y_POS)
	} else {
		screen.DrawImage(g.dino.IMAGE2, op)

	}

}

func (g *Game) DrawBG(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(0, screenHeight-128)
	screen.DrawImage(BG, op)
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("DinoGame")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
