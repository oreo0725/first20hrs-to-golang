package world

import (
	"testing"
)

func Test_move_WHEN_validPos_THEN_updatedPos(t *testing.T) {
	sheep := NewSheep("test", Point2D{1, 2})
	sheep.Move(South)

	if sheep.pos.Y != 3 {
		t.Errorf("Sheep should be at [1, 3], but %v", sheep.pos)
	}
}

func Test_move_WHEN_invalidPos_THEN_notUpdated(t *testing.T) {
	sheep := NewSheep("test1", Point2D{0, 2})
	sheep.Move(South)

	if sheep.pos.X != 0 {
		t.Errorf("Sheep should be at [0, 2], but %v", sheep.pos)
	}
}
