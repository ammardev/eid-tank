package main

import (
	"math"

	"github.com/ammardev/eid-tank/internal/engine"
	"github.com/faiface/pixel"
)

var bullets []*Bullet

type Tank struct {
	sprite    *pixel.Sprite
	vector    pixel.Vec
	direction float64
}

func (t *Tank) Accelerate(delta float64) {
	t.vector.Y += math.Sin(t.direction) * delta * 600
	t.vector.X += math.Cos(t.direction) * delta * 600
}

func (t *Tank) Reverse(delta float64) {
	t.vector.Y -= math.Sin(t.direction) * delta * 300
	t.vector.X -= math.Cos(t.direction) * delta * 300
}

func (t *Tank) SteerRight(delta float64) {
	t.direction -= 3 * math.Pi / 2 * delta
}

func (t *Tank) SteerLeft(delta float64) {
	t.direction += 3 * math.Pi / 2 * delta
}

func (t *Tank) Shoot() {
	bullet := t.createBullet()
	bullets = append(bullets, &bullet)
}

func (t *Tank) Draw() {
	matrix := pixel.
		IM.
		Scaled(pixel.ZV, 0.5).
		Rotated(pixel.ZV, t.direction+math.Pi/2).
		Moved(game.Window.Bounds().Center().Add(t.vector))

	t.sprite.Draw(game.Window, matrix)
}

type Bullet struct {
	sprite    *pixel.Sprite
	vector    pixel.Vec
	direction float64
	tank      *Tank
	ttl       byte
}

func (b *Bullet) Move(delta float64) {
	b.vector.Y += math.Sin(b.direction) * delta * 600
	b.vector.X += math.Cos(b.direction) * delta * 600

	if b.ttl <= 0 {
		bullets[0] = nil
		bullets = bullets[1:]
	}

	b.ttl--
}

func (b *Bullet) Draw() {
	matrix := pixel.
		IM.
		Scaled(pixel.ZV, 0.5).
		Rotated(pixel.ZV, b.direction+3*math.Pi/2).
		Moved(game.Window.Bounds().Center().Add(b.vector))

	b.sprite.Draw(game.Window, matrix)
}

func createTank() Tank {
	return Tank{
		sprite:    engine.CreateSprite("assets/tank_blue.png"),
		direction: math.Pi / 2,
	}
}

func (t *Tank) createBullet() Bullet {
	return Bullet{
		sprite:    engine.CreateSprite("assets/bullet_blue.png"),
		vector:    t.vector,
		direction: t.direction,
		ttl:       50,
	}
}
