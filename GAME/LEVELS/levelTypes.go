package levels

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Object struct {
	Tex          *sdl.Texture
	X, Y         int
	ObjectWidth  int
	ObjectHeight int
}

var objectData []Object

type StartData struct {
	X, Y    int
	EndData struct {
		X, Y int
	}
}

var backgroundData Object
var PlayerStart StartData
