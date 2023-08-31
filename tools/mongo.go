package tools

import (
	"crypto/rand"
	"errors"
	"math"
)

const alphabet string = "abcdef0123456789"

var alphabetLength int = len(alphabet)

// taken from https://github.com/matoous/go-nanoid/blob/master/gonanoid.go
// because i didnt need to use the whole package 8===D

const size int = 24

var chars = []rune(alphabet)
var length int = len(chars)

// GenerateMongoId returns a random string of length 24
func GenerateMongoID() (string, error) {
	if alphabetLength == 0 || alphabetLength > 255 {
		return "", errors.New("alphabet must not be empty and contain no more than 255 chars")
	}
	if size <= 0 {
		return "", errors.New("size must be positive integer")
	}

	mask := getMask(length)
	// estimate how many random bytes we will need for the ID, we might actually need more but this is tradeoff
	// between average case and worst case
	ceilArg := 1.6 * float64(mask*size) / float64(alphabetLength)
	step := int(math.Ceil(ceilArg))

	id := make([]rune, size)
	bytes := make([]byte, step)
	for j := 0; ; {
		_, err := rand.Read(bytes)
		if err != nil {
			return "", err
		}
		for i := 0; i < step; i++ {
			currByte := bytes[i] & byte(mask)
			if currByte < byte(length) {
				id[j] = chars[currByte]
				j++
				if j == size {
					return string(id[:size]), nil
				}
			}
		}
	}
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
