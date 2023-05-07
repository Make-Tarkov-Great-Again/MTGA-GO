package tools

import (
	"github.com/jaevor/go-nanoid"
)

// GenerateMongoId returns a random string of length 24
func GenerateMongoId() string {
	id, err := nanoid.CustomUnicode("1234567890abcdef", 24)
	if err != nil {
		panic(err)
	}
	return id()
}
