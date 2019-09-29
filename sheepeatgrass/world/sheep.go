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
	name      string
}

//Constructor of Sheep
func NewSheep(name string, pos Point2D, world *World) (*Sheep, error) {
	if world.MAP[pos.X][pos.Y] != nil {
		return nil, fmt.Errorf("Point: %v is not empty", pos)
	}

	newSheep := &Sheep{
		Life{
			AliveDays: 0,
			Pos:       pos,
			World:     world},
		SheepDefaultLifePoint,
		name}
	world.MAP[pos.X][pos.Y] = newSheep

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

	if !s.World.isAcceptPos(newPos.X, newPos.Y) {
		return fmt.Errorf("Not avail move to position: %v", newPos)
	}

	var listenser IWorldChangeListenser = s.World

	listenser.onPosChanged(s.Pos, newPos)

	s.Pos = newPos
	return err
}

func (s *Sheep) IsDead() bool {
	return s.AliveDays >= sheepDeathDayLong || s.LifePoint <= 0
}

func (s *Sheep) GetAliveDays() int {
	return s.AliveDays
}

func (s *Sheep) GetPos() Point2D {
	return s.Pos
}

func (s *Sheep) Eat(food IFood) {
	//TODO implement
	s.LifePoint += food.GetEnergyPoint()
	var listenser IWorldChangeListenser = s.World
	listenser.onCreatureEaten(s, food)
}
