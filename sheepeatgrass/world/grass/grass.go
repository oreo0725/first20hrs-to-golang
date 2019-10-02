package grass

import (
	"fmt"
	"zentest.io/sheepeatgrass/world"
	"zentest.io/sheepeatgrass/world/creature"
	"zentest.io/sheepeatgrass/world/geo"
)

const (
	grassDeathDayLong = 6
	GrassEnergy       = 5
)

//Grass - a grass
//go:generate stringer -type=Grass
type Grass struct {
	world.Life `default:"{\"lifePoint\":3}"`
}

func (g *Grass) Act() {
	fmt.Printf("I am %T\n", g)
}

func (g *Grass) GetName() string {
	return "🌱" + g.Life.Name
}

func (g *Grass) IsDead() bool {
	return g.AliveDays >= grassDeathDayLong
}

func (g *Grass) Die() {
	var listener creature.IWorldChangeListenser = g.World
	listener.OnLifeDead(g)
}

func (g *Grass) GetAliveDays() int {
	return g.AliveDays
}

func (g *Grass) GetEnergyPoint() int {
	return GrassEnergy
}

func (g *Grass) GetPos() geo.Point2D {
	return g.Pos
}

func (g *Grass) Breed() (creature.ICreature, error) {
	//TODO implementation
	return nil, nil
}

/**
 * Constructor of Grass
 */
func NewGrass(name string, pos geo.Point2D, w *world.World) (*Grass, error) {
	if !w.IsAcceptPos(pos.X, pos.Y) {
		return nil, fmt.Errorf("Point: %v is not empty", pos)
	}
	newGrass := &Grass{world.Life{pos, 0, w, 0, name}}
	w.OnNewLifeBorn(newGrass)

	return newGrass, nil
}
