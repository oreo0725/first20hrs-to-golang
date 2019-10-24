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
	IsDead      bool
}
