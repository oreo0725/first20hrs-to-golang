package world_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"zentest.io/sheepeatgrass/world"
)

var w *world.World
var sheep *world.Sheep
var initSheepPos world.Point2D

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	setup()
	result := m.Run()
	tearDown()
	os.Exit(result)
}

func setup() {
	fmt.Println("===== Test setup ======")
	w = &world.World{}
	sheep, _ = world.NewSheep("1", world.Point2D{0, 0}, w)
}

func tearDown() {
	fmt.Println("===== Test tearDown ======")
}

func TestEat(t *testing.T) {
	testCases := []struct {
		desc            string
		expectLifePoint int
	}{
		{
			desc:            "A Sheep ate a grass",
			expectLifePoint: world.SheepDefaultLifePoint + world.GrassEnergy},
	}
	for _, tC := range testCases {
		setup()
		var grass *world.Grass
		grass, _ = world.NewGrass("1", world.Point2D{0, 1}, w)
		var food world.IFood = grass

		t.Run(tC.desc, func(t *testing.T) {
			fmt.Println(w)
			assert.NotNil(t, w.MAP[grass.Pos.X][grass.Pos.Y])

			sheep.Eat(food)
			assert.Equal(t, tC.expectLifePoint, sheep.LifePoint)
			assert.Nil(t, w.MAP[grass.Pos.X][grass.Pos.Y])
		})
		tearDown()
	}
}

func TestNewSheepWHEN_atX0Y0_THEN_world_updated(t *testing.T) {
	setup()
	assert.EqualValues(t, sheep, w.MAP[0][0])
}

func TestMove_WHEN_validPos_THEN_updatedPos(t *testing.T) {
	//GIVEN
	setup()
	//WHEN
	sheep.Move(world.South)
	//THEN
	if sheep.Pos.Y != 1 {
		t.Errorf("Sheep should be at [0, 1], but %v", sheep.Pos)
	}

	// fmt.Print(w)
	if w.MAP[0][1] != sheep {
		t.Errorf("Sheep should be at [1, 3] in the world")
	}
}

func TestMove_WHEN_invalidPos_THEN_notUpdated(t *testing.T) {
	setup()
	err := sheep.Move(world.West)
	assert.NotNil(t, err)

	if sheep.Pos.X != 0 {
		t.Errorf("Sheep should be at [0, 2], but %v", sheep.Pos)
	}
}
