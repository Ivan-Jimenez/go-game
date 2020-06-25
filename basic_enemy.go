package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	enemyHeight = 105
	enemyWidth  = 105
)

type basicEnemy struct {
	text *sdl.Texture
	x, y float64
}

func newBasicEnemy(r *sdl.Renderer, x, y float64) (be basicEnemy) {
	be.text = textureFromBMP(r, "sprites/basic_enemy.bmp")
	be.x = x
	be.y = y
	return be
}

func (be *basicEnemy) draw(r *sdl.Renderer) {
	x := be.x - enemyWidth/2.0
	y := be.y - enemyHeight/2.0

	r.CopyEx(
		be.text,
		&sdl.Rect{X: 0, Y: 0, W: enemyWidth, H: enemyHeight},
		&sdl.Rect{X: int32(x), Y: int32(y), W: enemyWidth, H: enemyHeight},
		180,
		&sdl.Point{X: enemyWidth / 2, Y: enemyHeight / 2},
		sdl.FLIP_NONE)
}
