package world

import (
	"fmt"
	"testing"
)

func TestPoint2D_move(t *testing.T) {
	p := Point2D{0, 0}
	p.Move(East)

	expectPt := Point2D{1, 0}
	if p != expectPt {
		t.Errorf("%v doesn't go East", p)
	}
	fmt.Println(p)

	p.Move(North)

	fmt.Println(p)
	if p.X != 1 || p.Y != 1 {
		t.Errorf("%v doesn't go North", p)
	}
}
