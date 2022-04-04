package levels

import (
	"image/color"

	"github.com/veandco/go-sdl2/sdl"
)

func LevelOne(renderer *sdl.Renderer) (objectData []Object, LevelBG Object, PlayerStart StartData, err error) {

	Surf, _ := sdl.LoadBMP("LEVELS/LevelOneSprites/levelLayout.bmp")
	Tex, _ := renderer.CreateTextureFromSurface(Surf)
	FormattedSurf, _ := Surf.ConvertFormat(sdl.PIXELFORMAT_RGB888, 0)

	type initialPos struct {
		X, Y  int
		found bool
	}

	var corner initialPos
	var pixCountX int
	var pixCountY int

	var colorBlack = color.RGBA{0, 0, 0, 255}

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
				if corner.found != true {
					corner.X = X
					corner.Y = Y
					corner.found = true
				}
				for j, h := X, Y; FormattedSurf.At(j+1, Y) == colorBlack; {
					if FormattedSurf.At(j, h+1) == colorBlack {
						h++
					} else {
						pixCountY = h

						j++
					}
					pixCountX = j
					X = j
				}
				objectData = append(objectData, Object{
					Tex:          Tex,
					X:            corner.X,
					Y:            corner.Y,
					ObjectWidth:  pixCountX - corner.X,
					ObjectHeight: pixCountY - corner.Y,
				})
				corner.found = false
			}
		}
	}

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
