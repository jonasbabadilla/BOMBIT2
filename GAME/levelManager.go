package main

import (
	levels "chaseGame/GAME/LEVELS"
	"fmt"
)

var currentLvl = 7
var totalLvl = 16
var LevelObjects []levels.Object
var LevelBG levels.Object
var pStart levels.StartData
var textData levels.Object

func decideLevel() ([]levels.Object, levels.Object, levels.StartData, levels.Object) {
	fmt.Println("Current Level:", currentLvl)

	switch currentLvl {
	case 1:
		LevelObjects, LevelBG, pStart, textData, _ = levels.LevelOne(Renderer)
		return LevelObjects, LevelBG, pStart, textData
	case 2:
		LevelObjects, LevelBG, pStart, _ = levels.LevelTwo(Renderer)
		return LevelObjects, LevelBG, pStart, levels.Object{}
	case 3:
		LevelObjects, LevelBG, pStart, _ = levels.LevelThree(Renderer)
		return LevelObjects, LevelBG, pStart, levels.Object{}
	case 4:
		LevelObjects, LevelBG, pStart, _ = levels.LevelFour(Renderer)
		return LevelObjects, LevelBG, pStart, levels.Object{}
	case 5:
		LevelObjects, LevelBG, pStart, _ = levels.LevelFive(Renderer)
		return LevelObjects, LevelBG, pStart, levels.Object{}
	case 6:
		LevelObjects, LevelBG, pStart, _ = levels.LevelSix(Renderer)
		return LevelObjects, LevelBG, pStart, levels.Object{}
	case 7:
		LevelObjects, LevelBG, pStart, _ = levels.LevelSeven(Renderer)
		return LevelObjects, LevelBG, pStart, levels.Object{}
	}

	return nil, levels.Object{}, levels.StartData{}, levels.Object{}
}
