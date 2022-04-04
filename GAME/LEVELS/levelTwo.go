package levels

import (
	"github.com/veandco/go-sdl2/sdl"
)

func LevelTwo(renderer *sdl.Renderer) (levelData []Object, LevelBG Object, pStart StartData, err error) {

	Surf, _ := sdl.LoadBMP("LEVELS/LevelTwoSprites/levelLayout.bmp")
	Tex, _ := renderer.CreateTextureFromSurface(Surf)

	levelData = CreateLevel(Surf, Tex)

	defer Surf.Free()

	Surf, _ = sdl.LoadBMP("LEVELS/LevelTwoSprites/BG.bmp")
	BG, _ := renderer.CreateTextureFromSurface(Surf)

	backgroundData = Object{
		Tex:          BG,
		X:            0,
		Y:            0,
		ObjectWidth:  1280,
		ObjectHeight: 720,
	}

	defer Surf.Free()

	PlayerStart = StartData{X: 596, Y: 330, EndData: struct {
		X int
		Y int
	}{X: 1077, Y: 34}}

	return levelData, backgroundData, PlayerStart, nil

}
