package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
) // Game implements ebiten.Game interface.

func main() {

	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowTitle("Your game's title")
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(NewGameScene()); err != nil {
		log.Fatal(err)
	}
}
