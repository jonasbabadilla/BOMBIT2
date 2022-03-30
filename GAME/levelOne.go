package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type object struct {
	Tex          *sdl.Texture
	x, y         int
	objectWidth  int
	objectHeight int
}

var tempObj *sdl.Texture

var objectData []object

var backgroundData object
var PlayerStart map[string]int

func NewObject(renderer *sdl.Renderer) (objectData []object, err error) {

	Surf, _ := sdl.LoadBMP("SPRITES/LEVELONE/levelOneLayout.bmp")
	Tex, _ := renderer.CreateTextureFromSurface(Surf)

	PlayerStart = make(map[string]int, 2)
	PlayerStart["X"] = 96
	PlayerStart["Y"] = 430

	blockOne := object{
		Tex:          Tex,
		x:            0,
		y:            542,
		objectWidth:  388,
		objectHeight: 97,
	}
	blockTwo := object{
		Tex:          Tex,
		x:            481,
		y:            404,
		objectWidth:  347,
		objectHeight: 106,
	}
	blockThree := object{
		Tex:          Tex,
		x:            987,
		y:            267,
		objectWidth:  282,
		objectHeight: 106,
	}
	objectData = append(objectData, blockOne, blockTwo, blockThree)

	defer Surf.Free()

	Surf, _ = sdl.LoadBMP("SPRITES/LEVELONE/BG.bmp")
	BG, _ := renderer.CreateTextureFromSurface(Surf)

	backgroundData = object{
		Tex:          BG,
		x:            0,
		y:            0,
		objectWidth:  1280,
		objectHeight: 720,
	}

	defer Surf.Free()

	return objectData, nil

}

func (o object) Draw(renderer *sdl.Renderer, objData []object) {

	backgroundData.Tex.SetBlendMode(sdl.BLENDMODE_BLEND)
	renderer.Copy(backgroundData.Tex, nil, nil)

	for _, k := range objData {
		renderer.Copy(k.Tex,
			&sdl.Rect{X: int32(k.x), Y: int32(k.y), W: int32(k.objectWidth), H: int32(k.objectHeight)},
			&sdl.Rect{X: int32(k.x), Y: int32(k.y), W: int32(k.objectWidth), H: int32(k.objectHeight)},
		)
	}

}

func (o object) Update() {

}
