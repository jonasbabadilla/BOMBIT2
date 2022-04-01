package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type player struct {
	Tex                       *sdl.Texture
	x, y                      float64
	playerWidth, playerHeight int
}

var CharFrameX int32
var CharFrameY int32

var JumpTimer int

var gravity = 5.00
var JumpState bool

var charDirection sdl.RendererFlip

var PlayerSpeedX = 3.00
var PlayerSpeedY = 7.00

var Keys = sdl.GetKeyboardState()

func NewPlayer(renderer *sdl.Renderer) (p player, e error) {

	char, _ := sdl.LoadBMP("CharSprites/playerSprite.bmp")
	Tex, _ := renderer.CreateTextureFromSurface(char)

	p = player{
		Tex:          Tex,
		x:            float64(pStart.X),
		y:            float64(pStart.Y),
		playerWidth:  16,
		playerHeight: 16,
	}
	defer char.Free()
	return p, nil
}

func (p *player) Draw(renderer *sdl.Renderer) {

	renderer.CopyEx(p.Tex,
		&sdl.Rect{X: CharFrameX, Y: CharFrameY, W: 16, H: 16},
		&sdl.Rect{X: int32(p.x), Y: int32(p.y), W: 64, H: 64},
		0.0,
		&sdl.Point{},
		charDirection,
	)
}

func (p *player) Update() {
	Keys = sdl.GetKeyboardState()
	if Keys[sdl.SCANCODE_LEFT] == 1 {
		p.x -= PlayerSpeedX
		charDirection = sdl.FLIP_HORIZONTAL

	} else if Keys[sdl.SCANCODE_RIGHT] == 1 {
		p.x += PlayerSpeedX
		charDirection = sdl.FLIP_NONE
	}

	if Keys[sdl.SCANCODE_UP] == 1 && JumpState != true {
		JumpState = true
		gravity = 1.00
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
