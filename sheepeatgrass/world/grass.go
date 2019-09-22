package world

import (
	"fmt"
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

func (g *Grass) GetHealth() int {
	return g.lifePoint
}

func (g *Grass) GetName() string {
	return "🌱" + g.name
}

/**
 * Constructor of Grass
 */
func NewGrass(name string, pos Point2D) *Grass {
	return &Grass{Life{lifePoint: 3, pos: pos}, name}
}
