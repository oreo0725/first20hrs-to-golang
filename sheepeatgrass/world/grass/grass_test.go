package grass_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"zentest.io/sheepeatgrass/world"
	"zentest.io/sheepeatgrass/world/creature"
	"zentest.io/sheepeatgrass/world/geo"
	"zentest.io/sheepeatgrass/world/grass"
	"zentest.io/sheepeatgrass/world/sheep"
)

var w *world.World
var grass1 *grass.Grass

func setup() {
	fmt.Println("===== Test setup ======")
	w = &world.World{}
}

func tearDown() {
	fmt.Println("===== Test tearDown ======")
}

func TestGrass_Breed(t *testing.T) {
	type pos struct {
		x, y int
	}

	testCases := []struct {
		desc     string
		grassPos geo.Point2D
		sheepPos []pos
	}{
		{
			desc:     "Breed a new grass WHEN empty around",
			grassPos: geo.Point2D{1, 1},
			sheepPos: []pos{},
		},
		//{
		//	desc:     "Breed a new Sheep WHEN it is at corner, finally breed",
		//	sheepPos: world.Point2D{0, 0},
		//	grassPos: []pos{{0, 1}},
		//},
		//{
		//	desc:     "Breed a new Sheep WHEN it is sround by 3 others, finally breed",
		//	sheepPos: world.Point2D{1, 1},
		//	grassPos: []pos{{2, 1}, {1, 0}, {1, 2}},
		//},
	}
	for _, tC := range testCases {
		setup()
		t.Run(tC.desc, func(t *testing.T) {
			fmt.Printf("::::::::: Running test: %s ::::::::\n", tC.desc)

			grass1, _ = grass.NewGrass("1", tC.grassPos, w)
			for i, pos := range tC.sheepPos {
				sheep.NewSheep(fmt.Sprintf("%d", i), geo.Point2D{pos.x, pos.y}, w)
			}

			fmt.Println(w)

			var c creature.ICreature = grass1
			newBorn, err := c.Breed()

			assert.Nil(t, err)
			assert.NotNil(t, newBorn)

			fmt.Println(w)

			posNewBorn := newBorn.GetPos()
			assert.EqualValues(t, newBorn, w.MAP[posNewBorn.X][posNewBorn.Y])

		})
		tearDown()
	}
}
