package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type GameScene struct {
	player *Player
}

func NewGameScene() *GameScene {
	gc := &GameScene{}
	gc.player = NewPlayer(gc)
	return gc
}

func (g *GameScene) Update() error {
	g.player.Update()
	return nil
}

func (g *GameScene) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)
}

func (g *GameScene) Layout(outsideWidth, outsideHeight int) (ScreenWidth, ScreenHight int) {
	return outsideWidth, outsideHeight
}
