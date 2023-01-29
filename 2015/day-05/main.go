package main

import (
	"fmt"
	"os"
	"strings"
)

var totalStrings int
var totalPart2 int

func main() {
	text, err := os.ReadFile("2015/day-05/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	input := string(text)
	partOne(input)
	fmt.Println("Total strings:\t", totalStrings)
	counter := 0
	for _, line := range strings.Split(input, "\n") {
		if partTwo(line) {
			counter++
		}
	}
	fmt.Println("Part 2 Total:\t", totalPart2)
	fmt.Println("Part 2 Counter:\t", counter)

}

func partOne(input string) {
	for _, line := range strings.Split(input, "\n") {
		// check if contains at least three vowels (aeiou)
		var vowels int
		for _, char := range line {
			if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
				vowels++
			}
		}
		if vowels >= 3 {
			if !strings.Contains(line, "ab") && !strings.Contains(line, "cd") && !strings.Contains(line, "pq") && !strings.Contains(line, "xy") {
				var doubleLetter bool
				for i := 0; i < len(line)-1; i++ {
					if line[i] == line[i+1] {
						doubleLetter = true
						if doubleLetter {
							//fmt.Println(line)
							totalStrings++
							break
						}
					}
				}
			}
		} else {
			continue
		}

	}
}

func partTwo(input string) bool {
	for _, line := range strings.Split(input, "\n") {
		var pair bool

		for i := 0; i < len(line)-1; i++ {
			if strings.Contains(line[i+2:], line[i:i+2]) {
				pair = true
				break
			}
		}

		var repeat bool
		for i := 0; i < len(line)-2; i++ {
			if line[i] == line[i+2] {
				repeat = true
				break
			}
		}
		if pair && repeat {
			totalPart2++
			return true
		}
	}
	return false
}
