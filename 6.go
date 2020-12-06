package main

import "fmt"

func main6() {
	input := getDeclarationForms()

	newGroup := true
	answeredTotal := 0
	answered := make(map[string]struct{}) // maps have unique keys
	for i, row := range input {
		if len(row) == 0 || i == len(input)-1 {
			answeredTotal += len(answered)
			newGroup = true
			continue
		}

		if newGroup == true {
			answered = make(map[string]struct{})
			newGroup = false
		}

		for _, char := range row {
			answered[string(char)] = struct{}{}
		}
	}

	fmt.Println("part 1:", answeredTotal, "questions answered in total")
}
