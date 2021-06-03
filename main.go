/*
DINOGAME v1.0 by Maciej Sobczyk
*/

package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	screenHeight            = 600
	screenWidth             = 1100
	y_pos                   = 310
	GRAVITY                 = 9.81
	JV                      = 35
	bird_probability        = 0.98
	cactus_probability      = 0.75
	ModeTitle          Mode = iota
	ModeGame
	ModeGameOver
	titleFontSize = fontSize * 1.5
	fontSize      = 24
	smallFontSize = fontSize / 2
)

var (
	RUNNING  *ebiten.Image
	RUNNING2 *ebiten.Image
	JUMPING  *ebiten.Image
	DUCKING  *ebiten.Image
	DUCKING2 *ebiten.Image

	SMALL_CACTUS  *ebiten.Image
	LARGE_CACTUS  *ebiten.Image
	SMALL_CACTUS2 *ebiten.Image
	LARGE_CACTUS2 *ebiten.Image
	SMALL_CACTUS3 *ebiten.Image
	LARGE_CACTUS3 *ebiten.Image

	BIRD  *ebiten.Image
	BIRD2 *ebiten.Image
	CLOUD *ebiten.Image
	BG    *ebiten.Image

	titleArcadeFont font.Face
	arcadeFont      font.Face
	smallArcadeFont font.Face
)

//inicjalizacja dino,chmur, etc.
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
	SMALL_CACTUS2, _, err = ebitenutil.NewImageFromFile("assets/Cactus/SmallCactus2.png", ebiten.FilterDefault)

	if err != nil {
		log.Fatal(err)
	}
	SMALL_CACTUS3, _, err = ebitenutil.NewImageFromFile("assets/Cactus/SmallCactus3.png", ebiten.FilterDefault)

	if err != nil {
		log.Fatal(err)
	}
	LARGE_CACTUS, _, err = ebitenutil.NewImageFromFile("assets/Cactus/LargeCactus1.png", ebiten.FilterDefault)

	if err != nil {
		log.Fatal(err)
	}
	LARGE_CACTUS2, _, err = ebitenutil.NewImageFromFile("assets/Cactus/LargeCactus2.png", ebiten.FilterDefault)

	if err != nil {
		log.Fatal(err)
	}
	LARGE_CACTUS3, _, err = ebitenutil.NewImageFromFile("assets/Cactus/LargeCactus3.png", ebiten.FilterDefault)

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

//inicjalizacja z jakimi wartosciami ma uruchomic sie gra
func (g *Game) init() {
	g.Mode = ModeTitle
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
	g.Bird.X_VELOCITY = 100
	g.Bird.STATE = false

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
	g.Cactus.STATE = false
	g.Cactus.ILE = 0.3
	g.Cactus.IS_LARGE = false

	g.score = 0

}

//funkcja uzyta do animacji
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

type Mode int
type Game struct {
	Mode          Mode
	dino          Dinosaur
	Bird          entity
	Cloud         entity
	Cloud2        entity
	Cloud3        entity
	Cactus        entity
	gameoverCount int
	score         int
}

func random() float64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Float64()

}

func randomY() int {
	rand.Seed(time.Now().UnixNano())

	return (rand.Intn(350-220+1) + 220)

}
func (g *Game) Update(screen *ebiten.Image) error {
	//przemieszczanie sie chmur
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
	//tryby gry Title,Game,GameOver
	switch g.Mode {
	case ModeTitle:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.Mode = ModeGame
		}
	case ModeGame:
		g.score++
		//przemieszczanie ptaka
		if g.Bird.STATE {
			if g.Bird.X_POS < 0 {
				g.Bird.X_POS = 1600
				g.Bird.STATE = false
			}
			g.Bird.X_POS -= g.Bird.X_VELOCITY * 0.4

		} else {
			if random() > bird_probability {
				g.Bird.STATE = true
				g.Bird.Y_POS = float64(randomY())
			}
		}
		//przemieszczanie kaktusow
		if g.Cactus.STATE {
			if g.Cactus.X_POS < 0 {
				g.Cactus.X_POS = 2000
				g.Cactus.STATE = false

			}
			g.Cactus.X_POS -= g.Cactus.X_VELOCITY * 0.4
		} else {
			if random() > cactus_probability {
				g.Cactus.ILE = random()
				if random() > 0.90 {
					g.Cactus.IS_LARGE = true
				} else {
					g.Cactus.IS_LARGE = false
				}
				if g.Cactus.IS_LARGE {
					if g.Cactus.ILE < 0.7 {
						g.Cactus.IMAGE = LARGE_CACTUS
					} else if g.Cactus.ILE > 0.7 && g.Cactus.ILE < 0.90 {
						g.Cactus.IMAGE = LARGE_CACTUS2
					} else if g.Cactus.ILE > 0.90 {
						g.Cactus.IMAGE = LARGE_CACTUS3
					}
				} else {
					if g.Cactus.ILE < 0.8 {
						g.Cactus.IMAGE = SMALL_CACTUS
					} else if g.Cactus.ILE > 0.8 && g.Cactus.ILE < 0.90 {
						g.Cactus.IMAGE = SMALL_CACTUS2
					} else if g.Cactus.ILE > 0.90 {
						g.Cactus.IMAGE = SMALL_CACTUS3
					}
				}

				g.Cactus.STATE = true

			}

		}

		//reakcje na klawisze dla dinozaura
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.dino.JUMP_STATE = true
			g.dino.DUCK_STATE = false
			g.dino.RUN_STATE = false
		} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
			g.dino.JUMP_STATE = false
			g.dino.DUCK_STATE = true
			g.dino.RUN_STATE = false
		} else {
			g.dino.JUMP_STATE = false
			g.dino.DUCK_STATE = false
			g.dino.RUN_STATE = true
		}

		//reakcje dinozaura na poszczegolne stany
		if g.dino.Y_POS > 310 {

			if g.dino.JUMP_VEL > 0 {
				g.dino.JUMP_VEL -= 40 * 2

			}

			g.dino.JUMP_STATE = false
		} else {
			g.dino.JUMP_VEL = 0
			g.dino.Y_POS = 310
		}

		if g.dino.JUMP_STATE {
			g.dino.IMAGE = RUNNING
			g.dino.IMAGE2 = JUMPING
			g.dino.JUMP_VEL += 75
			g.dino.Y_POS += g.dino.JUMP_VEL
			g.dino.JUMP_VEL += 20

		}
		g.dino.Y_POS += g.dino.JUMP_VEL
		if g.dino.DUCK_STATE {
			g.dino.IMAGE = DUCKING
			g.dino.IMAGE2 = DUCKING2
			g.dino.STEP_INDEX++
			time.Sleep(100 * time.Millisecond)

		}

		if g.dino.RUN_STATE {
			g.dino.IMAGE = RUNNING
			g.dino.IMAGE2 = RUNNING2
			g.dino.STEP_INDEX++
			time.Sleep(100 * time.Millisecond)

		}

		if g.dino.STEP_INDEX >= 10 {
			g.dino.STEP_INDEX = 0
		}

		if g.Collision() {
			g.Mode = ModeGameOver
			g.gameoverCount = 30
		}
	case ModeGameOver:
		if g.gameoverCount > 0 {
			g.gameoverCount--
		}
		if g.gameoverCount == 0 && ebiten.IsKeyPressed(ebiten.KeySpace) {
			g.init()
			g.Mode = ModeTitle
		}
	}

	return nil
}

//inicjalizacja czcionki
func init() {
	tt, err := opentype.Parse(fonts.PressStart2P_ttf)
	if err != nil {
		log.Fatal(err)
	}
	const dpi = 72
	titleArcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    titleFontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	arcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    fontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	smallArcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    smallFontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Collision() bool {
	if g.Mode != ModeGame {
		return false

	}
	w_dino, h_dino := g.dino.IMAGE.Size()
	w_bird, h_bird := g.Bird.IMAGE.Size()
	w_cactus, h_cactus := g.Cactus.IMAGE.Size()

	if g.Bird.X_POS-(float64(w_bird)*0.3) <= g.dino.X_POS+(float64(w_dino)*0.4) {
		if g.dino.Y_POS-(float64(h_dino)*0.4) >= g.Bird.Y_POS+(float64(h_bird)*0.9) {
			return true
		}
	}
	if g.Cactus.IS_LARGE {
		if g.Cactus.X_POS-(float64(w_cactus)*0.3) <= g.dino.X_POS+(float64(w_dino)*0.4) {
			if g.dino.Y_POS-(float64(h_dino)*0.4) <= g.Cactus.Y_POS+(float64(h_cactus)*0.4) {
				return true
			}
		}
	} else {
		if g.Cactus.X_POS-(float64(w_cactus)*0.3) <= g.dino.X_POS+(float64(w_dino)*0.4) {
			if g.dino.Y_POS-(float64(h_dino)*0.4) <= g.Cactus.Y_POS+(float64(h_cactus)*0.4) {
				return true
			}
		}
	}

	return false
}
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0xff, 0xff, 0xff, 0xea})
	g.DrawBG(screen)
	g.DrawCloud(screen)
	g.DrawCloud2(screen)
	g.DrawCloud3(screen)

	var titleTexts []string
	var texts []string
	switch g.Mode {
	case ModeTitle:
		titleTexts = []string{"DINOGAME v1.0"}
		texts = []string{"WCIŚNIJ SPACE ABY ROZPOCZĄĆ GRĘ"}
	case ModeGameOver:
		texts = []string{"KONIEC GRY!"}
	}
	for i, l := range titleTexts {
		x := (screenWidth - len(l)*titleFontSize) / 2
		text.Draw(screen, l, titleArcadeFont, x, (i+4)*titleFontSize, color.Black)
	}
	for i, l := range texts {
		x := (screenWidth - len(l)*fontSize) / 2
		text.Draw(screen, l, arcadeFont, x, (i+4)*fontSize, color.Black)
	}
	if g.Mode == ModeTitle {
		msg := []string{
			"DINOGAME by Maciej Sobczyk",
		}
		for i, l := range msg {
			x := (screenWidth - len(l)*smallFontSize) / 2
			text.Draw(screen, l, smallArcadeFont, x, screenHeight-4+(i-1)*smallFontSize, color.Black)
		}
	}
	g.DrawDinoRun(screen)
	g.DrawBird(screen)
	g.DrawCactus(screen)
	scoreStr := fmt.Sprintf("%06d", g.score)
	text.Draw(screen, scoreStr, arcadeFont, screenWidth-len(scoreStr)*fontSize, fontSize, color.Black)

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
	if g.Cactus.IS_LARGE {
		op.GeoM.Translate(float64(g.Cactus.X_POS+float64(w)), float64(screenHeight-g.Cactus.Y_POS+float64(h+99)))
	} else {
		op.GeoM.Translate(float64(g.Cactus.X_POS+float64(w)), float64(screenHeight-g.Cactus.Y_POS+float64(h+128)))
	}

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
