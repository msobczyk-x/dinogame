package main

import "github.com/hajimehoshi/ebiten"

type entity struct {
	X_POS      float64
	Y_POS      float64
	IMAGE      *ebiten.Image
	IMAGE2     *ebiten.Image
	X_VELOCITY float64
}
