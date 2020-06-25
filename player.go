package main

import (
	"math"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerSpeed  = 0.05
	playerHeight = 105
	playerWidth  = 105

	playerShootCooldown = time.Millisecond * 250
)

type player struct {
	text *sdl.Texture
	x, y float64

	lastShoot time.Time
}

func newPlayer(r *sdl.Renderer) (p player) {
	p.text = textureFromBMP(r, "sprites/player.bmp")
	p.x = screenWidth / 2
	p.y = screenHeight - playerHeight/2.0
	return p
}

func (p *player) draw(r *sdl.Renderer) {
	x := p.x - playerWidth/2.0
	y := p.y - playerHeight/2.0

	r.Copy(
		p.text,
		&sdl.Rect{X: 0, Y: 0, W: playerWidth, H: playerHeight},
		&sdl.Rect{X: int32(x), Y: int32(y), W: playerWidth, H: playerHeight},
	)
}

func (p *player) update() {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_LEFT] == 1 {
		if p.x-(playerWidth/2.0) > 0 {
			p.x -= playerSpeed
		}
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		if p.x+(playerWidth/2.0) < screenWidth {
			p.x += playerSpeed
		}
	}

	if keys[sdl.SCANCODE_SPACE] == 1 {
		if time.Since(p.lastShoot) < playerShootCooldown {
			return
		}
		p.shoot(p.x+25, p.y-20)
		p.shoot(p.x-25, p.y-20)
		p.lastShoot = time.Now()
	}
}

func (p *player) shoot(x, y float64) {
	if b, ok := bulletFromPool(); ok {
		b.active = true
		b.x = x
		b.y = y
		b.angle = 270 * (math.Pi / 180)
	}
}
