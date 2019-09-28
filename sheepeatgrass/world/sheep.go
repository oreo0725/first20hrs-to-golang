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
func NewSheep(name string, pos Point2D, world *World) *Sheep {
	return &Sheep{
		Life{
			aliveDays: 0,
			pos:       pos,
			world:     world},
		20,
		name}
}

func (s *Sheep) Act() {
	// c.lifePoint - DAILY_CONSUME
	fmt.Printf("I am %T\n", s)
}

func (s *Sheep) GetName() string {
	return "🐏" + s.name
}

func (s *Sheep) Move(dir Direction) error {
	return s.pos.Move(dir)
}

func (s *Sheep) IsDead() bool {
	return s.aliveDays >= sheepDeathDayLong || s.lifePoint <= 0
}

func (s *Sheep) GetAliveDays() int {
	return s.aliveDays
}
