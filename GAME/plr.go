package main

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type player struct {
	Tex       *sdl.Texture
	x, y      float64
	LastPress time.Time
}

var gravity = 1.00
var jump int
var gravCount int

const (
	playerSpeedX = 1.5
	playerSpeedY = 30
	playerSize   = 64
	CoolDown     = time.Millisecond * 500
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
		p.x -= playerSpeedX
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		p.x += playerSpeedX
	}

	if keys[sdl.SCANCODE_UP] == 1 {
		if time.Since(p.LastPress) > CoolDown {
			p.y -= (playerSpeedY)
			p.LastPress = time.Now()
			sdl.Delay(500)
		} else {
			p.y += (playerSpeedY)
		}

	}
}
