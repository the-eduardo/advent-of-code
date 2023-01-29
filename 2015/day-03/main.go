package main

import (
	"fmt"
	"os"
)

func main() {
	text, err := os.ReadFile("2015/day-03/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	input := string(text)
	// create a map to store the visited houses
	visited := make(map[string]bool)
	// start at 0,0
	santaX, santaY, roboX, roboY := 0, 0, 0, 0 // add the starting house to the map
	visited[fmt.Sprintf("%d,%d", santaX, santaY)] = true
	// loop through the input
	for i, char := range input {
		// Check if Santa or Robot is moving
		var x, y int
		if i%2 == 0 {
			x, y = santaX, santaY
		} else {
			x, y = roboX, roboY
		}
		switch char {
		case '^':
			y++
		case 'v':
			y--
		case '>':
			x++
		case '<':
			x--
		}
		// add the house to the map
		visited[fmt.Sprintf("%d,%d", x, y)] = true
		if i%2 == 0 {
			santaX, santaY = x, y
		} else {
			roboX, roboY = x, y
		}
	}
	// print the number of houses visited
	fmt.Printf("Santa and Robot visited %d houses\n", len(visited))
}
