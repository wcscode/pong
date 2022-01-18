package main

import (
	"fmt"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/wcscode/pong/engine"
)

type MenuScene engine.Scene

func (s *MenuScene) Init() {

	var err error
	img, _, err = ebitenutil.NewImageFromFile("images/sprites.png")
	if err != nil {
		log.Fatal(err)
	}
}

func (s *MenuScene) Draw(screen *ebiten.Image) {

	//screen.Fill(color.RGBA{200, 0, 0, 0xff})
	fmt.Printf("Ok")
}


func (s *MenuScene) Update(g *Game) error {
	
	return nil
}


