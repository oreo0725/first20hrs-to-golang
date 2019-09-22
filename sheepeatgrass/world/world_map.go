package world

import (
	"fmt"
	"golang.org/x/text/width"
)

const (
	WIDTH  = 12
	HEIGHT = 12
)

/**
 * World
 */
type World struct {
	MAP [WIDTH][HEIGHT]ICreature
	DAY int
}

func (w World) String() string {
	var str = fmt.Sprintf("===== Day: %v ======\n", w.DAY)
	for _, col := range w.MAP {
		for _, creature := range col {
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
