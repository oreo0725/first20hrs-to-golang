package creature

import (
	"zentest.io/sheepeatgrass/world/geo"
)

//ICreature interface
type ICreature interface {
	Act()
	GetName() string
	IsDead() bool
	GetAliveDays() int
	GetPos() geo.Point2D
	Breed() (ICreature, error)
}

//IAnimal -
type IAnimal interface {
	ICreature
	Move(dir geo.Direction) error
	Eat(food ICreature)
}

type IFood interface {
	GetEnergyPoint() int
}
