package world

//Life -  An abstract class
//go:generate stringer -type=Life
type Life struct {
	pos       Point2D
	aliveDays int
	world     *World
}

//Act - How a creature act within a day
func (c *Life) Act() {
	// c.lifePoint - DAILY_CONSUME
}

//ICreature interface
type ICreature interface {
	Act()
	GetName() string
	IsDead() bool
	GetAliveDays() int
}

//IAnimal -
type IAnimal interface {
	Move(dir Direction) bool
	Eat()
}
