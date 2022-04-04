package levels

import (
	"image/color"

	"github.com/veandco/go-sdl2/sdl"
)

type Object struct {
	Tex          *sdl.Texture
	X, Y         int
	ObjectWidth  int
	ObjectHeight int
}

type StartData struct {
	X, Y    int
	EndData struct {
		X, Y int
	}
}

var backgroundData Object
var PlayerStart StartData

func CreateLevel(Surf *sdl.Surface, Tex *sdl.Texture) (levelData []Object) {
	FormattedSurf, _ := Surf.ConvertFormat(sdl.PIXELFORMAT_RGB888, 0)

	type initialPos struct {
		X, Y  int
		found bool
	}

	var corner initialPos
	var pixCountX int
	var pixCountY int

	var colorBlack = color.RGBA{0, 0, 0, 255}

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
				levelData = append(levelData, Object{
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
	return levelData
}
