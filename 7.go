package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type bagContent struct {
	num   int
	color string
}

var rules = map[string]map[int]bagContent{}

func main7() {
	input := getBagRules()

	for _, rule := range input {
		first, _ := regexp.Compile("^([\\w ]+) bags contain (.*).$")
		firstColor := first.FindAllStringSubmatch(rule, -1)
		rules[firstColor[0][1]] = map[int]bagContent{}

		colors := strings.Split(firstColor[0][2], ", ")
		for i, color := range colors {
			if color == "no other bags" {
				rules[firstColor[0][1]][i] = bagContent{num: 0, color: ""}
				break
			}
			bagContents := strings.SplitN(color, " ", 2)
			num, _ := strconv.Atoi(bagContents[0])
			color := strings.TrimSuffix(bagContents[1], " bags")
			color = strings.TrimSuffix(color, " bag")

			rules[firstColor[0][1]][i] = bagContent{num: num, color: color}
		}
	}

	goldCandidates := map[string]int{}
	getOuterBags("shiny gold", &goldCandidates)
	newCandidates := len(goldCandidates)
	for newCandidates > 0 {
		newCandidates = 0
		for candidate := range goldCandidates {
			newCandidates += getOuterBags(candidate, &goldCandidates)
		}
	}

	fmt.Println("part 1: number of bag colors which can contain at least one shiny gold bag:", len(goldCandidates))

	inGoldBag := 0
	candidates := map[int]bagContent{}
	candidates = rules["shiny gold"]
	inGoldBag += getInnerBags(&candidates)

	fmt.Println("part 2: number of bags in a shiny gold bag", inGoldBag)
}

func getOuterBags(searchColor string, candidates *map[string]int) int {
	newCandidates := 0
	before := len(*candidates)
	for color, colors := range rules {
		for _, contents := range colors {
			if contents.color == searchColor && color != "shiny gold" {
				(*candidates)[color] = 0
				if len(*candidates) > before {
					newCandidates++
				}
			}
		}
	}
	return newCandidates
}

func getInnerBags(candidates *map[int]bagContent) int {
	inGoldBag := 0
	for _, content := range *candidates {
		inGoldBag += content.num
		if content.num > 0 {
			for i := 0; i < content.num; i++ {
				moreCandidates := rules[content.color]
				inGoldBag += getInnerBags(&moreCandidates)
			}
		}
	}
	return inGoldBag
}
