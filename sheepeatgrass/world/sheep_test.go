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
		sheep, _ = world.NewSheep("1", world.Point2D{0, 0}, w)
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

func TestDie(t *testing.T) {
	testCases := []struct {
		desc            string
		expectLifePoint int
	}{
		{
			desc: "A Sheep is dead, then remove it in the World",
		},
	}
	for _, tC := range testCases {
		setup()
		sheep, _ = world.NewSheep("1", world.Point2D{0, 0}, w)
		t.Run(tC.desc, func(t *testing.T) {
			fmt.Println(w)
			assert.NotNil(t, w.MAP[sheep.Pos.X][sheep.Pos.Y])

			sheep.Die()
			fmt.Println(w)
			assert.Nil(t, w.MAP[sheep.Pos.X][sheep.Pos.Y])
		})
		tearDown()
	}
}

func TestBreed(t *testing.T) {
	type pos struct {
		x, y int
	}

	testCases := []struct {
		desc     string
		sheepPos world.Point2D
		grassPos []pos
	}{
		{
			desc:     "Breed a new Sheep WHEN it is surround by space",
			sheepPos: world.Point2D{1, 1},
			grassPos: []pos{},
		},
		{
			desc:     "Breed a new Sheep WHEN it is at corner, finally breed",
			sheepPos: world.Point2D{0, 0},
			grassPos: []pos{{0, 1}},
		},
		{
			desc:     "Breed a new Sheep WHEN it is sround by 3 others, finally breed",
			sheepPos: world.Point2D{1, 1},
			grassPos: []pos{{2, 1}, {1, 0}, {1, 2}},
		},
	}
	for _, tC := range testCases {
		setup()
		t.Run(tC.desc, func(t *testing.T) {
			fmt.Printf("::::::::: Running test: %s ::::::::\n", tC.desc)

			sheep, _ = world.NewSheep("1", tC.sheepPos, w)
			for i, pos := range tC.grassPos {
				world.NewGrass(fmt.Sprintf("%d", i), world.Point2D{pos.x, pos.y}, w)
			}

			fmt.Println(w)

			var creature world.ICreature = sheep
			newBorn, err := creature.Breed()

			assert.Nil(t, err)
			assert.NotNil(t, newBorn)

			fmt.Println(w)

			posNewBorn := newBorn.GetPos()
			assert.EqualValues(t, newBorn, w.MAP[posNewBorn.X][posNewBorn.Y])

		})
		tearDown()
	}
}

func TestNewSheepWHEN_atX0Y0_THEN_world_updated(t *testing.T) {
	setup()
	sheep, _ = world.NewSheep("1", world.Point2D{0, 0}, w)
	assert.EqualValues(t, sheep, w.MAP[0][0])
}

func TestMove_WHEN_validPos_THEN_updatedPos(t *testing.T) {
	//GIVEN
	setup()
	sheep, _ = world.NewSheep("1", world.Point2D{0, 0}, w)
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
	//GIVEN
	setup()
	sheep, _ = world.NewSheep("1", world.Point2D{0, 0}, w)
	//WHEN
	err := sheep.Move(world.West)
	//THEN
	assert.NotNil(t, err)
	if sheep.Pos.X != 0 {
		t.Errorf("Sheep should be at [0, 2], but %v", sheep.Pos)
	}
}
