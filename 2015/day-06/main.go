package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// make a grip 1000x100
	grip := make([][]bool, 1000)
	gripV2 := make([][]uint, 1000)
	for i := range grip {
		grip[i] = make([]bool, 1000)
		gripV2[i] = make([]uint, 1000)
	}

	text, err := os.ReadFile("2015/day-06/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	input := string(text)

	var turnOn, turnOff, toggle int

	for _, line := range strings.Split(input, "\n") {

		var x1, y1, x2, y2 int
		if strings.Contains(line, "turn on") {
			turnOn++
			fmt.Sscanf(line, "turn on %d,%d through %d,%d", &x1, &y1, &x2, &y2)
			//fmt.Println(x1, y1, x2, y2)
			turnOnLights(grip, gripV2, x1, y1, x2, y2)
		} else if strings.Contains(line, "turn off") {
			turnOff++
			fmt.Sscanf(line, "turn off %d,%d through %d,%d", &x1, &y1, &x2, &y2)
			//fmt.Println(x1, y1, x2, y2)
			turnOffLights(grip, gripV2, x1, y1, x2, y2)
		} else if strings.Contains(line, "toggle") {
			toggle++
			fmt.Sscanf(line, "toggle %d,%d through %d,%d", &x1, &y1, &x2, &y2)
			//fmt.Println(x1, y1, x2, y2)
			togleLights(grip, gripV2, x1, y1, x2, y2)
		}
	}
	fmt.Println("turn on:\t", turnOn)
	fmt.Println("turn off:\t", turnOff)
	fmt.Println("toggle:\t\t", toggle)

	// check how many Bools are true
	var counter int
	for i := range grip {
		for j := range grip[i] {
			if grip[i][j] {
				counter++
			}
		}
	}

	fmt.Println("Total lights on:\t", counter)
	// part 2:
	var counter2 uint
	for i := range gripV2 {
		for j := range gripV2[i] {
			counter2 += gripV2[i][j]
		}
	}
	fmt.Println("Total brightness:\t", counter2)
}

func togleLights(grip [][]bool, gripV2 [][]uint, x1, y1, x2, y2 int) {
	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			grip[i][j] = !grip[i][j]
			gripV2[i][j] += 2
		}
	}
}
func turnOnLights(grip [][]bool, gripV2 [][]uint, x1, y1, x2, y2 int) {
	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			grip[i][j] = true
			gripV2[i][j]++
		}
	}
}

func turnOffLights(grip [][]bool, gripV2 [][]uint, x1, y1, x2, y2 int) {
	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			grip[i][j] = false
			if gripV2[i][j] > 0 {
				gripV2[i][j]--
			}
		}
	}
}
