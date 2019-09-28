package world

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSheepWHEN_atX0Y0_THEN_world_updated(t *testing.T) {
	var w = &World{}
	sheep, err := NewSheep("test", Point2D{0, 0}, w)

	assert.Nil(t, err)

	assert.EqualValues(t, sheep, w.MAP[0][0])
}

func TestMove_WHEN_validPos_THEN_updatedPos(t *testing.T) {
	var w = &World{}
	sheep, _ := NewSheep("1", Point2D{1, 2}, w)

	sheep.Move(South)

	if sheep.pos.Y != 3 {
		t.Errorf("Sheep should be at [1, 3], but %v", sheep.pos)
	}

	fmt.Print(w)
	if w.MAP[1][3] != sheep {
		t.Errorf("Sheep should be at [1, 3] in the world")
	}
}

func TestMove_WHEN_invalidPos_THEN_notUpdated(t *testing.T) {
	var w = &World{}
	sheep, _ := NewSheep("test2", Point2D{0, 2}, w)

	sheep.Move(South)

	if sheep.pos.X != 0 {
		t.Errorf("Sheep should be at [0, 2], but %v", sheep.pos)
	}
}
