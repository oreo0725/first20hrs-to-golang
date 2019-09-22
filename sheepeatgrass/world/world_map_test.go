package world

import (
	"fmt"
	"testing"
)

func TestWorld_String(t *testing.T) {

	w := World{DAY: 0}
	w.MAP[1][1] = NewSheep("1", Point2D{})
	w.MAP[1][2] = NewGrass("11", Point2D{})
	fmt.Println(w)
}
