package main

import (
	"fmt"
	"strconv"

	lib "monkadev.xyz/aoc"
)

func Day3B() {
	var input []string = lib.Read("3/input.txt")

	sum := 0

	for i := 0; i < len(input); i++ {
		number := input[i]

		var batteries [12]byte
		copy(batteries[:], []byte(input[i]))

		for j := 12; j < len(number); j++ {

			for b := 0; b < len(batteries)-1; b++ {
				if batteries[b] < batteries[b+1] {
					batteries = shiftLeftByOne(batteries, b)
					break
				}
			}

			digit := number[j]
			if digit > batteries[11] {
				batteries[11] = digit
			}
		}
		result, err := strconv.Atoi(string(batteries[:]))
		if err != nil {
			return
		}

		sum += result
	}

	fmt.Println(sum)
}

func shiftLeftByOne(batteries [12]byte, startIndex int) [12]byte {
	copy(batteries[startIndex:], batteries[startIndex+1:])
	batteries[len(batteries)-1] = 0
	return batteries
}
