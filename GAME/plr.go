package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type player struct {
	Tex  *sdl.Texture
	x, y float64
}

const playerSpeed = 0.75

func newPlayer(renderer *sdl.Renderer) (p player, e error) {

	char, err := sdl.LoadBMP("SPRITES/playerSprite.bmp")
	if err != nil {
		return player{}, fmt.Errorf("loading player sprite: %v", err)
	}
	defer char.Free()
	p.Tex, err = renderer.CreateTextureFromSurface(char)
	if err != nil {
		return player{}, fmt.Errorf("loading player texture: %v", err)
	}

	p.x = 0.0
	p.y = 680

	return p, nil
}

func (p *player) Draw(renderer *sdl.Renderer) {
	// x: 3 y: 375
	renderer.Copy(p.Tex,
		&sdl.Rect{X: 0, Y: 0, W: 32, H: 32},
		&sdl.Rect{X: int32(p.x), Y: int32(p.y), W: 32, H: 32})

}

func (p *player) Update() {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_LEFT] == 1 {
		p.x -= playerSpeed
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		p.x += playerSpeed
	} else if keys[sdl.SCANCODE_LEFT] == 1 && keys[sdl.SCANCODE_RIGHT] == 1 {
		p.x += 0
	}

}
