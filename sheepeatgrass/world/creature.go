package world

//Life -  An abstract class
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
}
