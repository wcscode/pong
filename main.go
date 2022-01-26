package main

import (
	_ "image/png"
	"log"

	//"github.com/hajimehoshi/ebiten/ebitenutil"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/wcscode/pong/engine"
)

//var menuScene MenuScene // = MenuScene{}
// Game implements ebiten.Game interface.
type Game struct {
	keys []ebiten.Key
}

//var rect = image.Rectangle{}

func init() {

	menuScene := MenuScene{Name: "Menu", Active: true, GamesObjects: []*engine.GameObject{}}
	playScene := PlayScene{Name: "Play", GamesObjects: []*engine.GameObject{}}

	menuScene.Init()
	playScene.Init()

	engine.SetActiveScene("Play")

	engine.ScenesBehaviors = append(engine.ScenesBehaviors, &menuScene, &playScene)

	//x = image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{50, 50}})
	//rect := image.Rect(255, 255, 255, 255) //  geometry of 2nd rectangle which we draw atop above rectangle
	//myred := color.RGBA{200, 0, 0, 255}

}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	// Write your game's logical update.

	for _, scene := range engine.ScenesBehaviors {
		if scene.GetActive() {
			scene.Update(g.keys)
		}
	}

	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {

	// create a red rectangle atop the green surface
	//draw.Draw(screen, red_rect.Bounds(), &image.Uniform{myred}, image.ZP, draw.Src)

	for _, scene := range engine.ScenesBehaviors {
		if scene.GetActive() {
			scene.Draw(screen)
		}
	}
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	game := &Game{}

	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(int(engine.ScreenWidth), int(engine.ScreenHeight))
	ebiten.SetWindowTitle("My Pong version!!!")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
