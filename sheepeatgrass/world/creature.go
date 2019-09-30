package world

//Life -  An abstract class
//go:generate stringer -type=Life
type Life struct {
	Pos         Point2D
	AliveDays   int
	World       *World
	ChildrenNum int
	name        string
}

//Act - How a creature act within a day
func (c *Life) Act() {
	// c.lifePoint - DAILY_CONSUME
}

func (s *Life) Die() {

}

func (s *Life) Breed() (ICreature, error) {
	return nil, nil
}

func (s *Life) GetAliveDays() int {
	return s.AliveDays
}

func (s *Life) GetName() string {
	return s.name
}

func (s *Life) GetPos() Point2D {
	return s.Pos
}

func (s *Life) IsDead() bool {
	return false
}

//ICreature interface
type ICreature interface {
	Act()
	GetName() string
	IsDead() bool
	GetAliveDays() int
	GetPos() Point2D
	Breed() (ICreature, error)
}

//IAnimal -
type IAnimal interface {
	Move(dir Direction) error
	Eat(food ICreature)
}

type IFood interface {
	GetEnergyPoint() int
}
