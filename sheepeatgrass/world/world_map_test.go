package world_test

import (
	"fmt"
	"testing"
	"zentest.io/sheepeatgrass/world"
	"zentest.io/sheepeatgrass/world/creature"
	"zentest.io/sheepeatgrass/world/geo"
	"zentest.io/sheepeatgrass/world/grass"
	"zentest.io/sheepeatgrass/world/sheep"
)

func TestWorld_String(t *testing.T) {

	w := world.World{DAY: 0}
	sheep.NewSheep("1", geo.Point2D{1, 1}, &w)
	grass.NewGrass("11", geo.Point2D{2, 2}, &w)
	fmt.Println(w)
}

func TestWorld_isAcceptPos(t *testing.T) {
	type fields struct {
		MAP [world.WIDTH][world.HEIGHT]creature.ICreature
		DAY int
	}
	type args struct {
		x int
		y int
	}

	w := world.World{}
	tests := []struct {
		x    int
		y    int
		want bool
	}{
		{0, 0, true},
		{1, 1, true},
		{1, 0, true},
		{-1, 0, false},
		{-1, -1, false},
		{0, -1, false},
		{12, 0, false},
		{11, 13, false},
		{11, 0, true},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {

			if got := w.IsAcceptPos(tt.x, tt.y); got != tt.want {
				t.Errorf("World.IsAcceptPos() = %v, want %v", got, tt.want)
			}
		})
	}
}
