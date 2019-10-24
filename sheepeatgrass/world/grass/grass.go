package grass

import (
	"fmt"
	"strconv"
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
	//fmt.Printf("I am %v\n", g.GetName())

	if g.IsDead() {
		return
	} else if g.isDying() {
		g.Die()
	} else if g.canBreed() {
		g.Breed()
	}
	g.AliveDays++
}

func (g *Grass) canBreed() bool {
	return g.AliveDays >= 3 && g.AliveDays <= 5
}

func (g *Grass) isDying() bool {
	return g.AliveDays >= grassDeathDayLong
}

func (g *Grass) GetName() string {
	return "🌱" //+ g.Life.Name
}

func (g *Grass) IsDead() bool {
	return g.Life.IsDead
}

func (g *Grass) Die() {
	g.Life.IsDead = true
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
	availPos := g.World.GetAnEmptyNeighbour(g.Pos)
	if availPos == nil {
		return nil, fmt.Errorf("No empty space to breed, at[%v]", g.Pos)
	}
	g.ChildrenNum++
	gen, _ := strconv.Atoi(g.Name)
	newSheep, err := NewGrass(fmt.Sprintf("%v", gen*10+g.ChildrenNum),
		geo.Point2D{availPos.X, availPos.Y},
		g.World)
	return newSheep, err
}

/**
 * Constructor of Grass
 */
func NewGrass(name string, pos geo.Point2D, w *world.World) (*Grass, error) {
	if !w.IsAcceptPos(pos.X, pos.Y) {
		return nil, fmt.Errorf("Point: %v is not empty", pos)
	}
	newGrass := &Grass{world.Life{Pos: pos, World: w, Name: name}}
	w.OnNewLifeBorn(newGrass)

	return newGrass, nil
}
