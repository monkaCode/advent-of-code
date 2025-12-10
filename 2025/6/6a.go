package main

import (
	"strconv"
	"strings"

	lib "monkadev.xyz/aoc"
)

func Day6A() int {
	input := lib.Read("6/input.txt")

	amount := len(strings.Fields(input[0]))

	formulas := make([][]string, amount)

	for r := range input {
		parts := strings.Fields(input[r])
		for c := range parts {
			if formulas[c] == nil {
				formulas[c] = make([]string, len(input))
			}
			formulas[c][r] = parts[c]
		}
	}

	sum := 0

	for f := range formulas {
		formula := formulas[f]
		first, err := strconv.Atoi(formula[0])
		if err != nil {
			return 0
		}
		calc := first
		for s := range formula {
			if s == 4 || s == 0 {
				continue
			}
			i, err := strconv.Atoi(formula[s])
			if err != nil {
				return 0
			}
			switch formula[4] {
			case "+":
				calc += i
			case "*":
				calc *= i
			}
		}
		sum += calc
	}

	return sum
}
