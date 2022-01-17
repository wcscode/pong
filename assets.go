package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	Name        string
	PositionX   float64
	PositionY   float64
	ImageWidth  int
	ImageHeight int
	Image       *ebiten.Image
}

type BoxCollision struct {
	x0 float64
	y0 float64
	x1 float64
	y1 float64
}

type GameObject struct {
	Sprite       Sprite
	BoxCollision BoxCollision
}

type GamesObjects struct {
	Sprites []*Sprite
}

func (s *Sprite) LoadAndCutImage(img *ebiten.Image, initialX int, initialY int) {

	s.Image = img.SubImage(image.Rect(initialX, initialY, initialX+s.ImageWidth, initialY+s.ImageHeight)).(*ebiten.Image)
}

func (bc *BoxCollision) SetBorders(offset int) {
	//bc.x0 =
}
