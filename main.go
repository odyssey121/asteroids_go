package main

import (
	"asteroids_go/assets"
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
) // Game implements ebiten.Game interface.
type Game struct {
	player *Player
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	// Write your game's logical update.
	g.player.Update()
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	// ebitenutil.DebugPrint(screen, "Hello, World!")
	// Write your game's rendering.
	g.player.Draw(screen)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	fmt.Println("assets.PlayerSprite: ", assets.PlayerSprite)

	game := &Game{}
	game.player = NewPlayer(game)
	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowTitle("Your game's title")
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
