package levels

import (
	"fmt"
	"image/color"

	"github.com/veandco/go-sdl2/sdl"
)

var byteCount int

func LevelOne(renderer *sdl.Renderer) (objectData []Object, LevelBG Object, PlayerStart StartData, err error) {

	Surf, _ := sdl.LoadBMP("LEVELS/LevelOneSprites/levelLayout.bmp")
	Tex, _ := renderer.CreateTextureFromSurface(Surf)
	FormattedSurf, _ := Surf.ConvertFormat(sdl.PIXELFORMAT_RGB888, 0)
	var colorBlack color.RGBA
	var Count int

	colorBlack = color.RGBA{0, 0, 0, 255}

	//THIS FINALLY WORKS
	//IT SEARCHES THE LEVELLAYOUT.BMP FOR BLACK PIXELS
	//DOESNT DO ANYTHING BUT PRINT THE X AND Y AND COUNT HOW MANY BLACK PXIELS IN TOTAL
	//BUT LATER IT WILL AUTOMATICALLY TAKE THE WIDTH AND HEIGHT OF EVERY BLACK PLATFORM
	//AND ADD IT TO THE LEVEL'S OBJECT DATA
	//THAT WAY PLATFORMS GET ADDED AUTOMATICALLY
	//NO MORE MANUALLY ADDING PLATFORMS!!!
	for Y := 0; Y < FormattedSurf.Bounds().Dy(); Y++ {
		for X := 0; X < FormattedSurf.Bounds().Dx(); X++ {
			if FormattedSurf.At(X, Y) == colorBlack {
				Count++
				fmt.Println(X+1, Y+1, Count)
			}
		}
	}
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

	PlayerStart = StartData{X: 96, Y: 430, EndData: struct {
		X int
		Y int
	}{X: 1118, Y: 178}}

	return objectData, backgroundData, PlayerStart, nil

}
