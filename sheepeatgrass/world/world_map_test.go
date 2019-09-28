package world

import (
	"fmt"
	"testing"
)

func TestWorld_String(t *testing.T) {

	w := World{DAY: 0}
	NewSheep("1", Point2D{1, 1}, &w)
	NewGrass("11", Point2D{2, 2}, &w)
	fmt.Println(w)
}

func TestWorld_isAcceptPos(t *testing.T) {
	type fields struct {
		MAP [WIDTH][HEIGHT]ICreature
		DAY int
	}
	type args struct {
		x int
		y int
	}

	w := World{}
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

			if got := w.isAcceptPos(tt.x, tt.y); got != tt.want {
				t.Errorf("World.isAcceptPos() = %v, want %v", got, tt.want)
			}
		})
	}
}
