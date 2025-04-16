package random

import (
	"math/rand"
	"time"
)

func NewRandomString(size int) string {

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	chars := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]byte, size)
	for i := range b {
		b[i] = chars[rnd.Intn(len(chars))]
	}

	return string(b)

}
