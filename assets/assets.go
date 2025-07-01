package assets

import (
	"embed"
	"image"
	_ "image/png"
	"io/fs"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

//go:embed *
var assets embed.FS

var PlayerSprite = mustLoadImage("images/player.png")
var TitleFont = mustLoadFontFace("fonts/title.ttf")
var MeteorsSprites = mustLoadImages("images/meteors/*.png")
var MeteorsSpritesSmall = mustLoadImages("images/meteors-small/*.png")

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

func mustLoadImages(path string) []*ebiten.Image {
	matches, err := fs.Glob(assets, path)
	if err != nil {
		panic(err)
	}

	images := make([]*ebiten.Image, len(matches))

	for i, p := range matches {
		images[i] = mustLoadImage(p)
	}

	return images

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
