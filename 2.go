package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main2() {

	input := getPasswords()

	// part 1
	valid := 0

	for i := range input {

		min, max, char, pass := split(input[i])
		rxp, _ := regexp.Compile(char)
		matches := rxp.FindAllString(pass, -1)

		if len(matches) >= min && len(matches) <= max {
			valid++
		}
	}

	fmt.Println("part 1:", valid, "valid passwords of", len(input))

	// part 2
	valid = 0
	for i := range input {

		pos1, pos2, char, pass := split(input[i])
		rxp1, _ := regexp.Compile("^.{" + strconv.Itoa(pos1-1) + "}" + char + ".{" + strconv.Itoa(pos2-pos1-1) + "}[^" + char + "]")
		rxp2, _ := regexp.Compile("^.{" + strconv.Itoa(pos1-1) + "}[^" + char + "].{" + strconv.Itoa(pos2-pos1-1) + "}" + char)

		if rxp1.MatchString(pass) != rxp2.MatchString(pass) {
			valid++
		}
	}

	fmt.Println("part 2:", valid, "valid passwords of", len(input))

}

func split(input string) (int, int, string, string) {
	line := strings.Split(input, ": ")
	rule := strings.Split(line[0], " ")
	num := strings.Split(rule[0], "-")
	min, _ := strconv.Atoi(num[0])
	max, _ := strconv.Atoi(num[1])
	char := rule[1]
	pass := line[1]

	return min, max, char, pass
}
