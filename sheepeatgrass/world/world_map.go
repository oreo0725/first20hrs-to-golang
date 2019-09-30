package world

import (
	"fmt"
	"golang.org/x/text/width"
	"log"
	"os"
)

const (
	WIDTH  = 12
	HEIGHT = 12
)

var logger = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime)

/**
 * World
 */
type World struct {
	MAP [WIDTH][HEIGHT]ICreature
	DAY int
}

func (w World) IsAcceptPos(x int, y int) bool {
	// fmt.Printf("check if accept pos (%d,%d)\n", x, y)
	if !(x >= 0 && x < WIDTH && y >= 0 && y < HEIGHT) || w.MAP[x][y] != nil {
		return false
	}
	return true
}

func (w World) GetAnEmptyNeighbour(pos Point2D) *Point2D {
	dir := RandDirection()
	for i := 0; i < 4; i++ {
		if newPos, err := pos.Move(dir); err == nil && w.IsAcceptPos(newPos.X, newPos.Y) {
			return &newPos
		}
		dir = dir.Next()
	}
	return nil
}

func (w World) String() string {
	var str = fmt.Sprintf("===== Day: %v ======\n", w.DAY)
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			creature := w.MAP[x][y]
			if creature == nil {
				str += width.Widen.String(fmt.Sprintf("%5s", "◼️"))
			} else {
				str += width.Widen.String(fmt.Sprintf("%4s", creature.GetName()))
			}
		}
		str += "\n"
	}
	return str
}

func (w *World) OnPosChanged(oldPos Point2D, newPos Point2D) {
	w.MAP[newPos.X][newPos.Y], w.MAP[oldPos.X][oldPos.Y] = w.MAP[oldPos.X][oldPos.Y], nil
	logger.Printf("position changed: \n %v", w)
}

func (w *World) OnCreatureEaten(creature ICreature, food IFood) {
	var eaten ICreature = food.(ICreature)
	pos := eaten.GetPos()
	logger.Printf("%v ate food[%v at %v]\n", creature.GetName(), eaten.GetName(), pos)
	w.MAP[pos.X][pos.Y] = nil
}

func (w *World) OnLifeDead(creature ICreature) {
	pos := creature.GetPos()
	w.MAP[pos.X][pos.Y] = nil
	logger.Printf("%v die at [%v], it lived for %d days\n", creature.GetName(), pos, creature.GetAliveDays())
}

func (w *World) OnNewLifeBorn(creature ICreature) {
	pos := creature.GetPos()
	w.MAP[pos.X][pos.Y] = creature
	logger.Printf("%v was born at [%v]\n", creature.GetName(), pos)
}
