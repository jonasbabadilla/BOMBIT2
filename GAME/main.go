package main

import (
	"fmt"

	levels "chaseGame/GAME/LEVELS"

	"github.com/veandco/go-sdl2/sdl"
)

var Renderer *sdl.Renderer

const (
	screenWidth  = 1280
	screenHeight = 720
)

var ObjectData []levels.Object
var backgroundData levels.Object

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
		pChar.x = float64(pStart.X)
		pChar.y = float64(pStart.Y)
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
	Renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("initializing Renderer:", err)
		return
	}
	defer Renderer.Destroy()

	ObjectData, backgroundData, pStart = decideLevel()

	pChar, err = NewPlayer(Renderer)
	if err != nil {
		fmt.Println("creating player:", err)
		return
	}

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		Renderer.SetDrawColor(144, 144, 144, 0)
		Renderer.Clear()

		//levels.ObjectData[0].Draw(Renderer, levels.ObjectData)
		Draw(Renderer, ObjectData)

		pChar.Draw(Renderer)
		pChar.Update()

		Renderer.Present()

		checkCollision()

	}
}

func Draw(Renderer *sdl.Renderer, ObjectData []levels.Object) {

	backgroundData.Tex.SetBlendMode(sdl.BLENDMODE_BLEND)
	Renderer.Copy(backgroundData.Tex, nil, nil)

	for _, k := range ObjectData {
		Renderer.Copy(k.Tex,
			&sdl.Rect{X: int32(k.X), Y: int32(k.Y), W: int32(k.ObjectWidth), H: int32(k.ObjectHeight)},
			&sdl.Rect{X: int32(k.X), Y: int32(k.Y), W: int32(k.ObjectWidth), H: int32(k.ObjectHeight)},
		)
	}

}
