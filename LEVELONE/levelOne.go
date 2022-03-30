package levelOne

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Object struct {
	Tex          *sdl.Texture
	X, Y         int
	ObjectWidth  int
	ObjectHeight int
}

var tempObj *sdl.Texture

var objectData []Object

var backgroundData Object
var PlayerStart map[string]int

func NewObject(renderer *sdl.Renderer) (objectData []Object, err error) {

	Surf, _ := sdl.LoadBMP("SPRITES/LEVELONE/levelOneLayout.bmp")
	Tex, _ := renderer.CreateTextureFromSurface(Surf)

	PlayerStart = make(map[string]int, 2)
	PlayerStart["X"] = 96
	PlayerStart["Y"] = 430

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

	Surf, _ = sdl.LoadBMP("SPRITES/LEVELONE/BG.bmp")
	BG, _ := renderer.CreateTextureFromSurface(Surf)

	backgroundData = Object{
		Tex:          BG,
		X:            0,
		Y:            0,
		ObjectWidth:  1280,
		ObjectHeight: 720,
	}

	defer Surf.Free()

	return objectData, nil

}

func (o Object) Draw(renderer *sdl.Renderer, objData []Object) {

	backgroundData.Tex.SetBlendMode(sdl.BLENDMODE_BLEND)
	renderer.Copy(backgroundData.Tex, nil, nil)

	for _, k := range objData {
		renderer.Copy(k.Tex,
			&sdl.Rect{X: int32(k.X), Y: int32(k.Y), W: int32(k.ObjectWidth), H: int32(k.ObjectHeight)},
			&sdl.Rect{X: int32(k.X), Y: int32(k.Y), W: int32(k.ObjectWidth), H: int32(k.ObjectHeight)},
		)
	}

}

func (o Object) Update() {

}
