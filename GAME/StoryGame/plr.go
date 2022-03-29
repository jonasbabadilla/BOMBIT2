package StoryGame

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type player struct {
	Tex  *sdl.Texture
	x, y float64
}

var Width int32
var Height int32

const (
	playerSpeed = 0.75
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
	p.x = float64(Width)/2.0 - playerSize/2.0
	p.y = float64(Height)/2.0 - playerSize/2.0
	return p, nil
}

func (p *player) Draw(renderer *sdl.Renderer) {
	// For some reason when I use this instead of p.x/p.y in Copy, the whole thing doesn't work anymore
	// will need to look over that
	//x := playerSize / 2.0
	//y := playerSize / 2.0

	renderer.Copy(p.Tex,
		&sdl.Rect{X: 0, Y: 0, W: 64, H: 64},
		&sdl.Rect{X: int32(p.x), Y: int32(p.y), W: 64, H: 64})

}

func (p *player) Update() {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_LEFT] == 1 {
		p.x -= playerSpeed
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {

		p.x += playerSpeed
	}

}

// renderer.CopyEx() another important function, has other parameters like center point, rotate, flip, etc
