package main

import (
	"sort"
	"strconv"
	"strings"

	lib "monkadev.xyz/aoc"
)

func Day5B() int {
	input := lib.Read("5/input.txt")
	inputSeperator := ""

	allRanges := make(map[int]int)
	freshCount := 0

	for r := range input {
		row := input[r]

		if row == inputSeperator {
			break
		}

		numbers := strings.Split(row, "-")
		first, err1 := strconv.Atoi(numbers[0])
		last, err2 := strconv.Atoi(numbers[1])
		if err1 != nil || err2 != nil {
			return 0
		}

		lastInCurrentRange, ok := allRanges[first]
		if (ok && last > lastInCurrentRange) || !ok {
			allRanges[first] = last
		}
	}

	simplifiedRanges := SimplifyRanges(allRanges)

	keys := make([]int, 0, len(simplifiedRanges))
	for k := range simplifiedRanges {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for i := 0; i < len(simplifiedRanges); i++ {
		freshCount += simplifiedRanges[keys[i]] - keys[i] + 1
	}

	return freshCount
}
