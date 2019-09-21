package world

import (
	"testing"
)

func TestCreature(t *testing.T) {

	var creature ICreature = &Sheep{}

	creature.Act()

	creature = &Grass{}

	creature.Act()

	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		if got := tt.d.String(); got != tt.want {
	// 			t.Errorf("Direction.String() = %v, want %v", got, tt.want)
	// 		}
	// 	})
	// }
}
