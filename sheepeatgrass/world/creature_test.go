package world

import (
	"testing"
)

func TestCreature(t *testing.T) {

	var creature ICreature = NewSheep("1", Point2D{})

	creature.Act()

	creature = NewGrass("1", Point2D{})

	creature.Act()

	tt, ok := creature.(*Grass)
	if !ok {
		t.Errorf("%v is not Grass, but %T", tt, creature)
	}

	if creature.GetHealth() != 3 {
		t.Errorf("Grass, should has 3 pts, but %v", creature.GetHealth())
	}

	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		if got := tt.d.String(); got != tt.want {
	// 			t.Errorf("Direction.String() = %v, want %v", got, tt.want)
	// 		}
	// 	})
	// }
}
