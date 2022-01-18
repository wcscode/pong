package main

import (
	_ "image/png"
	"log"

	//"github.com/hajimehoshi/ebiten/ebitenutil"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/wcscode/pong/engine"
)

//var scenes []engine.Scene
var scenes []engine.DrawerUpdater
//var scenes []engine.Scene
var menuScene MenuScene
var gamesObjects []engine.GameObject

// Game implements ebiten.Game interface.
type Game struct {
	keys []ebiten.Key
}

func init() {

	 menuScene := MenuScene{Name: "Menu", Active: true}	 
	// playScene := PlayScene{Name: "Play", Active: false}	 
	 
	scenes = []engine.DrawerUpdater{}
	 //scenes = append(scenes, engine.Scene(menuScene), engine.Scene(playScene))
	 scenes = append(scenes, *menuScene)
	 
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	// Write your game's logical update.

	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	
	for _, scene := range scenes {		
		/*&if(scene.Active) {
            
		}*/
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
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Your game's title")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
