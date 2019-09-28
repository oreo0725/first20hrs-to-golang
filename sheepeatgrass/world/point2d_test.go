package world

import (
	"fmt"
	"testing"
)

func TestMove_WHEN_valid_THEN_update_pos(t *testing.T) {
	p := Point2D{0, 0}
	p.Move(East)

	expectPt := Point2D{1, 0}
	if p != expectPt {
		t.Errorf("%v doesn't go East", p)
	}
	fmt.Println(p)

	p.Move(South)

	fmt.Println(p)
	if p.X != 1 || p.Y != 1 {
		t.Errorf("%v doesn't go North", p)
	}
}

func TestMove_WHEN_invalid_THEN_getErr(t *testing.T) {
	p := Point2D{0, 0}
	err := p.Move(West)
	t.Log(err)
	if err == nil {
		t.Errorf("Should throw err")
	}

	err = p.Move(North)
	t.Log(err)
	if err == nil {
		t.Errorf("Should throw err")
	}
}
