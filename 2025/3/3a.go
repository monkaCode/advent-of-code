package main

import (
	"fmt"
	"strconv"

	lib "monkadev.xyz/aoc"
)

func Day3A() {
	var input []string = lib.Read("3/input.txt")

	sum := 0

	for i := 0; i < len(input); i++ {
		number := input[i]

		first := number[0]
		second := number[1]
		for j := 2; j < len(number); j++ {

			digit := number[j]
			if second > first {
				first = second
				second = 0
			}
			if digit > second {
				second = digit
			}
		}
		result, err := strconv.Atoi(string(first) + string(second))
		if err != nil {
			return
		}

		sum += result
	}

	fmt.Println(sum)
}
