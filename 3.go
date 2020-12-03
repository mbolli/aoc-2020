package main

import (
	"fmt"
	"strings"
)

var input []string

func main3() {
	input = getTreeMap()

	r1 := rideSlope(1, 1)
	r2 := rideSlope(3, 1)
	r3 := rideSlope(5, 1)
	r4 := rideSlope(7, 1)
	r5 := rideSlope(1, 2)

	fmt.Println("result:", r1*r2*r3*r4*r5)
}

func rideSlope(right int, down int) int {
	width := len(input[0])
	height := len(input)
	maxWidth := height * right
	repetitions := maxWidth / width
	xPos := -right
	treesEncountered := 0

	for i := range input {
		if i%down != 0 {
			continue
		}

		line := strings.Repeat(input[i], repetitions+1)
		xPos += right

		if xPos < maxWidth {
			if string(line[xPos]) == "#" {
				treesEncountered++
			}
		}
	}

	fmt.Println("encountered", treesEncountered, "trees")
	return treesEncountered
}
