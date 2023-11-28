package tools

import (
	"log"
	"math"
	"math/rand"
	"time"

	"golang.org/x/exp/constraints"
)

var thyme = time.Now().UnixNano()
var random = rand.New(rand.NewSource(thyme))

/*
GetRandomInt returns a random int between min and max

	If max is less than or equal to min, max is set to 100
	To get a random int between 0 and 100, pass 0 as min and 100 as max
*/
func GetRandomInt(min int, max int) int {
	if max <= min {
		max = 100
	} else {
		return random.Intn(max-min+1) + min
	}
	return min
}

func RoundToThousandths(value float32) float32 {
	return float32(math.Round(float64(value)*1000) / 1000)
}

// GetFloat returns a random float between min and max
func GetRandomFloat(min float64, max float64) float64 {
	return random.Float64()*(max-min) + min
}

// GetRandomSplitInt returns a slice of random ints that add up to the total
func GetRandomSplitInt(total int) []int {
	output := make([]int, 0)
	for total > 0 {
		sum := random.Intn(total + 1)
		output = append(output, sum)
		total -= sum
	}
	return output
}

// GetRandomFromArray returns a random value from the array passed
func GetRandomFromArray(array []int) int {
	randIndex := rand.Intn(len(array))
	return array[randIndex]
}

// GetRandomFromObject returns a random value from the object passed
func GetRandomFromObject(obj map[string]int) int {
	keys := make([]string, 0, len(obj))
	for k := range obj {
		keys = append(keys, k)
	}
	randIndex := rand.Intn(len(keys))
	return obj[keys[randIndex]]
}

// GetPercentRandomBool returns a random bool based on the percentage int passed
func GetPercentRandomBool(percentage int) bool {
	return rand.Intn(100) < percentage
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
