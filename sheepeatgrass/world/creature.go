package world

//Life -  An abstract class
//go:generate stringer -type=Life
type Life struct {
	lifePoint int
	pos       Point2D
}

//Act - How a creature act within a day
func (c *Life) Act() {
	// c.lifePoint - DAILY_CONSUME
}

//ICreature interface
type ICreature interface {
	Act()
	GetHealth() int
	GetName() string
}

//IAnimal -
type IAnimal interface {
	Move(dir Direction) bool
}
