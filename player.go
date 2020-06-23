package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type player struct {
	text *sdl.Texture
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

	return p, nil
}

func (p *player) draw(r *sdl.Renderer) {
	r.Copy(
		p.text,
		&sdl.Rect{X: 0, Y: 0, W: 105, H: 105},
		&sdl.Rect{X: 0, Y: 0, W: 105, H: 105},
	)
}
