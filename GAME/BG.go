package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type player struct{
	texture *sdl.Texture
}

func Stickman(rednerer *sdl.Renderer) (p player, err error) {
	var p player
	
	img, err := sdl.LoadBMP("SPRITES/BG1.bmp")
	if err != nil {
		return player{}, fmt.Errorf("loading player sprite:", err)
	}
	
	defer img.Free() 
	p.texture, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		return player{}, fmt.Errof("loading player texture:", err)
	}
	return p, nil
}

// currently not in use