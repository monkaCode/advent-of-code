package main

import (
	"fmt"
	"strings"

	lib "monkadev.xyz/aoc"
)

func Day6B() int {
	input := lib.Read("6/input.txt")

	amount := len(strings.Fields(input[0]))

	formulas := make([][]string, amount)

	for r := range input {
		parts := strings.Fields(input[r])
		for c := range parts {
			if formulas[c] == nil {
				formulas[c] = make([]string, 5)
			}
			if r == 4 {
				formulas[c][4] = parts[c]
			} else {
				if input[r][] {
					formulas[c][0] += string(parts[c][3])
				}
				if length >= 3 {
					formulas[c][1] += string(parts[c][2])
				}
				if length >= 2 {
					formulas[c][2] += string(parts[c][1])
				}
				if length >= 1 {
					formulas[c][3] += string(parts[c][0])
				}
			}
		}
	}

	fmt.Println(formulas)

	sum := 0

	// for f := range formulas {
	// 	formula := formulas[f]
	// 	first, err := strconv.Atoi(formula[0])
	// 	if err != nil {
	// 		return 0
	// 	}
	// 	calc := first
	// 	for s := range formula {
	// 		if s == 4 || s == 0 {
	// 			continue
	// 		}
	// 		i, err := strconv.Atoi(formula[s])
	// 		if err != nil {
	// 			return 0
	// 		}
	// 		switch formula[4] {
	// 		case "+":
	// 			calc += i
	// 		case "*":
	// 			calc *= i
	// 		}
	// 	}
	// 	sum += calc
	// }

	return sum
}
