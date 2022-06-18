package random

import (
	"math/rand"
	"time"
)

func DefaultSeed() int64 {
	return time.Now().UnixNano()
}

func Random(seed int64) *rand.Rand {
	return rand.New(rand.NewSource(seed))
}

func Intn(n int) int {
	return Random(DefaultSeed()).Intn(n)
}
