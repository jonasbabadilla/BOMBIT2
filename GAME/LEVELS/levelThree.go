package levels

import (
	"github.com/veandco/go-sdl2/sdl"
)

func LevelThree(renderer *sdl.Renderer) (objectData []Object, LevelBG Object, PlayerStart StartData, err error) {

	Surf, _ := sdl.LoadBMP("LEVELS/LevelThreeSprites/levelLayout.bmp")
	Tex, _ := renderer.CreateTextureFromSurface(Surf)

	/*
		ScanTex, _ := renderer.CreateTexture(Surf.Format.Format, sdl.TEXTUREACCESS_STREAMING, 1280, 720)

		ScanTex.Lock(&sdl.Rect{W: 1280, H:720, X: 0, Y: 0})

		PixelFormat := sdl.MapRGB(Surf.Format, 0, 0, 0)

		if err := Surf.SetColorKey(true, PixelFormat); err == nil {
		}
	*/

	blockOne := Object{
		Tex:          Tex,
		X:            6,
		Y:            189,
		ObjectWidth:  147,
		ObjectHeight: 56,
	}
	blockTwo := Object{
		Tex:          Tex,
		X:            166,
		Y:            341,
		ObjectWidth:  189,
		ObjectHeight: 56,
	}
	blockThree := Object{
		Tex:          Tex,
		X:            645,
		Y:            418,
		ObjectWidth:  155,
		ObjectHeight: 50,
	}
	blockFour := Object{
		Tex:          Tex,
		X:            938,
		Y:            322,
		ObjectWidth:  200,
		ObjectHeight: 54,
	}

	objectData = append(objectData, blockOne, blockTwo, blockThree, blockFour)

	defer Surf.Free()

	Surf, _ = sdl.LoadBMP("LEVELS/LevelThreeSprites/BG.bmp")
	BG, _ := renderer.CreateTextureFromSurface(Surf)

	backgroundData = Object{
		Tex:          BG,
		X:            0,
		Y:            0,
		ObjectWidth:  1280,
		ObjectHeight: 720,
	}

	defer Surf.Free()

	PlayerStart = StartData{X: 36, Y: 129, EndData: struct {
		X int
		Y int
	}{X: 1006, Y: 273}}

	return objectData, backgroundData, PlayerStart, nil

}
