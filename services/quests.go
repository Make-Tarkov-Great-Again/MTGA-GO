package services

import (
	"fmt"
)

var bearOnlyQuests = map[string]bool{
	"6179b5eabca27a099552e052": true,
	"5e383a6386f77465910ce1f3": true,
	"5e4d515e86f77438b2195244": true,
	"639282134ed9512be67647ed": true,
}

var usecOnlyQuests = map[string]bool{
	"6179b5eabca27a099552e052": true,
	"5e383a6386f77465910ce1f3": true,
	"5e4d515e86f77438b2195244": true,
	"639282134ed9512be67647ed": true,
}

func CheckIfQuestForOtherFaction(side string, qid string) bool {
	if side == "Bear" {
		return usecOnlyQuests[qid]
	} else {
		return bearOnlyQuests[qid]
	}
}

func LevelComparisonCheck(requiredLevel float64, currentLevel float64, compareMethod string) bool {
	switch compareMethod {
	case ">=":
		return currentLevel >= requiredLevel
	case ">":
		return currentLevel > requiredLevel
	case "<":
		return currentLevel < requiredLevel
	case "<=":
		return currentLevel <= requiredLevel
	case "=":
		return currentLevel == requiredLevel
	default:
		log.Println("Unknown comparison method of " + compareMethod)
		return false
	}
}
