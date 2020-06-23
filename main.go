package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 600
	screenHeight = 800
)

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

	render, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("[ERROR] initializing renderer: ", err)
	}
	defer render.Destroy()

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}

		render.SetDrawColor(255, 255, 255, 255)
		render.Clear()

		render.Present()
	}
}