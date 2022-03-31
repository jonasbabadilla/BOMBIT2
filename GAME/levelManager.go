package main

import (
	levels "chaseGame/GAME/LEVELONE"
)

var currentLvl = 1
var LevelObjData [][]levels.Object

func decideLevel() ([]levels.Object, levels.Object) {

	LevelObjects, LevelBG, _ := levels.LevelOne(Renderer)
	switch currentLvl {
	case 1:
		LevelObjData = append(LevelObjData, LevelObjects)
		return LevelObjData[0], LevelBG
	}
	return nil, levels.Object{}
}
