package tools

import (
	"crypto/rand"
	"math"
)

// taken from https://github.com/matoous/go-nanoid/blob/master/gonanoid.go
// because i didnt need to use the whole package 8===D

const size int = 24

var chars = []rune("abcdef0123456789")
var alphabetLength = len(chars)

// GenerateMongoID returns a random string of length 24
func GenerateMongoID() string {
	mask := getMask(alphabetLength)

	id := make([]rune, size)
	bytes := make([]byte, int(math.Ceil(1.6*float64(mask*size)/float64(alphabetLength))))
	if _, err := rand.Read(bytes); err != nil {
		return ""
	}

	for j := 0; j < size; {
		for i := 0; i < size && j < size; i++ {
			currByte := bytes[i] & byte(mask)
			if currByte < byte(alphabetLength) {
				id[j] = chars[currByte]
				j++
			}
		}
	}

	return string(id)
}

func getMask(alphabetSize int) int {
	for i := 1; i <= 8; i++ {
		mask := (2 << uint(i)) - 1
		if mask >= alphabetSize-1 {
			return mask
		}
	}
	return 0
}
