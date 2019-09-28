package world

import (
	"testing"
)

func TestCreature(t *testing.T) {
	w := &World{}

	var creature ICreature = NewSheep("1", Point2D{}, w)

	creature.Act()

	creature = NewGrass("1", Point2D{}, w)

	creature.Act()

	tt, ok := creature.(*Grass)
	if !ok {
		t.Errorf("%v is not Grass, but %T", tt, creature)
	}

	if creature.IsDead() {
		t.Errorf("Grass, should has 3 pts, but %v", creature.GetAliveDays())
	}

}
