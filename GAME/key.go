package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Key struct {
	Tex                       *sdl.Texture
	x, y                      float64
	keyWidth, keyHeight int
	
}


func NewKey(renderer *sdl.Renderer) (p player, e error) {

	key, _ := sdl.LoadBMP("CharSprites/KEY.bmp")
	Tex, _ := renderer.CreateTextureFromSurface(key)

	p = player{
		Tex:          Tex,
		x:            float64(pStart.X),
		y:            float64(pStart.Y),
		playerWidth:  16,
		playerHeight: 16,
	}
	defer key.Free()
	return p, nil
}

func (k *Key) Draw(renderer *sdl.Renderer) {

	renderer.Copy(k.Tex,
	&sdl.Rect{},
	&sdl.Rect{})

	&sdl.Rect{X: CharFrameX, Y: CharFrameY, W: 16, H: 16},
		&sdl.Rect{X: int32(p.x), Y: int32(p.y), W: 64, H: 64},
	
}


	