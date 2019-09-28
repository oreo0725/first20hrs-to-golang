package world

import (
	"fmt"
)

const (
	grassDeathDayLong = 6
)

//Grass - a grass
//go:generate stringer -type=Grass
type Grass struct {
	Life `default:"{\"lifePoint\":3}"`
	name string
}

func (g *Grass) Act() {
	fmt.Printf("I am %T\n", g)
}

func (g *Grass) GetName() string {
	return "🌱" + g.name
}

func (g *Grass) IsDead() bool {
	return g.aliveDays >= grassDeathDayLong
}

func (g *Grass) GetAliveDays() int {
	return g.aliveDays
}

/**
 * Constructor of Grass
 */
func NewGrass(name string, pos Point2D, world *World) (*Grass, error) {
	if world.MAP[pos.X][pos.Y] != nil {
		return nil, fmt.Errorf("Point: %v is not empty", pos)
	}

	newGrass := &Grass{
		Life{
			aliveDays: 0,
			pos:       pos,
			world:     world},
		name}

	world.MAP[pos.X][pos.Y] = newGrass

	return newGrass, nil
}
