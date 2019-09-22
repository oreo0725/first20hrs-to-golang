package world

import (
	"fmt"
)

//Sheep - a sheep
type Sheep struct {
	Life `default:"{\"lifePoint\":5}"`
	name string
}

//Constructor of Sheep
func NewSheep(name string, pos Point2D) *Sheep {
	return &Sheep{
		Life{lifePoint: 5, pos: pos}, name}
}

func (s *Sheep) Act() {
	// c.lifePoint - DAILY_CONSUME
	fmt.Printf("I am %T\n", s)
}

func (s *Sheep) GetHealth() int {
	return s.lifePoint
}

func (s *Sheep) GetName() string {
	return "🐏" + s.name
}

func (s *Sheep) Move(dir Direction) error {
	return s.pos.Move(dir)
}
