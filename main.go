package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 600
	screenHeight = 800
)

func textureFromBMP(r *sdl.Renderer, filename string) *sdl.Texture {
	img, err := sdl.LoadBMP(filename)
	if err != nil {
		panic(fmt.Errorf("[ERROR] loading %v %v", filename, err))
	}
	defer img.Free()
	text, err := r.CreateTextureFromSurface(img)
	if err != nil {
		panic(fmt.Errorf("[ERROR] creating texture from %v: %v", filename, err))
	}

	return text
}

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("[ERROR] initializing SLD: ", err)
	}

	window, err := sdl.CreateWindow(
		"Go Game",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight,
		sdl.WINDOW_OPENGL,
	)
	if err != nil {
		fmt.Println("[ERROR] initializing window: ", err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("[ERROR] initializing renderer: ", err)
	}
	defer renderer.Destroy()

	player := newPlayer(renderer)
	if err != nil {
		fmt.Println("[ERROR] creating player: ", err)
		return
	}

	var enemies []basicEnemy

	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			x := (float64(i)/5)*screenWidth + (enemyWidth / 2.0)
			y := float64(j)*enemyWidth + (enemyHeight / 2.0)

			enemy := newBasicEnemy(renderer, x, y)
			if err != nil {
				fmt.Println("[ERROR] creating basic enemy: ", err)
			}

			enemies = append(enemies, enemy)
		}
	}

	initBulletPool(renderer)

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}

		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

		player.draw(renderer)
		player.update()

		for _, enemy := range enemies {
			enemy.draw(renderer)
		}

		for _, b := range bulletPool {
			b.draw(renderer)
			b.update()
		}

		renderer.Present()
	}
}
