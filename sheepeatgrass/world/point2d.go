package world

import (
	"fmt"
)

/**
 * Point2D
 */
type Point2D struct {
	X int
	Y int
}

func (p *Point2D) Move(dir Direction) error {
	newX, newY := p.X, p.Y
	switch dir {
	case East:
		newX = p.X + 1
	case West:
		newX = p.X - 1
	case North:
		newY = p.Y - 1
	case South:
		newY = p.Y + 1
	}
	if newX < 0 || newY < 0 {
		return fmt.Errorf("Invalid Point: [%v,%v]", newX, newY)
	}
	p.X, p.Y = newX, newY
	return nil
}
