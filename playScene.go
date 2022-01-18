package main

import (
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/wcscode/pong/engine"
)

type PlayScene engine.Scene

var img *ebiten.Image

//var scene  engine.Scene
var paddle1 engine.GameObject
var paddle2 engine.GameObject
var ball engine.GameObject

func (s *PlayScene) Init(gamesObjects []*engine.GameObject) {

	var err error
	img, _, err = ebitenutil.NewImageFromFile("images/sprites.png")
	if err != nil {
		log.Fatal(err)
	}
	
	paddle1.Sprite.Name = "Player 1"
	paddle1.PositionX = 0
	paddle1.PositionY = 45
	paddle1.Sprite.ImageWidth = 50
	paddle1.Sprite.ImageHeight = 145
	paddle1.Sprite.LoadAndCutImage(img, 0, 0)

	gamesObjects = append(gamesObjects, &paddle1)

	paddle2.Sprite.Name = "Player 2"
	paddle2.PositionX = 270
	paddle2.PositionY = 45
	paddle2.Sprite.ImageWidth = 50
	paddle2.Sprite.ImageHeight = 145
	paddle2.Sprite.LoadAndCutImage(img, 50, 0)

	gamesObjects = append(gamesObjects, &paddle2)

	ball.PositionX = 320 * .5
	ball.PositionY = 240 * .5
	ball.VelocityX = 1
	ball.VelocityY = 1
	ball.Sprite.ImageWidth = 50
	ball.Sprite.ImageHeight = 50
	ball.Sprite.LoadAndCutImage(img, 100, 0)

	gamesObjects = append(gamesObjects, &ball)
}

func (s *PlayScene) Update(g *Game) error {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])

	ball.PositionX += ball.VelocityX
	ball.PositionY += ball.VelocityY

	if ball.PositionX >= 320 {
	 	ball.InvertVelocity(true, false)
	}
	if ball.PositionX <= 0 {
		ball.InvertVelocity(true, false)
	}

	if ball.PositionY >= 240 {
		ball.InvertVelocity(false, true)
	}
	if ball.PositionY <= 0 {
		ball.InvertVelocity(false, true)
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
func (s *PlayScene) Draw(screen *ebiten.Image) {

	screen.Fill(color.RGBA{0, 0, 0, 0xff})

	//for _, gameObject := range gamesObjects {
//		gameObject.Draw(screen)
//	}
}
