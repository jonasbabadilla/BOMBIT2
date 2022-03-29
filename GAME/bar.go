package MainGame

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

type object struct {
	Tex   *sdl.Texture
	x, y  float64
	angle float64
}

const (
	objectSpeed  = 1.1
	objectWidth  = 128
	objectHeight = 16
)

func NewObject(renderer *sdl.Renderer) (o object, err error) {

	object, err := sdl.LoadBMP("SPRITES/BAR.bmp")
	if err != nil {
		return
	}
	defer object.Free()
	o.Tex, err = renderer.CreateTextureFromSurface(object)

	o.angle = 270 * (math.Pi / 180)
	o.x = 200
	o.y = 400

	return o, nil
}

func (o object) Draw(renderer *sdl.Renderer) {

	renderer.Copy(o.Tex,
		&sdl.Rect{X: 0, Y: 0, W: objectWidth, H: objectHeight},
		&sdl.Rect{X: int32(o.x), Y: int32(o.y), W: objectWidth, H: objectHeight})

}

func (o object) Update() {
	// trying to get it to move left and right, rn trying to just make it move in one direction lol
	o.y += objectSpeed * math.Sin(o.angle)

}
