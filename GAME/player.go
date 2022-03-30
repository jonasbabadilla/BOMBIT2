package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type player struct {
	Tex                       *sdl.Texture
	x, y                      float64
	playerWidth, playerHeight int
	//LastPress                 time.Time
}

var JumpTimer int

var gravity = 5.00
var JumpState bool

var charDirection sdl.RendererFlip

var PlayerSpeedX = 3.00
var PlayerSpeedY = 7.00

func NewPlayer(renderer *sdl.Renderer) (p player, e error) {

	char, err := sdl.LoadBMP("SPRITES/playerSprite.bmp")
	p.playerWidth = char.Bounds().Dx()
	p.playerHeight = char.Bounds().Dy()
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

	renderer.CopyEx(p.Tex,
		nil,
		&sdl.Rect{X: int32(p.x), Y: int32(p.y), W: 42, H: 62},
		0.0,
		&sdl.Point{X: 10, Y: 0},
		charDirection,
	)
}

func (p *player) Update() {
	keys := sdl.GetKeyboardState()
	if keys[sdl.SCANCODE_LEFT] == 1 {
		p.x -= PlayerSpeedX
		charDirection = sdl.FLIP_HORIZONTAL

	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		p.x += PlayerSpeedX
		charDirection = sdl.FLIP_NONE
	}

	if keys[sdl.SCANCODE_UP] == 1 && JumpState != true {
		JumpState = true
	}

	p.CalcJump()
	p.y += float64(gravity)
	sdl.Delay(10)
}

func (p *player) CalcJump() {
	if JumpState {
		JumpTimer++
		p.y -= PlayerSpeedY
		if JumpTimer >= 20 {
			PlayerSpeedY *= 0.90
		}
		if JumpTimer >= 37 {
			if gravity < 5 {
				gravity *= 1.1
			}
		}
	}
	checkCollision()
}
