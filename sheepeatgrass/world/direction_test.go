package world_test

import (
	"testing"

	"zentest.io/sheepeatgrass/util/rand"

	"github.com/stretchr/testify/assert"
	"zentest.io/sheepeatgrass/world"
)

func TestDirection_String(t *testing.T) {
	tests := []struct {
		name string
		d    world.Direction
		want string
	}{
		{"North", world.North, "North"},
		{"South", world.South, "South"},
		{"East", world.East, "East"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.String(); got != tt.want {
				t.Errorf("Direction.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandDirection(t *testing.T) {
	rand.FixRandInt(2)
	d := world.RandDirection()
	t.Log(d)
	assert.Equal(t, world.South, d)
}
