package main

import (
	"fmt"
	"strconv"
	"strings"

	lib "monkadev.xyz/aoc"
)

func Day2B() {
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
			if isNumberInvalid(number) {
				sum += number
				// fmt.Println(first, number, last)
			}
		}
	}

	fmt.Println(sum)
}

func isNumberInvalid(number int) bool {
	length := len(strconv.Itoa(number))
	if length == 1 {
		return false
	}

	primes := [8]int{2, 3, 5, 7, 11, 13, 17}
	streak := map[int]struct{}{}

	for i := 0; i < len(primes); i++ {
		div := float64(length) / float64(primes[i])
		if div == float64(int64(div)) {
			streak[int(div)] = struct{}{}
			streak[primes[i]] = struct{}{}
		}
	}

	streak[1] = struct{}{}

	correct := 0
	pointer1 := 0
	pointer2 := 1

	for ; pointer2 < length; pointer2++ {
		if number == -1 {
			fmt.Println(pointer1, strconv.Itoa(number)[pointer1], pointer2, strconv.Itoa(number)[pointer2], correct)
		}
		if strconv.Itoa(number)[pointer1] == strconv.Itoa(number)[pointer2] {
			correct++
			pointer1++
		} else {
			if correct > 0 {
				pointer2 -= 1
			}
			pointer1 -= correct
			correct = 0
		}
		if number == -1 {
			fmt.Println(pointer1, strconv.Itoa(number)[pointer1], pointer2, strconv.Itoa(number)[pointer2], correct)
		}
	}

	_, ok := streak[pointer2-pointer1]

	return (length/2 <= correct) && ok
}
