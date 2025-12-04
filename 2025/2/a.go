package main

import (
	"fmt"
	"strconv"
	"strings"

	lib "monkadev.xyz/aoc"
)

func Day2A() {
	var input []string = lib.Read("2/input.txt")

	var ranges []string = strings.Split(input[0], ",")

	for i := 0; i < len(ranges); i++ {
		numbers := strings.Split(ranges[i], "-")
		first, err1 := strconv.Atoi(numbers[0])
		last, err2 := strconv.Atoi(numbers[1])
		if err1 != nil || err2 != nil {
			return
		}

		fmt.Println(first, last)
	}
}
