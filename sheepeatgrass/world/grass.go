package world

import (
	"fmt"
)

//Grass - a grass
type Grass struct {
}

func (g *Grass) Act() {
	fmt.Println("I am Grass")
}
