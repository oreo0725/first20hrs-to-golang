package world

import (
	"zentest.io/sheepeatgrass/world/geo"
)

//Life -  An abstract class
//go:generate stringer -type=Life
type Life struct {
	Pos         geo.Point2D
	AliveDays   int
	World       *World
	ChildrenNum int
	Name        string
}


func (l *Life) Die() {
	//TODO
}

func (l *Life) GetAliveDays() int {
	return l.AliveDays
}

func (l *Life) GetName() string {
	return l.Name
}

func (l *Life) GetPos() geo.Point2D {
	return l.Pos
}

func (l *Life) IsDead() bool {
	return false
}
