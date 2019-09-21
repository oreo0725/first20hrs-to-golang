package world

const (
	WIDTH  = 12
	HEIGHT = 12
)

/**
 * World
 */
type World struct {
	MAP [WIDTH][HEIGHT]ICreature
	DAY int
}
