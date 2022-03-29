package MainGame

import (
	"fmt"

	"chaseGame/GAME/StoryGame"

	"github.com/veandco/go-sdl2/sdl"
)

var ScreenWidth int32
var ScreenHeight int32

func main() {

	StoryGame.Resolution = StoryGame.Dimensions{1280, 720}

	if err := sdl.Init(uint32(sdl.INIT_EVERYTHING)); err != nil {
		fmt.Println("initializing sdl:", err)
		return
	}

	window, err := sdl.CreateWindow(
		"GAMING WITH SDL",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		ScreenWidth, ScreenHeight,
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

	plr, err := StoryGame.NewPlayer(renderer)
	if err != nil {
		fmt.Println("creating player:", err)
		return
	}

	bar, err := StoryGame.NewObject(renderer)

	for {

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}

		renderer.SetDrawColor(119, 23, 23, 0)
		renderer.Clear()

		plr.Draw(renderer)
		plr.Update()

		bar.Draw(renderer)
		bar.Update()

		renderer.Present()

	}
}
