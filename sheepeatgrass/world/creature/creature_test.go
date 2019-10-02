package creature_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"zentest.io/sheepeatgrass/world"
	"zentest.io/sheepeatgrass/world/creature"
	"zentest.io/sheepeatgrass/world/geo"
	"zentest.io/sheepeatgrass/world/grass"
	"zentest.io/sheepeatgrass/world/sheep"
)

func TestCreature(t *testing.T) {
	w := &world.World{}

	var c creature.ICreature
	var err error
	c, err = sheep.NewSheep("1", geo.Point2D{0, 0}, w)

	assert.Nil(t, err)

	c.Act()

	c, err = grass.NewGrass("1", geo.Point2D{0, 1}, w)

	assert.Nil(t, err)
	c.Act()

	tt, ok := c.(*grass.Grass)
	if !ok {
		t.Errorf("%v is not Grass, but %T", tt, c)
	}

	if c.IsDead() {
		t.Errorf("Grass, should has 3 pts, but %v", c.GetAliveDays())
	}

}
