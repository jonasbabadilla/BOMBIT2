package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type player struct {
	Tex                       *sdl.Texture
	Tex2                      *sdl.Texture
	x, y                      float64
	playerWidth, playerHeight int
}

var CharFrameX int32
var CharFrameY int32

var JumpTimer int

var gravity = 5.00
var JumpState bool

var charDirection sdl.RendererFlip
var BotDirection sdl.RendererFlip

var PlayerSpeedX = 3.00
var PlayerSpeedY = 7.00

var BotX int
var BotY int

var Keys = sdl.GetKeyboardState()

func NewPlayer(renderer *sdl.Renderer) (p player, e error) {

	char, _ := sdl.LoadBMP("CharSprites/playerOneSprite.bmp")
	Tex, _ := renderer.CreateTextureFromSurface(char)

	char2, _ := sdl.LoadBMP("CharSprites/playerTwoSprite.bmp")
	Tex2, _ := renderer.CreateTextureFromSurface(char2)

	p = player{
		Tex:          Tex,
		Tex2:         Tex2,
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

func (p *player) DrawTwo(renderer *sdl.Renderer) {

	switched = true
	renderer.CopyEx(p.Tex2,
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

func (p *player) BotOne(renderer *sdl.Renderer) {

	if switched == false {
		renderer.CopyEx(p.Tex2,
			&sdl.Rect{X: 0, Y: 0, W: 16, H: 16},
			&sdl.Rect{X: int32(BotX), Y: int32(BotY), W: 64, H: 64},
			0.0,
			&sdl.Point{},
			BotDirection,
		)
	} else {
		renderer.CopyEx(p.Tex,
			&sdl.Rect{X: 0, Y: 0, W: 16, H: 16},
			&sdl.Rect{X: int32(BotX), Y: int32(BotY), W: 64, H: 64},
			0.0,
			&sdl.Point{},
			BotDirection,
		)
	}
}

func (p *player) EndSpawn(renderer *sdl.Renderer) {

	renderer.CopyEx(p.Tex2,
		&sdl.Rect{X: CharFrameX, Y: CharFrameY, W: 16, H: 16},
		&sdl.Rect{X: 1000, Y: 600, W: 64, H: 64},
		0.0,
		&sdl.Point{},
		charDirection,
	)
}

func (p *player) FinalDraw(renderer *sdl.Renderer) {

	renderer.CopyEx(p.Tex2,
		&sdl.Rect{X: CharFrameX, Y: CharFrameY, W: 16, H: 16},
		&sdl.Rect{X: 1050, Y: 485, W: 64, H: 64},
		0.0,
		&sdl.Point{},
		charDirection,
	)

}
