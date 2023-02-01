package main

import (
	"fmt"
	"os"
)

var x1, y1, x2, y2, z1, z2 string
var stopCounter int

func main() {
	text, err := os.ReadFile("2015/day-07/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	input := string(text)
	// TODO: everything

}

func and(x, y uint16) uint16 {
	return x & y
}
func not(x uint16) uint16 {
	return ^x
}
func lshift(x uint16, y uint16) uint16 {
	return x << y
}
func rshift(x uint16, y uint16) uint16 {
	return x >> y
}
func or(x, y uint16) uint16 {
	return x | y
}
