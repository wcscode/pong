package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Paddle struct {
	name      string
	positionX float64
	positionY float64
	image     *ebiten.Image
}

type Ball struct {
	positionX float64
	positionY float64
	image     *ebiten.Image
}
