package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type Dinosaur struct {
	X_POS      float64
	Y_POS      float64
	Y_POS_DUCK float64
	JUMP_VEL   float64

	STEP_INDEX int
	IMAGE      *ebiten.Image
	IMAGE2     *ebiten.Image
	DUCK_STATE bool
	RUN_STATE  bool
	JUMP_STATE bool
}

func (d Dinosaur) duck() {
	d.IMAGE = DUCKING
	d.Y_POS = d.Y_POS_DUCK
	d.STEP_INDEX += 1
}

func (d Dinosaur) run() {
	d.STEP_INDEX += 1
}

func (d Dinosaur) jump() {
	d.IMAGE = JUMPING
	if d.JUMP_STATE {
		d.Y_POS -= d.JUMP_VEL * 4
		d.JUMP_VEL -= 0.8
		if d.JUMP_VEL > -JV {
			d.JUMP_STATE = false
			d.JUMP_VEL = JV
		}
	}
}

func (d *Dinosaur) update() {
	if d.DUCK_STATE {
		d.duck()

	}
	if d.RUN_STATE {
		d.run()

	}
	if d.JUMP_STATE {
		d.jump()

	}

	if d.STEP_INDEX >= 10 {
		d.STEP_INDEX = 0
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyW) && !d.JUMP_STATE {

		d.DUCK_STATE = false
		d.RUN_STATE = false
		d.JUMP_STATE = true
	} else if inpututil.IsKeyJustPressed(ebiten.KeyS) && !d.JUMP_STATE {

		d.DUCK_STATE = true
		d.RUN_STATE = false
		d.JUMP_STATE = false
	} else if !(d.JUMP_STATE || inpututil.IsKeyJustPressed((ebiten.KeyS))) {
		d.DUCK_STATE = false
		d.RUN_STATE = true
		d.JUMP_STATE = false

	}
}
