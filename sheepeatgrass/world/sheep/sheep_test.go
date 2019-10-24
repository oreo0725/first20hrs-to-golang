package sheep_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"zentest.io/sheepeatgrass/world"
	"zentest.io/sheepeatgrass/world/creature"
	"zentest.io/sheepeatgrass/world/geo"
	"zentest.io/sheepeatgrass/world/grass"
	"zentest.io/sheepeatgrass/world/sheep"
)

var w *world.World
var sheep1 *sheep.Sheep
var initSheepPos geo.Point2D

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
			expectLifePoint: sheep.SheepDefaultLifePoint + grass.GrassEnergy},
	}
	for _, tC := range testCases {
		setup()
		sheep1, _ = sheep.NewSheep("1", geo.Point2D{0, 0}, w)
		var grass1 *grass.Grass
		grass1, _ = grass.NewGrass("1", geo.Point2D{0, 1}, w)
		var food creature.IFood = grass1

		t.Run(tC.desc, func(t *testing.T) {
			fmt.Println(w)
			assert.NotNil(t, w.MAP[grass1.Pos.X][grass1.Pos.Y])

			sheep1.Eat(food)
			assert.Equal(t, tC.expectLifePoint, sheep1.LifePoint)
			assert.Nil(t, w.MAP[grass1.Pos.X][grass1.Pos.Y])
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
		sheep1, _ = sheep.NewSheep("1", geo.Point2D{0, 0}, w)
		t.Run(tC.desc, func(t *testing.T) {
			fmt.Println(w)
			assert.NotNil(t, w.MAP[sheep1.Pos.X][sheep1.Pos.Y])

			sheep1.Die()
			fmt.Println(w)
			assert.Nil(t, w.MAP[sheep1.Pos.X][sheep1.Pos.Y])
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
		sheepPos geo.Point2D
		grassPos []pos
	}{
		{
			desc:     "Breed a new Sheep WHEN it is surround by space",
			sheepPos: geo.Point2D{1, 1},
			grassPos: []pos{},
		},
		{
			desc:     "Breed a new Sheep WHEN it is at corner, finally breed",
			sheepPos: geo.Point2D{0, 0},
			grassPos: []pos{{0, 1}},
		},
		{
			desc:     "Breed a new Sheep WHEN it is sround by 3 others, finally breed",
			sheepPos: geo.Point2D{1, 1},
			grassPos: []pos{{2, 1}, {1, 0}, {1, 2}},
		},
	}
	for _, tC := range testCases {
		setup()
		t.Run(tC.desc, func(t *testing.T) {
			fmt.Printf("::::::::: Running test: %s ::::::::\n", tC.desc)

			sheep1, _ = sheep.NewSheep("1", tC.sheepPos, w)
			for i, pos := range tC.grassPos {
				grass.NewGrass(fmt.Sprintf("%d", i), geo.Point2D{pos.x, pos.y}, w)
			}

			fmt.Println(w)

			var creature creature.ICreature = sheep1
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

func TestAct(t *testing.T) {
	type pos struct {
		x, y int
	}

	testCases := []struct {
		desc           string
		sheepPos       geo.Point2D
		sheepLivedDays int
		sheepLifePoint int
		grassPos       []pos
		assertFunc     func()
	}{
		{
			desc:           "Sheep should move",
			sheepPos:       geo.Point2D{1, 1},
			sheepLivedDays: 49,
			sheepLifePoint: sheep.SheepDefaultLifePoint,
			grassPos:       []pos{},
			assertFunc: func() {
				assert.NotEqual(t, sheep1.Pos, geo.Point2D{1, 1})
			},
		},
		{
			desc:           "Sheep should die at day70",
			sheepPos:       geo.Point2D{0, 0},
			sheepLivedDays: 70,
			sheepLifePoint: sheep.SheepDefaultLifePoint,
			grassPos:       []pos{{0, 1}},
			assertFunc: func() {
				assert.Nil(t, w.MAP[sheep1.Pos.X][sheep1.Pos.Y])
			},
		},
		{
			desc:           "Sheep should die without lifePoints",
			sheepPos:       geo.Point2D{0, 0},
			sheepLivedDays: 49,
			sheepLifePoint: 0,
			grassPos:       []pos{{0, 1}},
			assertFunc: func() {
				assert.Nil(t, w.MAP[sheep1.Pos.X][sheep1.Pos.Y])
			},
		},
		{
			desc:           "Sheep should breed a baby at day50",
			sheepPos:       geo.Point2D{0, 0},
			sheepLivedDays: 49,
			sheepLifePoint: 0,
			grassPos:       []pos{{0, 1}},
			assertFunc: func() {
				assert.Nil(t, w.MAP[sheep1.Pos.X][sheep1.Pos.Y])
			},
		},
	}
	for _, tC := range testCases {
		setup()
		t.Run(tC.desc, func(t *testing.T) {
			fmt.Printf("::::::::: Running test: %s ::::::::\n", tC.desc)

			sheep1, _ = sheep.NewSheep("1", tC.sheepPos, w)
			sheep1.LifePoint = tC.sheepLifePoint
			sheep1.AliveDays = tC.sheepLivedDays
			for i, pos := range tC.grassPos {
				grass.NewGrass(fmt.Sprintf("%d", i), geo.Point2D{pos.x, pos.y}, w)
			}

			fmt.Println(w)

			var c creature.ICreature = sheep1
			c.Act()
			// THEN
			fmt.Println(w)
			tC.assertFunc()
		})
		tearDown()
	}
}

func TestNewSheepWHEN_atX0Y0_THEN_world_updated(t *testing.T) {
	setup()
	sheep1, _ = sheep.NewSheep("1", geo.Point2D{0, 0}, w)
	assert.EqualValues(t, sheep1, w.MAP[0][0])
}

func TestMove_WHEN_validPos_THEN_updatedPos(t *testing.T) {
	//GIVEN
	setup()
	sheep1, _ = sheep.NewSheep("1", geo.Point2D{0, 0}, w)
	//WHEN
	sheep1.Move(geo.South)
	//THEN
	if sheep1.Pos.Y != 1 {
		t.Errorf("Sheep should be at [0, 1], but %v", sheep1.Pos)
	}

	// fmt.Print(w)
	if w.MAP[0][1] != sheep1 {
		t.Errorf("Sheep should be at [1, 3] in the world")
	}
}

func TestMove_WHEN_invalidPos_THEN_notUpdated(t *testing.T) {
	//GIVEN
	setup()
	sheep1, _ = sheep.NewSheep("1", geo.Point2D{0, 0}, w)
	//WHEN
	err := sheep1.Move(geo.West)
	//THEN
	assert.NotNil(t, err)
	if sheep1.Pos.X != 0 {
		t.Errorf("Sheep should be at [0, 2], but %v", sheep1.Pos)
	}
}
