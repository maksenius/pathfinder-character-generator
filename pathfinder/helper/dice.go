package helper

import (
	"math/rand"
	"time"
)

func D6() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(6) + 1
}
