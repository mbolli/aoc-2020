package main

import "fmt"

func main9() {
	input := getXmas()
	notValid := 0

	for i, num := range input {
		if i < 25 {
			continue
		}

		found := false

		for _, preamble := range input[i-25 : i] {
			for _, preamble2 := range input[i-25 : i] {
				if preamble+preamble2 == num {
					found = true
					goto afterLoop
				}
			}
		}

	afterLoop:
		if found == false {
			notValid = num
			break
		}
	}

	fmt.Println("part 1:", notValid)

	for i := range input {
		iTemp := i
		result := 0
		for result != notValid && iTemp < len(input)-1 {
			result += input[iTemp]
			iTemp++
		}

		if result == notValid {
			min := notValid
			max := 0
			for _, x := range input[i:iTemp] {
				if x > max {
					max = x
				}

				if x < min {
					min = x
				}
			}
			fmt.Println("part 2:", min+max)
			break
		}
	}

}
