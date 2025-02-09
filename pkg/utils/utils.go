package utils

import (
	"math/rand"
	"time"
)

func GenerateRandomString(alphabet string, length int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = alphabet[r.Intn(len(alphabet))]
	}

	return string(b)
}
