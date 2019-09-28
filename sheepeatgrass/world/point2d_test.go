package world

import (
	"fmt"
	"testing"
)

func TestMove_WHEN_valid_THEN_update_pos(t *testing.T) {
	p := Point2D{0, 0}
	newP, _ := p.Move(East)

	expectPt := Point2D{1, 0}
	if newP != expectPt {
		t.Errorf("%v doesn't go East", p)
	}
	if p.X != 0 || p.Y != 0 {
		t.Errorf("origin P should not change, but now is %v", p)
	}
	fmt.Println(p)

	newP, _ = newP.Move(South)

	fmt.Println(newP)
	if newP.X != 1 || newP.Y != 1 {
		t.Errorf("%v doesn't go North", p)
	}
}

func TestMove_WHEN_invalid_THEN_getErr(t *testing.T) {
	p := Point2D{0, 0}
	_, err := p.Move(West)
	t.Log(err)
	if err == nil {
		t.Errorf("Should throw err")
	}

	_, err = p.Move(North)
	t.Log(err)
	if err == nil {
		t.Errorf("Should throw err")
	}
}
