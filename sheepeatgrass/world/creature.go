package world

//Life -  An abstract class
//go:generate stringer -type=Life
type Life struct {
	Pos       Point2D
	AliveDays int
	World     *World
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
	GetPos() Point2D
}

//IAnimal -
type IAnimal interface {
	Move(dir Direction) error
	Eat(food ICreature)
}

type IFood interface {
	GetEnergyPoint() int
}
