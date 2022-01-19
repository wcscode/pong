package engine

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

const ScreenWidth float64 = 640
const ScreenHeight float64 = 480

type Scener interface {
	GetName() string
	GetActive() bool
	SetActive(active bool) 
	Draw(screen *ebiten.Image)
	Update(e []ebiten.Key) error 
	//Update(g *ebiten.Game) error
}

type Scene struct {
	Name   string
	Active bool	
    GamesObjects []*GameObject
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


//Define list of scenes

var ScenesBehaviors []Scener


func SetActiveScene(name string) {

	for _, scene := range ScenesBehaviors {

		if(scene.GetName() == name) {

			scene.SetActive(true)
		}else {
			scene.SetActive(false)
		}
	}
}

func (gOb *GameObject) Draw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(gOb.PositionX, gOb.PositionY)

	screen.DrawImage(gOb.Sprite.Image, op)
}




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
