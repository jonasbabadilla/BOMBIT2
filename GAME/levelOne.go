package main

import (
	"fmt"

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

func NewObject(renderer *sdl.Renderer) (objectData []object, err error) {

	Surf, _ := sdl.LoadBMP("SPRITES/LEVELONE/1.bmp")
	Tex, _ := renderer.CreateTextureFromSurface(Surf)
	tempObj, _ = renderer.CreateTextureFromSurface(Surf)
	o := object{
		Tex:          Tex,
		x:            600,
		y:            500,
		objectWidth:  Surf.Bounds().Dx(),
		objectHeight: Surf.Bounds().Dy(),
	}
	objectData = append(objectData, o)
	fmt.Print(objectData)
	defer Surf.Free()

	Surf, _ = sdl.LoadBMP("SPRITES/LEVELONE/BG.bmp")
	BG, _ := renderer.CreateTextureFromSurface(Surf)
	backgroundData = object{
		Tex:          BG,
		x:            0,
		y:            0,
		objectWidth:  Surf.Bounds().Dx(),
		objectHeight: Surf.Bounds().Dy(),
	}
	defer Surf.Free()
	return objectData, nil

}

func (o object) Draw(renderer *sdl.Renderer, objData []object) {

	backgroundData.Tex.SetBlendMode(sdl.BLENDMODE_BLEND)
	renderer.Copy(backgroundData.Tex, nil, nil)

	for _, k := range objData {
		renderer.Copy(k.Tex,
			nil,
			&sdl.Rect{X: int32(k.x), Y: int32(k.y), W: 429, H: 74},
		)
	}

}

func (o object) Update() {

}
