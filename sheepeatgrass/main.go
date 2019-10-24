package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"zentest.io/sheepeatgrass/world"
	"zentest.io/sheepeatgrass/world/config"
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

		var days int
		if text == "" {
			days = 1

		} else if inputDays, err := strconv.Atoi(text); err == nil {
			fmt.Printf("Fastforward %v days.\n", inputDays)
			days = inputDays
		}
		for i := 0; i < days; i++ {
			w.PlayADay()
		}
		fmt.Println(w)
	}

}

