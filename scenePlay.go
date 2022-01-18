package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/wcscode/pong/engine"
)

var gamesObjects []*GameObject

var img *ebiten.Image

var dirX float64 = 1
var dirY float64 = 1

//var scene  engine.Scene
var paddle1 engine.GameObject
var paddle2 engine.GameObject
var ball engine.GameObject

func (s *engine.Scene) Init() {

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
	ball.Sprite.ImageWidth = 50
	ball.Sprite.ImageHeight = 50
	ball.Sprite.LoadAndCutImage(img, 100, 0)

	gamesObjects = append(gamesObjects, &ball)
}

func () Update() error {
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
func (engine.Scene) Draw(screen *ebiten.Image) {

	screen.Fill(color.RGBA{200, 0, 0, 0xff})
}
