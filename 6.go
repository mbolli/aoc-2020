package main

import (
	"fmt"
	"strings"
)

func main6() {
	input := getDeclarationForms()
	newGroup := false
	var groups []string
	group := ""

	// easier to parse group array
	for i, row := range input {
		if len(row) == 0 {
			newGroup = true
		} else {
			group += row + "\n"
		}

		if newGroup == true || i == len(input)-1 {
			groups = append(groups, group)
			group = ""
			newGroup = false
		}
	}

	answeredByAnyone := 0
	for _, group := range groups {
		answered := make(map[string]struct{})
		for _, char := range group {
			if string(char) == "\n" {
				continue
			}
			answered[string(char)] = struct{}{}
		}
		answeredByAnyone += len(answered)
	}

	fmt.Println("part 1:", answeredByAnyone, "questions anyone answered") // 6778

	answeredByEveryone := 0
	for _, group := range groups {
		var answered []interface{}
		persons := strings.Split(group, "\n")
		for pid, person := range persons {
			if pid == 0 {
				for _, answer := range person {
					answered = append(answered, answer)
				}
			} else if len(person) != 0 {
				answered = intersect(answered, []rune(person))
			}
		}
		answeredByEveryone += len(answered)
	}

	fmt.Println("part 2:", answeredByEveryone, "questions everybody answered")
}
