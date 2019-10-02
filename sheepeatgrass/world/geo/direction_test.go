package geo_test

import (
	"testing"
	"zentest.io/sheepeatgrass/world/geo"

	"zentest.io/sheepeatgrass/util/rand"

	"github.com/stretchr/testify/assert"
)

func TestDirection_String(t *testing.T) {
	tests := []struct {
		name string
		d    geo.Direction
		want string
	}{
		{"North", geo.North, "North"},
		{"South", geo.South, "South"},
		{"East", geo.East, "East"},
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
	d := geo.RandDirection()
	t.Log(d)
	assert.Equal(t, geo.South, d)
}
