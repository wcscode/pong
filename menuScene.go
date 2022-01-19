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

type MenuScene engine.Scene

func (s *MenuScene) Init() {	

	var err error
	img, _, err = ebitenutil.NewImageFromFile("images/sprites.png")
	if err != nil {
		log.Fatal(err)
	}
}


func (ms *MenuScene) GetName() string {
	
	return ms.Name
}

func (ms *MenuScene) GetActive() bool {
	
	return ms.Active
}
func (ms *MenuScene) SetActive(active bool) {
	
	ms.Active = active
}

func (ms *MenuScene) Draw(screen *ebiten.Image) {
	
	screen.Fill(color.RGBA{200, 0, 0, 0xff})
	
}

func (m *MenuScene) Update(keys []ebiten.Key) error {	

	keys = inpututil.AppendPressedKeys(keys[:0])

	for _, key := range keys {

		if(key == ebiten.KeyEnter){

			engine.SetActiveScene("Play")
		}
	}

	return nil
}


