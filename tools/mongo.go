package tools

import (
	crand "crypto/rand"
	"errors"
	"sync"
)

// GenerateMongoId returns a random string of length 24
func GenerateMongoId() string {
	id, err := standard(24)
	if err != nil {
		panic(err)
	}
	return id()
}

// all functions below were created @ https://github.com/jaevor/go-nanoid/blob/v1.3.0/nanoid.go#L22
// i stripped out what i needed to avoid external dependencies that werent needed

type generator = func() string

var standardAlphabet [64]byte = [64]byte{
	'A', 'B', 'C', 'D', 'E',
	'F', 'G', 'H', 'I', 'J',
	'K', 'L', 'M', 'N', 'O',
	'P', 'Q', 'R', 'S', 'T',
	'U', 'V', 'W', 'X', 'Y',
	'Z', 'a', 'b', 'c', 'd',
	'e', 'f', 'g', 'h', 'i',
	'j', 'k', 'l', 'm', 'n',
	'o', 'p', 'q', 'r', 's',
	't', 'u', 'v', 'w', 'x',
	'y', 'z', '0', '1', '2',
	'3', '4', '5', '6', '7',
	'8', '9', '-', '_',
}

var ErrInvalidLength = errors.New("length for ID is invalid (must be within 2-255)")

func standard(length int) (generator, error) {
	if length < 2 || length > 255 {
		return nil, ErrInvalidLength
	}

	// Multiplying to increase the 'buffer' so that .Read()
	// has to be called less, which is more efficient.
	// b holds the random crypto bytes.
	size := length * length * 7
	b := make([]byte, size)
	crand.Read(b)
	offset := 0

	// Since the standard alphabet is ASCII, we don't have to use runes.
	// ASCII max is 128, so byte will be perfect.
	id := make([]byte, length)

	var mu sync.Mutex

	return func() string {
		mu.Lock()
		defer mu.Unlock()

		// If all the bytes in the slice
		// have been used, refill.
		if offset == size {
			crand.Read(b)
			offset = 0
		}

		for i := 0; i < length; i++ {
			/*
				"It is incorrect to use bytes exceeding the alphabet size.
				The following mask reduces the random byte in the 0-255 value
				range to the 0-63 value range. Therefore, adding hacks such
				as empty string fallback or magic numbers is unneccessary because
				the bitmask trims bytes down to the alphabet size (64)."
			*/
			// Index using the offset.
			id[i] = standardAlphabet[b[i+offset]&63]
		}

		// Extend the offset.
		offset += length

		return string(id)
	}, nil
}
