package main

import (
	levels "chaseGame/GAME/LEVELONE"
)

var currentLvl = 1
var LevelObjData [][]levels.Object

func decideLevel() []levels.Object {

	LevelObjects, _ := levels.NewObject(Renderer)
	switch currentLvl {
	case 1:
		LevelObjData = append(LevelObjData, LevelObjects)
		return LevelObjData[0]
	}
	return nil
}
