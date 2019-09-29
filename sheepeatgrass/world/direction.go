package world

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

func (d Direction) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}

// Get a direction randomly
func RandDirection() Direction {
	// rand.Seed(time.Now().UnixNano())
	return [...]Direction{North, East, South, West}[rand.RandInt(4)]
}
