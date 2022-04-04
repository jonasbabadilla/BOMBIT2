package main

import (
	levels "chaseGame/GAME/LEVELS"
)

var currentLvl = 1
var totalLvl = 4
var LevelObjects []levels.Object
var LevelBG levels.Object
var pStart levels.StartData

func decideLevel() ([]levels.Object, levels.Object, levels.StartData) {

	switch currentLvl {
	case 1:
		LevelObjects, LevelBG, pStart, _ = levels.LevelOne(Renderer)
		return LevelObjects, LevelBG, pStart
	case 2:
		LevelObjects, LevelBG, pStart, _ = levels.LevelTwo(Renderer)
		return LevelObjects, LevelBG, pStart
	case 3:
		LevelObjects, LevelBG, pStart, _ = levels.LevelThree(Renderer)
		return LevelObjects, LevelBG, pStart
	case 4:
		LevelObjects, LevelBG, pStart, _ = levels.LevelFour(Renderer)
		return LevelObjects, LevelBG, pStart
	}

	return nil, levels.Object{}, levels.StartData{}
}
