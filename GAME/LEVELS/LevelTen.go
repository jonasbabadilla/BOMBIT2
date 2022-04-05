package levels

import (
	"github.com/veandco/go-sdl2/sdl"
)

//type LevelUpdate int

//var Projectile = sdl.Rect{X: 10, Y: 10, W: 10, H: 10}

func LevelTen(renderer *sdl.Renderer) (levelData []Object, LevelBG Object, PlayerStart StartData, err error) {

	Surf, _ := sdl.LoadBMP("LEVELS/LevelTenSprites/levelLayout.bmp")
	Tex, _ := renderer.CreateTextureFromSurface(Surf)

	levelData = CreateLevel(Surf, Tex)

	defer Surf.Free()

	Surf, _ = sdl.LoadBMP("LEVELS/LevelTenSprites/BG.bmp")
	BG, _ := renderer.CreateTextureFromSurface(Surf)

	backgroundData = Object{
		Tex:          BG,
		X:            0,
		Y:            0,
		ObjectWidth:  1280,
		ObjectHeight: 720,
	}

	defer Surf.Free()

	// When you look at the levelLayout in LevelThreeSprites folder, do not be fooled by the x coordinates
	// of the pink box. I was and I lost 45 minutes in real-time minutes and 2 years of my life span
	PlayerStart = StartData{X: 50, Y: 477, EndData: struct {
		X int
		Y int
	}{X: 1132, Y: 630}}

	return levelData, backgroundData, PlayerStart, nil

}

/*
func (s *LevelUpdate) Update() {
	Projectile.X += 1

}
*/
