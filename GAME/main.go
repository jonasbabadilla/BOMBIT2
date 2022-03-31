package main

import (
	"fmt"

	levels "chaseGame/GAME/LEVELONE"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 1280
	screenHeight = 720
)

var ObjectData []levels.Object

var pChar player

func checkCollision() {
	//Check if player is on same Y level
	for i := range ObjectData {
		if pChar.y+float64(pChar.playerHeight*4) >= float64(ObjectData[i].Y) && pChar.y+float64(pChar.playerHeight*4) < float64(ObjectData[i].Y+(ObjectData[i].ObjectHeight/2)) {

			if pChar.x >= float64(ObjectData[i].X-16) && pChar.x <= float64(ObjectData[i].X+ObjectData[i].ObjectWidth-16) {
				pChar.y -= float64(gravity)
				JumpState = false
				JumpTimer = 0
				PlayerSpeedY = 7.00
				gravity = 5.00
			}
		}
	}
	if pChar.y+float64(pChar.playerHeight*4) >= screenHeight {
		pChar.x = float64(levels.PlayerStart["X"])
		pChar.y = float64(levels.PlayerStart["Y"])
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

	ObjectData = decideLevel()

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		renderer.SetDrawColor(144, 144, 144, 0)
		renderer.Clear()

		//levels.ObjectData[0].Draw(renderer, levels.ObjectData)

		pChar.Draw(renderer)
		pChar.Update()

		renderer.Present()

		checkCollision()

	}
}
