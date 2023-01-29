package main

import (
	"fmt"
	"os"
	"strings"
)

var totalRibbon int

func main() {

	text, err := os.ReadFile("2015/day-02/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	input := string(text)
	var totalSqFt int

	for _, line := range strings.Split(input, "\n") {
		var length, width, height int
		_, err := fmt.Sscanf(line, "%dx%dx%d", &length, &width, &height)
		if err != nil {
			panic(err)
		}

		totalSqFt += length * width * 2
		totalSqFt += length * height * 2
		totalSqFt += height * width * 2
		// find the smallest side
		smallestSide := length * width
		if width*height < smallestSide {
			smallestSide = width * height
		}
		if height*length < smallestSide {
			smallestSide = height * length
		}
		totalSqFt += smallestSide
		// verify the part2
		part2(length, width, height)

	}

	// print total wrapping paper needed
	fmt.Printf("Total wrapping paper needed: %d square feet\n", totalSqFt)
	// print total ribbon needed
	fmt.Printf("Total ribbon needed: %d feet\n", totalRibbon)

}
func part2(length, width, height int) {
	// find the shortest distance around its sides, or the smallest perimeter of any one face
	var smallestPerimeter int
	smallestPerimeter = length*2 + width*2
	if width*2+height*2 < smallestPerimeter {
		smallestPerimeter = width*2 + height*2
	}
	if height*2+length*2 < smallestPerimeter {
		smallestPerimeter = height*2 + length*2
	}
	// find the cubic feet of volume of the present
	volume := length * width * height
	// add the cubic feet of volume
	totalRibbon += smallestPerimeter + volume
}
