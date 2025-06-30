package main

import (
	"asteroids_go/assets"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font"
)

type TitleScene struct{}

func (t *TitleScene) Draw(screen *ebiten.Image) {
	drawTxt := "press space to start"
	op := &text.DrawOptions{LayoutOptions: text.LayoutOptions{PrimaryAlign: text.AlignCenter}}
	op.ColorScale.ScaleWithColor(color.White)
	op.GeoM.Translate(float64(ScreenWidth/2), ScreenHeight-200)
	text.Draw(screen, drawTxt, &text.GoTextFace{Source: assets.TitleFont, Size: 48}, op)
	// tw := widthOfText(assets.TitleFont, drawTxt)
	// text.Draw(screen, drawTxt, assets.TitleFont, ScreenWidth/2-tw/2, ScreenHeight-200, color.White)
}

func (t *TitleScene) Update(state *State) error {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		state.SceneManager.GoToScene(NewGameScene())
		return nil
	}
	return nil
}

func widthOfText(f font.Face, t string) int {
	_, textWidth := font.BoundString(f, t)
	return textWidth.Round()
}
