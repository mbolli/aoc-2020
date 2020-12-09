package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main8() {
	input := getBootcode()
	//input := []string{"nop +0", "acc +1", "jmp +4", "acc +3", "jmp -3", "acc -99", "acc +1", "jmp -4", "acc +6"}
	acc := 0
	visits := map[int]int{}
	nops := map[int]int{}
	nopToIgnore := 0
	hasLoop := true
	replacedOp := false

	for hasLoop == true {
		acc = 0
		for i := 0; i < len(input); i++ {
			hasLoop = false
			opline := input[i]
			opsplit := strings.Split(opline, " ")
			op := opsplit[0]
			val, _ := strconv.Atoi(opsplit[1])
			if replacedOp == false && op == "jmp" {
				nops[i] = 0
				if len(nops)-1 > nopToIgnore {
					fmt.Println("replacing", op, val, "with nop")
					op = "nop"
					replacedOp = true
				}
			}
			//fmt.Println(i, op, val, "nopToIgnore", nopToIgnore, "nops", len(nops))

			visitsBefore := len(visits)
			visits[i] = 0

			if visitsBefore == len(visits) {
				fmt.Println("part 1: loop detected. accumulator is", acc)
				hasLoop = true
				break
			}

			if op == "nop" {

			} else if op == "acc" {
				acc += val
			} else if op == "jmp" {
				i += val - 1
			}
		}
		replacedOp = false
		visits = map[int]int{}
		nopToIgnore++
	}

	fmt.Println("?part 2: without loop, accumulator is", acc)
}
