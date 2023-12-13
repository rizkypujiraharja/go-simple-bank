package utils

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func RandomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := make([]byte, n)
	for i := range bytes {
		bytes[i] = letters[rand.Intn(len(letters))]
	}
	return string(bytes)
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return int64(RandomInt(0, 1000))
}
