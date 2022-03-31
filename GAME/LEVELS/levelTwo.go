package levels

import (
	"github.com/veandco/go-sdl2/sdl"
)

func LevelTwo(renderer *sdl.Renderer) (objectData []Object, LevelBG Object, pStart StartData, err error) {

	Surf, _ := sdl.LoadBMP("LEVELS/LevelTwoSprites/levelLayout.bmp")
	Tex, _ := renderer.CreateTextureFromSurface(Surf)

	objectData = append(objectData,
		Object{
			Tex:          Tex,
			X:            483,
			Y:            476,
			ObjectWidth:  270,
			ObjectHeight: 77,
		},
		Object{
			Tex:          Tex,
			X:            225,
			Y:            679,
			ObjectWidth:  347,
			ObjectHeight: 87,
		},
		Object{
			Tex:          Tex,
			X:            8,
			Y:            583,
			ObjectWidth:  128,
			ObjectHeight: 81,
		},
		Object{
			Tex:          Tex,
			X:            8,
			Y:            409,
			ObjectWidth:  128,
			ObjectHeight: 81,
		},
		Object{
			Tex:          Tex,
			X:            8,
			Y:            239,
			ObjectWidth:  128,
			ObjectHeight: 81,
		},
		Object{
			Tex:          Tex,
			X:            195,
			Y:            179,
			ObjectWidth:  240,
			ObjectHeight: 51,
		},
		Object{
			Tex:          Tex,
			X:            561,
			Y:            190,
			ObjectWidth:  218,
			ObjectHeight: 51,
		},
		Object{
			Tex:          Tex,
			X:            960,
			Y:            180,
			ObjectWidth:  288,
			ObjectHeight: 82,
		},
	)

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

	return objectData, backgroundData, PlayerStart, nil

}
