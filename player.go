package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerSpeed  = 0.05
	playerHeight = 105
	playerWidth  = 105
)

type player struct {
	text *sdl.Texture
	x, y float64
}

func newPlayer(r *sdl.Renderer) (p player, err error) {
	img, err := sdl.LoadBMP("sprites/player.bmp")
	if err != nil {
		return player{}, fmt.Errorf("[ERROR] loading player sprite: %v", err)
	}
	defer img.Free()
	p.text, err = r.CreateTextureFromSurface(img)
	if err != nil {
		return player{}, fmt.Errorf("[ERROR] creating player texture: %v", err)
	}

	p.x = screenWidth / 2
	p.y = screenHeight - playerHeight/2.0

	return p, nil
}

func (p *player) draw(r *sdl.Renderer) {
	x := p.x - playerWidth/2.0
	y := p.y - playerHeight/2.0

	r.Copy(
		p.text,
		&sdl.Rect{X: 0, Y: 0, W: 105, H: 105},
		&sdl.Rect{X: int32(x), Y: int32(y), W: 105, H: 105},
	)
}

func (p *player) update() {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_LEFT] == 1 {
		p.x -= playerSpeed
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		p.x += playerSpeed
	}
}
