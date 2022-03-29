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

var gravity = 1
var JumpState bool

const (
	playerSpeedX = 1.00
	playerSpeedY = 5.00
	//CoolDown     = time.Millisecond * 5000
)

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

	renderer.Copy(p.Tex,
		nil,
		&sdl.Rect{X: int32(p.x), Y: int32(p.y), W: 42, H: 62},
	)
}
func (p *player) Update() {
	keys := sdl.GetKeyboardState()
	if keys[sdl.SCANCODE_LEFT] == 1 {
		p.x -= playerSpeedX
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		p.x += playerSpeedX
	}

	if keys[sdl.SCANCODE_UP] == 1 && JumpState != true {
		JumpState = true
	}

	p.CalcJump()
	p.y += float64(gravity)
	sdl.Delay(1000 / 60)
}

func (p *player) CalcJump() {
	if JumpState {
		p.y -= playerSpeedY
		JumpTimer++
		if JumpTimer >= 20 {
			p.y += playerSpeedY + float64(gravity)
		}
	}
	checkCollision()
}
