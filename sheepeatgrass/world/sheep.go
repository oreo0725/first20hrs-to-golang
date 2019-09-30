package world

import (
	"fmt"
)

const (
	sheepDeathDayLong     = 70
	SheepDefaultLifePoint = 20
)

//Sheep - a sheep
type Sheep struct {
	Life      `default:"{\"lifePoint\":5}"`
	LifePoint int
}

//Constructor of Sheep
func NewSheep(name string, pos Point2D, world *World) (*Sheep, error) {
	if !world.IsAcceptPos(pos.X, pos.Y) {
		return nil, fmt.Errorf("Point: %v is not empty", pos)
	}

	newSheep := &Sheep{
		Life{
			AliveDays:   0,
			Pos:         Point2D{pos.X, pos.Y},
			World:       world,
			ChildrenNum: 0,
			name:        name,
		},
		SheepDefaultLifePoint,
	}
	world.OnNewLifeBorn(newSheep)

	return newSheep, nil
}

func (s *Sheep) Act() {
	// c.lifePoint - DAILY_CONSUME
	fmt.Printf("I am %T\n", s)
}

func (s *Sheep) GetName() string {
	return "🐏" + s.name
}

func (s *Sheep) Move(dir Direction) error {
	newPos, err := s.Pos.Move(dir)

	if !s.World.IsAcceptPos(newPos.X, newPos.Y) {
		return fmt.Errorf("Not avail move to position: %v", newPos)
	}

	var listenser IWorldChangeListenser = s.World

	listenser.OnPosChanged(s.Pos, newPos)

	s.Pos = newPos
	return err
}

func (s *Sheep) IsDead() bool {
	return s.AliveDays >= sheepDeathDayLong || s.LifePoint <= 0
}

func (s *Sheep) GetAliveDays() int {
	return s.AliveDays
}

func (s *Sheep) Breed() (ICreature, error) {
	availPos := s.World.GetAnEmptyNeighbour(s.Pos)
	if availPos == nil {
		return nil, fmt.Errorf("No empty space to breed, at[%v]", s.Pos)
	}
	s.ChildrenNum++
	newSheep, err := NewSheep(fmt.Sprintf("%v.%v", s.name, s.ChildrenNum),
		Point2D{availPos.X, availPos.Y},
		s.World)
	fmt.Printf("newBorn of [%v] at [%v]\n", s.GetName(), availPos)
	return newSheep, err
}

func (s *Sheep) GetPos() Point2D {
	return s.Pos
}

func (s *Sheep) Eat(food IFood) {
	s.LifePoint += food.GetEnergyPoint()
	var listenser IWorldChangeListenser = s.World
	listenser.OnCreatureEaten(s, food)
}
