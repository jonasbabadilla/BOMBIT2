package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenHeight = 560
	screenWidth  = 960
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
	for {
		// event compiler
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent: // this is for if the event type (quit event) == true
				return
			}
		}
		renderer.SetDrawColor(255, 255, 255, 255) // color value = white at the moment
		renderer.Clear()

		renderer.Present() // Presents the elements that the renderer currently has

	}
}
