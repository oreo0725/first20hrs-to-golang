package world

import (
	"fmt"
)

//Sheep - a sheep
type Sheep struct {
}

func (s *Sheep) Act() {
	// c.lifePoint - DAILY_CONSUME

	fmt.Println("I am Sheep")
}
