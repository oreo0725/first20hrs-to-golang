package geo

import (
	"zentest.io/sheepeatgrass/util/rand"
)

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

var Directions = []Direction{
	North, East, South, West,
}

func (d Direction) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}

func (d Direction) Next() Direction {
	return Directions[(d+1)%4]
}

// Get a direction randomly
func RandDirection() Direction {
	// rand.Seed(time.Now().UnixNano())
	return Directions[rand.RandInt(4)]
}
