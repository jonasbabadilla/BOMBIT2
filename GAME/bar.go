package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type object struct {
	Tex  *sdl.Texture
	x, y float64
}

const (
	objectSpeed  = 1.1
	objectWidth  = 128
	objectHeight = 16
)

/*
func newBar(renderer *sdl.Renderer) (o object) {
	o.Tex = renderer.CreateTextureFromSurface(renderer, "SPRITES/BAR.bmp")
	return o
}
*/

func newObject(renderer *sdl.Renderer) (o object, e error) {

	object, err := sdl.LoadBMP("SPRITES/BAR.bmp")
	if err != nil {
		return o, fmt.Errorf("loading object sprite: %v", err)
	}
	defer object.Free()
	o.Tex, err = renderer.CreateTextureFromSurface(object)
	if err != nil {
		return o, fmt.Errorf("loading object texture: %v", err)
	}

	o.x = 200
	o.y = 200
	return o, nil
}

func (o object) Draw(renderer *sdl.Renderer) {

	renderer.Copy(o.Tex,
		&sdl.Rect{X: 0, Y: 0, W: 128, H: 16},
		&sdl.Rect{X: int32(o.x), Y: int32(o.y), W: 128, H: 16})

}

func (o object) Update() {

	for i := 0; i < 25; i++ {
		o.x -= objectSpeed
	}

	for i := 0; i < 25; i++ {
		o.x += objectSpeed
	}
}
