package main

import "fmt"

type seat struct {
	row    int
	column int
	id     int
}

func main5() {
	input := getSeats()
	var seats []seat
	for _, identifier := range input {
		seats = append(seats, findSeat(identifier))
	}

	biggestSeatId := 0
	for _, s := range seats {
		if s.id > biggestSeatId {
			biggestSeatId = s.id
		}
	}

	fmt.Println("part 1: biggest seat id:", biggestSeatId)

	var lastSeat seat
	for i := 0; i < 128; i++ {
		for j := 0; j < 8; j++ {
			for _, s := range seats {
				if s.row == i && s.column == j {
					if lastSeat.id != 0 && lastSeat.id != s.id-1 {
						fmt.Println("part 2: my seat must be", lastSeat.id+1)
					}
					lastSeat = s
				}
			}
		}
	}
}

/**
 * 128x8
 * returns seat
 */
func findSeat(identifier string) seat {
	rows := 128
	rowsUp := 127
	rowsLo := 0

	columns := 8
	columnsUp := 7
	columnsLo := 0
	index := 0
	for index < len(identifier) {
		switch string(identifier[index]) {
		case "F":
			rows /= 2
			rowsUp -= rows
		case "B":
			rows /= 2
			rowsLo += rows
		case "L":
			columns /= 2
			columnsUp -= columns
		case "R":
			columns /= 2
			columnsLo += columns
		}
		index++
	}

	return seat{rowsLo, columnsLo, rowsLo*8 + columnsLo}
}
