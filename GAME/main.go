package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenHeight = 720
	screenWidth  = 1280
)

func main() {
	if err := sdl.Init(uint32(sdl.INIT_EVERYTHING)); err != nil {
		fmt.Println("initializing sdl:", err)
		return
	}

	// creates the window
	window, err := sdl.CreateWindow(
		"GAMING WITH SDL",                                // title of window
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, // where the window is placed on screen
		screenWidth, screenHeight, // paremeters
		sdl.WINDOW_OPENGL) // flag

	if err != nil {
		fmt.Println("initializing window:", err)
		return
	}

	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED) //
	if err != nil {
		fmt.Println("initializing renderer:", err)
		return
	}

	defer renderer.Destroy()

	if err != nil {
		fmt.Println("loading texture:", err)
		return
	}

	char, err := sdl.LoadBMP("SPRITES/CHAR.bmp")
	if err != nil {
		fmt.Println("loading texture:", err)
		return
	}
	playerchar, err := renderer.CreateTextureFromSurface(char)
	if err != nil {
		fmt.Println("loading texture:", err)
		return
	}

	for {

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		renderer.SetDrawColor(84, 84, 84, 255) // color value = white at the moment
		renderer.Clear()

		renderer.Copy(playerchar,
			&sdl.Rect{X: 0, Y: 0, W: 64, H: 64},
			&sdl.Rect{X: 640, Y: 360, W: 64, H: 64})

		renderer.Present()

	}
}
