package world

import (
	"fmt"
)

const (
	grassDeathDayLong = 6
	GrassEnergy       = 5
)

//Grass - a grass
//go:generate stringer -type=Grass
type Grass struct {
	Life `default:"{\"lifePoint\":3}"`
}

func (g *Grass) Act() {
	fmt.Printf("I am %T\n", g)
}

func (g *Grass) GetName() string {
	return "🌱" + g.name
}

func (g *Grass) IsDead() bool {
	return g.AliveDays >= grassDeathDayLong
}

func (g *Grass) GetAliveDays() int {
	return g.AliveDays
}

func (g *Grass) GetEnergyPoint() int {
	return GrassEnergy
}

func (g *Grass) GetPos() Point2D {
	return g.Pos
}

func (g *Grass) Breed() (ICreature, error) {
	//TODO implementation
	return nil, nil
}

/**
 * Constructor of Grass
 */
func NewGrass(name string, pos Point2D, world *World) (*Grass, error) {
	if !world.IsAcceptPos(pos.X, pos.Y) {
		return nil, fmt.Errorf("Point: %v is not empty", pos)
	}
	newGrass := &Grass{Life{pos, 0, world, 0, name}}
	world.OnNewLifeBorn(newGrass)

	return newGrass, nil
}
