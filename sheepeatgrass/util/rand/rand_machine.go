package rand

import (
	"math/rand"
	"time"
)

var fixedInt int
var isFixed bool

func RandInt(max int) int {
	if !isFixed {
		rand.Seed(time.Now().UnixNano())
		return rand.Intn(max)
	}
	return fixedInt
}

func ResetRandom() {
	isFixed = false
}

func FixRandInt(fixed int) {
	isFixed = true
	fixedInt = fixed
}
