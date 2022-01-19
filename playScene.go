package main

import (
	"fmt"
	_ "image/png"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/wcscode/pong/engine"
)

type PlayScene engine.Scene
// Game implements ebiten.Game interface.

var img *ebiten.Image

//var scene  engine.Scene
var paddle1 engine.GameObject
var paddle2 engine.GameObject
var ball engine.GameObject


func (s *PlayScene) GetName() string {
	
	return s.Name
}


func (ps *PlayScene) GetActive() bool {
	
	return ps.Active
}

func (ps *PlayScene) SetActive(active bool) {
	
	ps.Active = active
}

//var gamesObjects []engine.GameObject
func (ps *PlayScene) Init() {
	

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

	ps.GamesObjects = append(ps.GamesObjects, &paddle1)

	paddle2.Sprite.Name = "Player 2"
	paddle2.PositionX = 270
	paddle2.PositionY = 45
	paddle2.Sprite.ImageWidth = 50
	paddle2.Sprite.ImageHeight = 145
	paddle2.Sprite.LoadAndCutImage(img, 50, 0)

	ps.GamesObjects = append(ps.GamesObjects, &paddle2)

	ball.PositionX = 320 * .5
	ball.PositionY = 240 * .5
	ball.VelocityX = 1
	ball.VelocityY = 1
	ball.Sprite.ImageWidth = 50
	ball.Sprite.ImageHeight = 50
	ball.Sprite.LoadAndCutImage(img, 100, 0)

	ps.GamesObjects = append(ps.GamesObjects, &ball)
}

func (ps *PlayScene) Update(keys []ebiten.Key) error {
   	
	keys = inpututil.AppendPressedKeys(keys[:0])
	

	ball.PositionX += ball.VelocityX
	ball.PositionY += ball.VelocityY

	if ball.PositionX >= engine.ScreenWidth {
	 	ball.InvertVelocity(true, false)
	}
	if ball.PositionX <= 0 {
		ball.InvertVelocity(true, false)
	}

	if ball.PositionY >= engine.ScreenHeight {
		ball.InvertVelocity(false, true)
	}
	if ball.PositionY <= 0 {
		ball.InvertVelocity(false, true)
	}

	for _, key := range keys {
		if key == ebiten.KeyS {
			paddle1.PositionY = math.Min(engine.ScreenWidth,  math.Min(engine.ScreenWidth, paddle1.PositionY + 4))
		}

		if key == ebiten.KeyW {
			paddle1.PositionY += -4
		}

		if key == ebiten.KeyArrowDown {
			fmt.Println(paddle2.PositionY)
			if(paddle2.PositionY <= engine.ScreenHeight - float64(paddle2.Sprite.ImageHeight)) {
				paddle2.PositionY +=  4
			}
		}

		if key == ebiten.KeyArrowUp {
			paddle2.PositionY += -4
		}
	}

	return nil
}
func (ps *PlayScene) Draw(screen *ebiten.Image) {  

	//screen.Fill(color.RGBA{0, 0, 0, 0xff}) 
	for _, gameObject := range ps.GamesObjects {
		gameObject.Draw(screen)
	}
}
