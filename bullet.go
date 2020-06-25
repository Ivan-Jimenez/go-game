package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	bulletSize  = 32
	bulletSpeed = 0.15
)

type bullet struct {
	text  *sdl.Texture
	x, y  float64
	angle float64

	active bool
}

func newBullet(r *sdl.Renderer) (b bullet) {
	b.text = textureFromBMP(r, "sprites/player_bullet.bmp")
	return b
}

func (b *bullet) draw(r *sdl.Renderer) {
	if !b.active {
		return
	}

	x := b.x - bulletSize/2.0
	y := b.y - bulletSize/2.0

	r.Copy(
		b.text,
		&sdl.Rect{X: 0, Y: 0, W: bulletSize, H: bulletSize},
		&sdl.Rect{X: int32(x), Y: int32(y), W: bulletSize, H: bulletSize},
	)
}

func (b *bullet) update() {
	b.x += bulletSpeed * math.Cos(b.angle)
	b.y += bulletSpeed * math.Sin(b.angle)

	if b.x > screenWidth || b.x < 0 || b.y > screenHeight || b.y < 0 {
		b.active = false
	}
}

var bulletPool []*bullet

func initBulletPool(r *sdl.Renderer) {
	for i := 0; i < 30; i++ {
		b := newBullet(r)
		bulletPool = append(bulletPool, &b)
	}
}

func bulletFromPool() (*bullet, bool) {
	for _, b := range bulletPool {
		if !b.active {
			return b, true
		}
	}
	return nil, false
}