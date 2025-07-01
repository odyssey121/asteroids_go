package main

import (
	"asteroids_go/assets"
	"fmt"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	rotationSpeedMin                    = -0.02
	ratationSpeedMax                    = 0.02
	numberOfSmallMeteorsFromLargeMeteor = 4
)

type Meteor struct {
	game          *GameScene
	position      Vector
	rotation      float64
	movement      Vector
	angle         float64
	rotationSpeed float64
	sprite        *ebiten.Image
}

func NewMeteor(baseVelosity float64, g *GameScene, index int) *Meteor {
	fmt.Println("index:", index)
	// target the center of the screen
	target := Vector{
		X: ScreenWidth / 2,
		Y: ScreenHeight / 2,
	}
	// pick a random angle.
	angle := rand.Float64() * 2 * math.Pi
	// the distance from the center that meteor should spawn at.
	r := ScreenWidth/2 + 500
	// create the position vector, using the angle and simple math
	pos := Vector{
		X: target.X + math.Cos(angle)*float64(r),
		Y: target.Y + math.Sin(angle)*float64(r),
	}
	// keep the meteor moving towards the center of the screen
	// random velocity
	velocity := baseVelosity + rand.Float64()*1.5

	// Create th direction vector and normalize it.
	direction := Vector{
		X: target.X - pos.X,
		Y: target.Y - pos.Y,
	}

	normolizedDirection := direction.Normalize()

	// create the movement vector
	movement := Vector{
		X: normolizedDirection.X * velocity,
		Y: normolizedDirection.Y * velocity,
	}

	// create a meteor object and return it
	sprite := assets.MeteorsSprites[rand.Intn(len(assets.MeteorsSprites))]
	m := &Meteor{
		game:          g,
		position:      pos,
		angle:         angle,
		movement:      movement,
		rotationSpeed: rotationSpeedMin + rand.Float64()*(ratationSpeedMax-rotationSpeedMin),
		sprite:        sprite,
	}
	return m

}

func (m *Meteor) Update() {
	dx := m.movement.X
	dy := m.movement.Y
	m.position.X += dx
	m.position.Y += dy
	m.rotation += m.rotationSpeed
	// keep meteor on screen
	m.keepOnScreen()

}

func (m *Meteor) Draw(screen *ebiten.Image) {
	bounds := m.sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Rotate(m.rotation)
	op.GeoM.Translate(halfW, halfH)

	op.GeoM.Translate(m.position.X, m.position.Y)
	screen.DrawImage(m.sprite, op)
}

func (m *Meteor) keepOnScreen() {
	if m.position.X >= float64(ScreenWidth) {
		m.position.X = 0
	}
	if m.position.X < 0 {
		m.position.X = ScreenWidth
	}
	if m.position.Y >= float64(ScreenHeight) {
		m.position.Y = 0
	}
	if m.position.Y < 0 {
		m.position.Y = ScreenHeight
	}
}
