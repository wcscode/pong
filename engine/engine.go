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
	VelocityX    float64
	VelocityY    float64
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

type DrawerUpdater interface {
	//Draw(screen *ebiten.Image)
	Update()
}

/*func (gOb *GameObject) Draw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(gOb.PositionX, gOb.PositionY)

	screen.DrawImage(gOb.Sprite.Image, op)
}*/

/*func (s *Scene) Draw(screen *ebiten.Image) {

	screen.Fill(color.RGBA{200, 0, 0, 0xff})
}*/



func (s *Sprite) LoadAndCutImage(img *ebiten.Image, initialX int, initialY int) {

	s.Image = img.SubImage(image.Rect(initialX, initialY, initialX+s.ImageWidth, initialY+s.ImageHeight)).(*ebiten.Image)
}


func (gOb *GameObject) AddVelocity(x float64, y float64) {

	 gOb.VelocityX += x
	 gOb.VelocityY += y
}

func (gOb *GameObject) InvertVelocity(x bool, y bool) {
	
	if(x) {
		gOb.VelocityX *= -1
	}

	if(y) {
		gOb.VelocityY *= -1
	}
}


/*func (bc *BoxCollision) SetBorders(offset int) {
	//bc.x0 =
}*/
