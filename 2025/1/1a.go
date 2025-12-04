package main

import (
	"fmt"
	"strconv"

	lib "monkadev.xyz/aoc"
)

func Day1A() {
	var input []string = lib.Read("1/input.txt")
	var location int = 50
	var zeroAmount int = 0

	for i := 0; i < len(input); i++ {
		var direction string = string(input[i][0])
		value, err := strconv.Atoi(input[i][1:])

		if err == nil {
			switch direction {
			case "R":
				location += value
			case "L":
				location -= value
			}
		}
		location %= 100
		if location == 0 {
			zeroAmount++
		}
	}

	fmt.Println(zeroAmount)
}
