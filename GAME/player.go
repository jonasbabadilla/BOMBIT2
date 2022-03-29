package sdl

import (
	"github.com/veandco/go-sdl2/sdl"
)

type player struct {
	pTex *sdl.Texture
}

func newPlayer(renderer *sdl.Renderer) (p player, e error) {

}

func (p *player) Draw(renderer *sdl.Renderer) {

}

func (p *player) Update() {
	keys := sdl.GetKeyBoardState()
}
