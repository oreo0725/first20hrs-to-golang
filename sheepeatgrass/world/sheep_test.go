package world

import (
	"testing"
)

func TestMove_WHEN_validPos_THEN_updatedPos(t *testing.T) {
	var w = &World{}
	sheep := NewSheep("test", Point2D{1, 2}, w)
	sheep.Move(South)

	if sheep.pos.Y != 3 {
		t.Errorf("Sheep should be at [1, 3], but %v", sheep.pos)
	}
}

func TestMove_WHEN_invalidPos_THEN_notUpdated(t *testing.T) {
	var w = &World{}
	sheep := NewSheep("test2", Point2D{0, 2}, w)
	sheep.Move(South)

	if sheep.pos.X != 0 {
		t.Errorf("Sheep should be at [0, 2], but %v", sheep.pos)
	}
}
