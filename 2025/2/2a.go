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

	sum := 0

	for i := 0; i < len(ranges); i++ {
		numbers := strings.Split(ranges[i], "-")
		first, err1 := strconv.Atoi(numbers[0])
		last, err2 := strconv.Atoi(numbers[1])
		if err1 != nil || err2 != nil {
			return
		}

		for number := first; number <= last; number++ {
			if isNumberInvalidSimple(number) {
				sum += number
				// fmt.Println(first, number, last)
			}
		}
	}

	fmt.Println(sum)
}

func isNumberInvalidSimple(number int) bool {
	length := len(strconv.Itoa(number))
	if length%2 == 1 {
		return false
	}

	pointer1 := 0
	pointer2 := length / 2

	for ; pointer2 < length; pointer2++ {
		if strconv.Itoa(number)[pointer1] != strconv.Itoa(number)[pointer2] {
			return false
		}
		pointer1++
	}
	return true
}
