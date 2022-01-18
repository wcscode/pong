package engine

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Scene struct {
	Name   string
	Active bool
}

type GameObject struct {
	PositionX    float64
	PositionY    float64
	Sprite       Sprite
	BoxCollision BoxCollision
}

type Sprite struct {
	Name        string
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

type Drawer interface {
	//	Draw(screen *ebiten.Image)
}

type Updater interface {
	Update()
}

func (gOb *GameObject) Draw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(gOb.PositionX, gOb.PositionY)

	screen.DrawImage(gOb.Sprite.Image, op)
}

func (s *Sprite) LoadAndCutImage(img *ebiten.Image, initialX int, initialY int) {

	s.Image = img.SubImage(image.Rect(initialX, initialY, initialX+s.ImageWidth, initialY+s.ImageHeight)).(*ebiten.Image)
}

func (bc *BoxCollision) SetBorders(offset int) {
	//bc.x0 =
}
