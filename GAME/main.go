package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 1280
	screenHeight = 720
)

var pChar player
var levelObject object

func checkCollision() {
	if pChar.y+float64(pChar.playerHeight*2) == float64(levelObject.y) {
		if pChar.x <= float64(levelObject.objectWidth+(levelObject.x)) && pChar.x >= float64(levelObject.x-(levelObject.objectWidth)) {
			pChar.y -= float64(gravity)
			JumpState = false
			JumpTimer = 0
		}
	}

}

func main() {
	if err := sdl.Init(uint32(sdl.INIT_EVERYTHING)); err != nil {
		fmt.Println("initializing sdl:", err)
		return
	}
	window, err := sdl.CreateWindow(
		"GAMING WITH SDL",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight,
		sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Println("initializing window:", err)
		return
	}
	defer window.Destroy()
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("initializing renderer:", err)
		return
	}
	defer renderer.Destroy()
	pChar, err = NewPlayer(renderer)
	if err != nil {
		fmt.Println("creating player:", err)
		return
	}
	levelObject, err = NewObject(renderer)
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		renderer.SetDrawColor(119, 23, 23, 0)
		renderer.Clear()
		pChar.Draw(renderer)
		pChar.Update()

		levelObject.Draw(renderer)
		levelObject.Update()

		renderer.Present()
		checkCollision()

	}
}
