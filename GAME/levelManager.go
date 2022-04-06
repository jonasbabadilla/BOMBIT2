package main

import (
	levels "chaseGame/GAME/LEVELS"
	"fmt"
)

var currentLvl = 1
var totalLvl = 18
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
	case 8:
		LevelObjects, LevelBG, pStart, _ = levels.LevelEight(Renderer)
		return LevelObjects, LevelBG, pStart, textData
	case 9:
		LevelObjects, LevelBG, pStart, _ = levels.LevelNine(Renderer)
		return LevelObjects, LevelBG, pStart, levels.Object{}
	case 10:
		LevelObjects, LevelBG, pStart, _ = levels.LevelTen(Renderer)
		return LevelObjects, LevelBG, pStart, levels.Object{}
	case 11:
		LevelObjects, LevelBG, pStart, _ = levels.LevelEleven(Renderer)
		return LevelObjects, LevelBG, pStart, levels.Object{}
	case 12:
		LevelObjects, LevelBG, pStart, _ = levels.LevelTwelve(Renderer)
		return LevelObjects, LevelBG, pStart, levels.Object{}
	case 13:
		LevelObjects, LevelBG, pStart, _ = levels.LevelThirteen(Renderer)
		return LevelObjects, LevelBG, pStart, levels.Object{}
	case 14:
		LevelObjects, LevelBG, pStart, _ = levels.LevelFourteen(Renderer)
		return LevelObjects, LevelBG, pStart, levels.Object{}
	case 15:
		LevelObjects, LevelBG, pStart, _ = levels.LevelFifteen(Renderer)
		return LevelObjects, LevelBG, pStart, levels.Object{}

	}

	return nil, levels.Object{}, levels.StartData{}, levels.Object{}
}
