package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var neededFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
var intFields = []string{"byr", "iyr", "eyr", "pid"}

func main4() {
	input := getPassports()
	validPassports := 0
	ri := 0

	for i := range input {

		if len(input[i]) == 0 {
			ri++
			continue
		}

		if ri > i {
			continue
		}

		var keys []string
		var values []string
		for len(input[ri]) != 0 && len(input)-1 > ri {
			pairs := strings.Split(input[ri], " ")
			for j := range pairs {
				pair := strings.Split(pairs[j], ":")
				keys = append(keys, pair[0])
				values = append(values, pair[1])
			}
			ri++
		}

		// check passport fields length
		if len(neededFields) > len(keys) {
			continue
		}

		// check passport fields exactly
		invalid := false
		for _, v := range neededFields {
			if contains(v, keys) == false {
				invalid = true
				fmt.Println("invalid", keys)
				break
			}
		}

		if invalid == false {
			for k, key := range keys {
				if checkValue(key, values[k]) == false {
					invalid = true
					break
				}
			}
		}

		if invalid == false {
			validPassports++
			fmt.Println("valid", keys)
		}
	}

	fmt.Println("found", validPassports, "valid passports in", len(input), "lines")
}

func checkValue(key string, value string) bool {
	fmt.Println(key, value)
	number := 0
	if contains(key, intFields) {
		parsedInt, err := strconv.Atoi(value)
		number = parsedInt
		if err != nil {
			return false
		}
	}

	switch key {
	case "byr":
		return number > 1919 && number < 2003
	case "iyr":
		return number > 2009 && number < 2021
	case "eyr":
		return number > 2019 && number < 2031
	case "hgt":
		if strings.Contains(value, "cm") {
			number, err := strconv.Atoi(value[0:3])
			return err == nil && number > 149 && number < 194
		} else if strings.Contains(value, "in") {
			number, err := strconv.Atoi(value[0:2])
			return err == nil && number > 58 && number < 77
		} else {
			return false
		}
	case "hcl":
		reg, _ := regexp.Compile("^#[0-9a-f]{6}$")
		return reg.MatchString(value)
	case "ecl":
		reg, _ := regexp.Compile("^(amb|blu|brn|gry|grn|hzl|oth)$")
		return reg.MatchString(value)
	case "pid":
		reg, _ := regexp.Compile("^[0-9]{9}$")
		return reg.MatchString(value)
	default:
		return true
	}
}
