package assets

import (
	"embed"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

//go:embed *
var assets embed.FS

var PlayerSprite = mustLoadImage("images/player.png")
var TitleFont = mustLoadFontFace("fonts/title.ttf")

func mustLoadFontFace(name string) *text.GoTextFaceSource {
	r, err := assets.Open(name)
	if err != nil {
		panic(err)
	}
	defer r.Close()

	ts, err := text.NewGoTextFaceSource(r)
	if err != nil {
		panic(err)
	}

	return ts

}

func mustLoadImage(imagePath string) *ebiten.Image {
	f, err := assets.Open(imagePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	image, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(image)

}
