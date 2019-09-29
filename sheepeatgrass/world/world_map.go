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

func (w World) isAcceptPos(x int, y int) bool {
	if !(x >= 0 && x < WIDTH && y >= 0 && y < HEIGHT) {
		return false
	}
	//TODO check if empty space
	return true
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

func (w *World) onPosChanged(oldPos Point2D, newPos Point2D) {
	w.MAP[newPos.X][newPos.Y], w.MAP[oldPos.X][oldPos.Y] = w.MAP[oldPos.X][oldPos.Y], nil
	logger.Printf("position changed: \n %v", w)
}
