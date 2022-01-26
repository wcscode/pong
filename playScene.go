package main

import (
	"fmt"
	"image/color"
	_ "image/png"
	"log"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/wcscode/pong/engine"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type PlayScene engine.Scene
type Score struct {
	Player1, Player2 int
}

// Game implements ebiten.Game interface.

var img *ebiten.Image

//var scene  engine.Scene
var paddle1 engine.GameObject
var paddle2 engine.GameObject
var ball engine.GameObject
var faceScore font.Face
var score Score

//var count int

func ResetBall(ball *engine.GameObject) {

	ball.PositionX = 320 * .5
	ball.PositionY = 240 * .5

	ballDir := [2]float64{-1, 1}
	ball.VelocityX = math.Min((math.Abs(ball.VelocityX)+.2)*ballDir[rand.Int31n(2)], 20)
	ball.VelocityY = math.Min((math.Abs(ball.VelocityY)+.2)*ballDir[rand.Int31n(2)], 20)
	//fmt.Print(ballDir[rand.Int31n(0)])
}

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

	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatalf("Parse error: %v", err)
	}

	const dpi = 72
	faceScore, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatalf("New face error: %v", err)
	}

	paddle1.Sprite.Name = "Player 1"
	paddle1.PositionX = 0
	paddle1.PositionY = 45
	paddle1.Sprite.ImageWidth = 50
	paddle1.Sprite.ImageHeight = 145
	paddle1.BoxCollision.X0 = 0
	paddle1.BoxCollision.Y0 = 0
	paddle1.BoxCollision.X1 = 10
	paddle1.BoxCollision.Y1 = 115
	paddle1.Sprite.LoadAndCutImage(img, 0, 0)

	ps.GamesObjects = append(ps.GamesObjects, &paddle1)

	paddle2.Sprite.Name = "Player 2"
	paddle2.PositionX = 270
	paddle2.PositionY = 45
	paddle2.Sprite.ImageWidth = 50
	paddle2.Sprite.ImageHeight = 145
	paddle2.BoxCollision.X0 = 0
	paddle2.BoxCollision.Y0 = 0
	paddle2.BoxCollision.X1 = 10
	paddle2.BoxCollision.Y1 = 115
	paddle2.Sprite.LoadAndCutImage(img, 50, 0)

	ps.GamesObjects = append(ps.GamesObjects, &paddle2)

	ballDir := [2]float64{-1, 1}

	ball.PositionX = 320 * .5
	ball.PositionY = 240 * .5
	ball.VelocityX = 3 * ballDir[rand.Int31n(1)]
	ball.VelocityY = 3 * ballDir[rand.Int31n(1)]

	ball.Sprite.ImageWidth = 50
	ball.Sprite.ImageHeight = 50
	ball.BoxCollision.X0 = 0
	ball.BoxCollision.Y0 = 0
	ball.BoxCollision.X1 = 13
	ball.BoxCollision.Y1 = 13
	ball.Sprite.LoadAndCutImage(img, 100, 0)

	ps.GamesObjects = append(ps.GamesObjects, &ball)

	score = Score{Player1: 0, Player2: 0}
}

func (ps *PlayScene) Update(keys []ebiten.Key) error {

	//if ball.PositionX >= 0 {
	if engine.IsColliding(&ball, &paddle2) {
		ball.InvertVelocity(true, false)
	}
	//} else {
	if engine.IsColliding(&ball, &paddle1) {
		ball.InvertVelocity(true, false)
	}
	//}

	keys = inpututil.AppendPressedKeys(keys[:0])

	ball.PositionX += ball.VelocityX
	ball.PositionY += ball.VelocityY

	if ball.PositionX > 285 {
		score.Player1 += 1
		ResetBall(&ball)
		ball.AddVelocity(1, 1)
	}

	if ball.PositionX < -15 {
		score.Player2 += 1
		ResetBall(&ball)
		ball.AddVelocity(1, 1)
	}

	if ball.PositionY > 205 {
		ball.InvertVelocity(false, true)
		ball.VelocityX += rand.NormFloat64() * -1
		ball.VelocityY += rand.NormFloat64() * -1
	}

	if ball.PositionY < -15 {
		ball.InvertVelocity(false, true)
		ball.VelocityX += rand.NormFloat64()
		ball.VelocityY += rand.NormFloat64()
	}

	for _, key := range keys {

		if key == ebiten.KeyEscape {
			engine.SetActiveScene("Menu")
			ps.Init()
		}

		if key == ebiten.KeyS {
			if paddle1.PositionY < 109 {
				paddle1.PositionY += 4
			}
		}

		if key == ebiten.KeyW {
			if paddle1.PositionY > -20 {
				paddle1.PositionY += -4
			}
		}

		if key == ebiten.KeyArrowDown {

			if paddle2.PositionY < 109 {
				paddle2.PositionY += 4
			}
		}

		if key == ebiten.KeyArrowUp {

			if paddle2.PositionY > -20 {
				paddle2.PositionY += -4
			}
		}
	}

	//count++

	return nil
}
func (ps *PlayScene) Draw(screen *ebiten.Image) {

	//text.Draw(screen, fmt.Sprint((count/60)%5), face, 120, 70, color.RGBA{255, 255, 255, 255})

	text.Draw(screen, fmt.Sprint(score.Player1), faceScore, 120, 30, color.RGBA{255, 255, 255, 255})
	text.Draw(screen, fmt.Sprint(score.Player2), faceScore, 180, 30, color.RGBA{255, 255, 255, 255})
	//screen.Fill(color.RGBA{0, 0, 0, 0xff})
	for _, gameObject := range ps.GamesObjects {
		gameObject.Draw(screen)
	}
}
