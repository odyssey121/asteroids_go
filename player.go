package main

import (
	"asteroids_go/assets"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	rotationPerSecond = math.Pi
	maxAcceleration   = 8.0
	ScreenWidth       = 1280
	ScreenHeight      = 768
)

type Player struct {
	game           *Game
	sprite         *ebiten.Image
	rotation       float64
	position       Vector
	playerVelocity float64
}

func NewPlayer(game *Game) *Player {
	sprite := assets.PlayerSprite
	//centered
	bounds := sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfY := float64(bounds.Dy()) / 2

	pos := Vector{
		X: ScreenWidth/2 - halfW,
		Y: ScreenHeight/2 - halfY,
	}
	player := &Player{sprite: assets.PlayerSprite, position: pos}
	return player
}

func (p *Player) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	bounds := p.sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	opt.GeoM.Translate(-halfW, -halfH)
	opt.GeoM.Rotate(float64(p.rotation))
	opt.GeoM.Translate(halfW, halfH)

	opt.GeoM.Translate(p.position.X, p.position.Y)

	screen.DrawImage(p.sprite, opt)
}

func (p *Player) Update() {
	speed := rotationPerSecond / float64(ebiten.TPS())
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.rotation -= float64(speed)
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.rotation += float64(speed)
	}

	p.accelerate()
	p.keepOnScreen()
}

var curAcceleration float64

func (p *Player) accelerate() {
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		if curAcceleration < maxAcceleration {
			curAcceleration = p.playerVelocity + 4
		}
		if curAcceleration >= 8 {
			curAcceleration = 8
		}

		p.playerVelocity = curAcceleration
		// Move in the direction we are pointing.
		dx := math.Sin(p.rotation) * curAcceleration
		dy := math.Cos(p.rotation) * -curAcceleration
		// Move player on the screen
		p.position.X += dx
		p.position.Y += dy
	}

}

func (p *Player) keepOnScreen() {
	if p.position.X > float64(ScreenWidth) {
		p.position.X = 0
	}
	if p.position.X < 0 {
		p.position.X = float64(ScreenWidth)
	}
	if p.position.Y > float64(ScreenHeight) {
		p.position.Y = 0
	}
	if p.position.Y < 0 {
		p.position.Y = float64(ScreenHeight)
	}

}
