package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type player struct {
	Tex  *sdl.Texture
	x, y float64
}

var gravCount = 0

const (
<<<<<<< HEAD
	playerSpeed = 1.5
=======
	playerSpeed = 0.2
>>>>>>> 6d822491bdb28ee250669dcc40326e60b0a9da51
	playerSize  = 64
)

func NewPlayer(renderer *sdl.Renderer) (p player, e error) {

	char, err := sdl.LoadBMP("SPRITES/playerSprite.bmp")
	if err != nil {
		return player{}, fmt.Errorf("loading player sprite: %v", err)
	}
	defer char.Free()
	p.Tex, err = renderer.CreateTextureFromSurface(char)
	if err != nil {
		return player{}, fmt.Errorf("loading player texture: %v", err)
	}

	p.x = 640 /// 2.0
	p.y = 360 //- playerSize/2.0

	return p, nil
}

func (p *player) Draw(renderer *sdl.Renderer) {
	//x := playerSize / 2.0
	//y := playerSize / 2.0

	renderer.Copy(p.Tex,
		&sdl.Rect{X: 0, Y: 0, W: 21, H: 31},
		&sdl.Rect{X: int32(p.x), Y: int32(p.y), W: 42, H: 62})
}
func (p *player) Update() {
	keys := sdl.GetKeyboardState()
	if keys[sdl.SCANCODE_LEFT] == 1 {
		p.x -= playerSpeed
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		p.x += playerSpeed
	}

	var jump = keys[sdl.SCANCODE_UP]

	if jump == 1 {
		for gravCount := 0; gravCount < 10; {
			p.y -= (playerSpeed / (p.y / 1.5))
			gravCount++
		}
	}
}
