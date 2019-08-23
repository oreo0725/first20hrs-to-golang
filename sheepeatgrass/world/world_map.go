package world

import (
	"zentest.io/sheepeatgrass/creature"
)

const (
	WIDTH  = 12
	HEIGHT = 12
)

type World struct {
	MAP [WIDTH][HEIGHT]creature.Creature
	DAY int
}
