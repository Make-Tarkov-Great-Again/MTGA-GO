package tools

import (
	"golang.org/x/exp/constraints"
	"log"
	"math"
	"math/rand"
)

/*
GetRandomInt returns a random int between min and max

	If max is less than or equal to min, max is set to 100
	To get a random int between 0 and 100, pass 0 as min and 100 as max
*/
func GetRandomInt(min, max int) int {
	if max <= min {
		log.Println("You've got your numbers backwards, fool")
		return rand.Intn(min-max+1) + max
	}

	return rand.Intn(max-min+1) + min
}

func RoundToThousandths(value float32) float32 {
	return float32(math.Round(float64(value)*1000) / 1000)
}

// GetRandomSplitInt returns a slice of random ints that add up to the total
func GetRandomSplitInt(total int) []int {
	output := make([]int, 0)
	for total > 0 {
		sum := rand.Intn(total + 1)
		output = append(output, sum)
		total -= sum
	}
	return output
}

type NumericType interface {
	constraints.Signed
}

func LevelComparisonCheck[T NumericType](requiredLevel T, currentLevel T, compareMethod string) bool {
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
