package world

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreature(t *testing.T) {
	w := &World{}

	var creature ICreature
	var err error
	creature, err = NewSheep("1", Point2D{0, 0}, w)

	assert.Nil(t, err)

	creature.Act()

	creature, err = NewGrass("1", Point2D{0, 1}, w)

	assert.Nil(t, err)
	creature.Act()

	tt, ok := creature.(*Grass)
	if !ok {
		t.Errorf("%v is not Grass, but %T", tt, creature)
	}

	if creature.IsDead() {
		t.Errorf("Grass, should has 3 pts, but %v", creature.GetAliveDays())
	}

}
