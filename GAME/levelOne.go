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

var backgroundData object

func NewObject(renderer *sdl.Renderer) (objectData []object, err error) {

	Surf, _ := sdl.LoadBMP("SPRITES/LEVELONE/1.bmp")
	objectData = append(objectData, object{})
	objectData[0].objectWidth = Surf.Bounds().Dx()
	objectData[0].objectHeight = Surf.Bounds().Dy()
	objectData[0].Tex, _ = renderer.CreateTextureFromSurface(Surf)
	objectData[0].x = 600
	objectData[0].y = 500
	defer Surf.Free()

	Surf, _ = sdl.LoadBMP("SPRITES/LEVELONE/BG.bmp")
	BG, _ := renderer.CreateTextureFromSurface(Surf)
	backgroundData.Tex = BG
	backgroundData.objectWidth = Surf.Bounds().Dx()
	backgroundData.objectHeight = Surf.Bounds().Dy()
	backgroundData.x = 0
	backgroundData.y = 0
	defer Surf.Free()
	return objectData, nil
}

func (o object) Draw(renderer *sdl.Renderer) {
	renderer.Copy(backgroundData.Tex, nil, nil)

	renderer.Copy(o.Tex,
		nil,
		&sdl.Rect{X: int32(o.x), Y: int32(o.y), W: int32(o.objectWidth), H: int32(o.objectHeight)},
	)

}

func (o object) Update() {

}
