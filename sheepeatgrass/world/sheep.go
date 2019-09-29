package world

import (
	"fmt"
)

const (
	sheepDeathDayLong = 70
)

//Sheep - a sheep
type Sheep struct {
	Life      `default:"{\"lifePoint\":5}"`
	lifePoint int
	name      string
}

//Constructor of Sheep
func NewSheep(name string, pos Point2D, world *World) (*Sheep, error) {
	if world.MAP[pos.X][pos.Y] != nil {
		return nil, fmt.Errorf("Point: %v is not empty", pos)
	}

	newSheep := &Sheep{
		Life{
			aliveDays: 0,
			pos:       pos,
			world:     world},
		20,
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
	newPos, err := s.pos.Move(dir)

	if !s.world.isAcceptPos(newPos.X, newPos.Y) {
		return fmt.Errorf("Not avail move to position: %v", newPos)
	}

	var listenser IPosChangeListenser = s.world

	listenser.onPosChanged(s.pos, newPos)

	s.pos = newPos
	return err
}

func (s *Sheep) IsDead() bool {
	return s.aliveDays >= sheepDeathDayLong || s.lifePoint <= 0
}

func (s *Sheep) GetAliveDays() int {
	return s.aliveDays
}

func (s *Sheep) Eat() {
	//TODO implement
}
