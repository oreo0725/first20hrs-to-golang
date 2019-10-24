package sheep

import (
	"fmt"
	"strconv"
	"zentest.io/sheepeatgrass/world"
	"zentest.io/sheepeatgrass/world/creature"
	"zentest.io/sheepeatgrass/world/geo"
)

const (
	sheepDeathDayLong     = 70
	SheepDefaultLifePoint = 20
)

//Sheep - a sheep
type Sheep struct {
	world.Life `default:"{\"lifePoint\":5}"`
	LifePoint  int
}

//Constructor of Sheep
func NewSheep(name string, pos geo.Point2D, w *world.World) (*Sheep, error) {
	if !w.IsAcceptPos(pos.X, pos.Y) {
		return nil, fmt.Errorf("Point: %v is not empty", pos)
	}

	newSheep := &Sheep{
		world.Life{
			AliveDays:   0,
			Pos:         geo.Point2D{pos.X, pos.Y},
			World:       w,
			ChildrenNum: 0,
			Name:        name,
		},
		SheepDefaultLifePoint,
	}
	w.OnNewLifeBorn(newSheep)

	return newSheep, nil
}

func (s *Sheep) Act() {
	// c.lifePoint - DAILY_CONSUME
	//fmt.Printf("I am %v\n", s.GetName())

	if s.IsDead() {
		return

	} else if s.isDying() {
		s.Die()

	} else if s.canBreed() {
		s.Breed()

	} else {
		//should eat a grass if the next position standing a grass
		nextDir := geo.RandDirection()
		newPos, err := s.Pos.Move(nextDir)
		if err == nil {
			if iCreature := s.World.MAP[newPos.X][newPos.Y]; iCreature != nil {
				food, ok := iCreature.(creature.IFood)
				if ok {
					s.Eat(food)
				}
			}
		}

		if e := s.Move(nextDir); e != nil {
			fmt.Printf("%v Cannot move, just stand\n", s.GetName())
		}
	}
	s.AliveDays++
	s.LifePoint--
}

func (s *Sheep) canBreed() bool {
	return s.AliveDays >= 50 && s.AliveDays <= 55
}

func (s *Sheep) GetName() string {
	return "🐏" + strconv.Itoa(s.LifePoint)
}

func (s *Sheep) Move(dir geo.Direction) error {
	newPos, err := s.Pos.Move(dir)

	if err != nil || !s.World.IsAcceptPos(newPos.X, newPos.Y) {
		return fmt.Errorf("Not avail move to position: %v", newPos)
	}

	var listener creature.IWorldChangeListenser = s.World

	listener.OnPosChanged(s.Pos, newPos)

	s.Pos = newPos
	return nil
}

func (s *Sheep) IsDead() bool {
	return s.Life.IsDead
}

func (s *Sheep) isDying() bool {
	return s.AliveDays >= sheepDeathDayLong || s.LifePoint <= 0
}

func (s *Sheep) Die() {
	s.Life.IsDead = true
	var listener creature.IWorldChangeListenser = s.World
	listener.OnLifeDead(s)
}

func (s *Sheep) GetAliveDays() int {
	return s.AliveDays
}

func (s *Sheep) Breed() (creature.ICreature, error) {
	availPos := s.World.GetAnEmptyNeighbour(s.Pos)
	if availPos == nil {
		return nil, fmt.Errorf("No empty space to breed, at[%v]", s.Pos)
	}
	s.ChildrenNum++
	gen, _ := strconv.Atoi(s.Name)
	newSheep, err := NewSheep(fmt.Sprintf("%v", gen*10+s.ChildrenNum),
		geo.Point2D{availPos.X, availPos.Y},
		s.World)
	return newSheep, err
}

func (s *Sheep) GetPos() geo.Point2D {
	return s.Pos
}

func (s *Sheep) Eat(food creature.IFood) {
	s.LifePoint += food.GetEnergyPoint()
	var listener creature.IWorldChangeListenser = s.World
	listener.OnCreatureEaten(s, food)
}
