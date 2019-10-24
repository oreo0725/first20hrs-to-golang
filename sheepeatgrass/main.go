package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"zentest.io/sheepeatgrass/world"
	"zentest.io/sheepeatgrass/world/config"
	"zentest.io/sheepeatgrass/world/creature"
	"zentest.io/sheepeatgrass/world/geo"
	"zentest.io/sheepeatgrass/world/grass"
	"zentest.io/sheepeatgrass/world/sheep"
)

func main() {
	//newSheepCount, newGrassCount := 0, 0

	w := &world.World{}
	for newSheepCount := 0; newSheepCount < config.DefaultSheepAmount; newSheepCount++ {
		newPos := w.GetARandomAvailSpace()
		if newPos != nil {
			sheep.NewSheep(strconv.Itoa(newSheepCount), geo.Point2D{newPos.X, newPos.Y}, w)
		}
	}
	for newGrassCount := 0; newGrassCount < config.DefaultGrassAmount; newGrassCount++ {
		newPos := w.GetARandomAvailSpace()
		if newPos != nil {
			grass.NewGrass(strconv.Itoa(newGrassCount), geo.Point2D{newPos.X, newPos.Y}, w)
		}
	}

	fmt.Print(w)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("-----::::The world is running ::::--------")
	fmt.Println(" --- Press Enter to continue --- ")
	for {
		fmt.Print("_ ")
		text, _ := reader.ReadString('\n')
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		unmovedList := list.New()
		if text == "" {
			for _, row := range w.MAP {
				for _, c := range row {
					if c != nil {
						unmovedList.PushBack(c)
					}
				}
			}
		}

		for unmovedList.Len() > 0 {
			e := unmovedList.Front()
			c, ok := e.Value.(creature.ICreature)// First element
			if ok {
				//FIXME, what else if the creature is already dead of eaten by the other
				c.Act()
			}
			unmovedList.Remove(e) // Dequeue
		}
		fmt.Println(w)
		fmt.Println(unmovedList.Len())

		w.DAY++


	}


}
