package main

import (
	"image"
	_ "image/png"
	"log"

	//"github.com/hajimehoshi/ebiten/ebitenutil"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Game implements ebiten.Game interface.
type Game struct {
	keys []ebiten.Key
}

var img *ebiten.Image
var paddle1 Paddle
var paddle2 Paddle
var ball Ball

var dirX float64 = 1
var dirY float64 = 1

func init() {
	var err error
	img, _, err = ebitenutil.NewImageFromFile("images/sprites.png")
	if err != nil {
		log.Fatal(err)
	}

	paddle1.name = "Player 1"
	paddle1.positionX = 0
	paddle1.positionY = 45
	paddle1.image = img.SubImage(image.Rect(0, 0, 50, 140)).(*ebiten.Image)

	paddle2.name = "Player 2"
	paddle2.positionX = 270
	paddle2.positionY = 45
	paddle2.image = img.SubImage(image.Rect(100, 0, 50, 140)).(*ebiten.Image)

	ball.positionX = 320 * .5
	ball.positionY = 240 * .5
	ball.image = img.SubImage(image.Rect(100, 0, 150, 50)).(*ebiten.Image)
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	// Write your game's logical update.
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])

	ball.positionX += dirX
	ball.positionY += dirY

	if ball.positionX >= 320 {
		dirX = -1
	}
	if ball.positionX <= 0 {
		dirX = 1
	}

	if ball.positionY >= 240 {
		dirY = -1
	}
	if ball.positionY <= 0 {
		dirY = 1
	}

	for _, key := range g.keys {

		if key == ebiten.KeyS {
			paddle1.positionY += 1
		}

		if key == ebiten.KeyW {
			paddle1.positionY += -1
		}

		if key == ebiten.KeyArrowDown {
			paddle2.positionY += 1
		}

		if key == ebiten.KeyArrowUp {
			paddle2.positionY += -1
		}
	}

	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {

	//screen.Fill(color.RGBA{200, 0, 0, 0xff})
	// Write your game's rendering.
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(paddle1.positionX, paddle1.positionY)

	op2 := &ebiten.DrawImageOptions{}
	op2.GeoM.Translate(paddle2.positionX, paddle2.positionY) //op.GeoM.Scale(1.5, 1)

	op3 := &ebiten.DrawImageOptions{}
	op3.GeoM.Translate(ball.positionX, ball.positionY) //op.GeoM.Scale(1.5, 1)

	screen.DrawImage(paddle1.image, op)
	screen.DrawImage(paddle2.image, op2)
	screen.DrawImage(ball.image, op3)

}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	game := &Game{}
	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Your game's title")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
