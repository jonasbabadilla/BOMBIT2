package levels

import (
	"github.com/veandco/go-sdl2/sdl"
)

func LevelOne(renderer *sdl.Renderer) (objectData []Object, LevelBG Object, PlayerStart StartData, err error) {

	Surf, _ := sdl.LoadBMP("LEVELS/LevelOneSprites/levelLayout.bmp")
	Tex, _ := renderer.CreateTextureFromSurface(Surf)

	blockOne := Object{
		Tex:          Tex,
		X:            0,
		Y:            542,
		ObjectWidth:  388,
		ObjectHeight: 97,
	}
	blockTwo := Object{
		Tex:          Tex,
		X:            481,
		Y:            404,
		ObjectWidth:  347,
		ObjectHeight: 106,
	}
	blockThree := Object{
		Tex:          Tex,
		X:            987,
		Y:            267,
		ObjectWidth:  282,
		ObjectHeight: 106,
	}
	objectData = append(objectData, blockOne, blockTwo, blockThree)

	defer Surf.Free()

	Surf, _ = sdl.LoadBMP("LEVELS/LevelOneSprites/BG.bmp")
	BG, _ := renderer.CreateTextureFromSurface(Surf)

	backgroundData = Object{
		Tex:          BG,
		X:            0,
		Y:            0,
		ObjectWidth:  1280,
		ObjectHeight: 720,
	}

	defer Surf.Free()

	PlayerStart = StartData{X: 96, Y: 430, EndData: struct {
		X int
		Y int
	}{X: 1118, Y: 178}}

	return objectData, backgroundData, PlayerStart, nil

}
