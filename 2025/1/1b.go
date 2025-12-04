package main

import (
	"fmt"
	"math"
	"strconv"

	lib "monkadev.xyz/aoc"
)

func Day1B() {
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

		var times int = int(math.Floor(float64(location / 100)))
		if times < 0 {
			times *= -1
		}

		zeroAmount += times
		if location <= 0 {
			zeroAmount++
		}

		location %= 100
		if location < 0 {
			location += 100
		}

	}
	fmt.Println(zeroAmount)
}
