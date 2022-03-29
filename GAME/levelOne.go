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

func NewObject(renderer *sdl.Renderer) (o object, err error) {

	Surf, _ := sdl.LoadBMP("SPRITES/LEVELONE/1.bmp")
	o.objectWidth = Surf.Bounds().Dx()
	o.objectHeight = Surf.Bounds().Dy()
	defer Surf.Free()

	o.Tex, _ = renderer.CreateTextureFromSurface(Surf)
	if err != nil {
		return
	}

	o.x = 200
	o.y = 200

	return o, nil
}

func (o object) Draw(renderer *sdl.Renderer) {
	renderer.Copy(o.Tex,
		nil,
		&sdl.Rect{X: int32(o.x), Y: int32(o.y), W: int32(o.objectWidth), H: int32(o.objectHeight)},
	)
}

func (o object) Update() {

}
