package main

import (
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
type BoxCollision struct {
}

var img *ebiten.Image
var paddle1 GameObject
var paddle2 GameObject
var ball GameObject

var dirX float64 = 1
var dirY float64 = 1

func init() {
	var err error
	img, _, err = ebitenutil.NewImageFromFile("images/sprites.png")
	if err != nil {
		log.Fatal(err)
	}

	paddle1.Sprite.Name = "Player 1"
	paddle1.Sprite.PositionX = 0
	paddle1.Sprite.PositionY = 45
	paddle1.Sprite.ImageWidth = 50
	paddle1.Sprite.ImageHeight = 145
	paddle1.Sprite.LoadAndCutImage(img, 0, 0)

	paddle2.Sprite.Name = "Player 2"
	paddle2.Sprite.PositionX = 270
	paddle2.Sprite.PositionY = 45
	paddle2.Sprite.ImageWidth = 50
	paddle2.Sprite.ImageHeight = 145
	paddle2.Sprite.LoadAndCutImage(img, 50, 0)

	ball.Sprite.PositionX = 320 * .5
	ball.Sprite.PositionY = 240 * .5
	ball.Sprite.ImageWidth = 50
	ball.Sprite.ImageHeight = 50
	ball.Sprite.LoadAndCutImage(img, 100, 0)
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	// Write your game's logical update.
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])

	ball.PositionX += dirX
	ball.PositionY += dirY

	if ball.PositionX >= 320 {
		dirX = -1
	}
	if ball.PositionX <= 0 {
		dirX = 1
	}

	if ball.PositionY >= 240 {
		dirY = -1
	}
	if ball.PositionY <= 0 {
		dirY = 1
	}

	for _, key := range g.keys {

		if key == ebiten.KeyS {
			paddle1.PositionY += 1
		}

		if key == ebiten.KeyW {
			paddle1.PositionY += -1
		}

		if key == ebiten.KeyArrowDown {
			paddle2.PositionY += 1
		}

		if key == ebiten.KeyArrowUp {
			paddle2.PositionY += -1
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
	op.GeoM.Translate(paddle1.PositionX, paddle1.PositionY)

	op2 := &ebiten.DrawImageOptions{}
	op2.GeoM.Translate(paddle2.PositionX, paddle2.PositionY) //op.GeoM.Scale(1.5, 1)

	op3 := &ebiten.DrawImageOptions{}
	op3.GeoM.Translate(ball.PositionX, ball.PositionY) //op.GeoM.Scale(1.5, 1)

	//fmt.Println(op, op2, op3)
	screen.DrawImage(paddle1.Image, op)
	screen.DrawImage(paddle2.Image, op2)
	screen.DrawImage(ball.Image, op3)

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
