package main

import "fmt"

func main1() {

	input := getReport()

	// part 1

	for i := range input {
		summand1 := input[i]

		for j := range input {
			summand2 := input[j]
			if summand1+summand2 == 2020 {
				fmt.Println("part 1: ", summand1, "*", summand2, "=", summand1*summand2)
				goto part2
			}
		}
	}

	// part 2
part2:
	for i := range input {
		summand1 := input[i]

		for j := range input {
			summand2 := input[j]

			for k := range input {
				summand3 := input[k]

				if summand1+summand2+summand3 == 2020 {
					fmt.Println("part 2: ", summand1*summand2*summand3)
					goto end
				}
			}
		}
	}

end:
}
