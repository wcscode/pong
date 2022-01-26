package main

import (
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/wcscode/pong/engine"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type MenuScene engine.Scene

var faceMenu font.Face
var alphaMenu uint8
var countMenu int

func (s *MenuScene) Init() {

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
	faceMenu, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatalf("New face error: %v", err)
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

func (m *MenuScene) Update(keys []ebiten.Key) error {

	alphaMenu = uint8(255 * ((countMenu / 30) % 2))
	countMenu++

	if countMenu == 60 {
		countMenu = 0
	}

	keys = inpututil.AppendPressedKeys(keys[:0])

	for _, key := range keys {

		if key == ebiten.KeyEnter {

			engine.SetActiveScene("Play")
		}
	}

	return nil
}

func (ms *MenuScene) Draw(screen *ebiten.Image) {

	//screen.Fill(color.RGBA{200, 0, 0, 0xff})
	text.Draw(screen, "Press ENTER", faceMenu, 90, 120, color.RGBA{255, 255, 255, alphaMenu})

}
