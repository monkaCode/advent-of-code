package main

import (
	"sort"
	"strconv"
	"strings"

	lib "monkadev.xyz/aoc"
)

func Day5A() int {
	input := lib.Read("5/input.txt")
	inputSeperator := ""

	allRanges := make(map[int]int)
	currentlyReadingIDs := false
	var ids []int
	freshCount := 0

	for r := range input {
		row := input[r]

		if row == inputSeperator {
			currentlyReadingIDs = true
			continue
		}

		if currentlyReadingIDs {
			number, err := strconv.Atoi(row)
			if err != nil {
				return 0
			}
			ids = append(ids, number)
		} else {
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
	}

	simplifiedRanges := SimplifyRanges(allRanges)

	keys := make([]int, 0, len(simplifiedRanges))
	for k := range simplifiedRanges {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	sort.Ints(ids)

	rangesIndex := 0
	for i := 0; i < len(ids); i++ {
		id := ids[i]

		if id >= keys[rangesIndex] && id <= simplifiedRanges[keys[rangesIndex]] {
			// Inside current range
			freshCount++
		} else if id >= keys[rangesIndex] && rangesIndex < len(simplifiedRanges)-1 {
			// Outside current range
			rangesIndex++
			i--
		}
	}
	return freshCount
}

func SimplifyRanges(ranges map[int]int) map[int]int {
	simplifiedRanges := make(map[int]int)

	keys := make([]int, 0, len(ranges))
	for k := range ranges {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	start := keys[0]
	end := ranges[keys[0]]
	for i := 1; i < len(keys); i++ {
		if keys[i] <= end && ranges[keys[i]] <= end {
			// Drop range
			delete(ranges, keys[i])
		} else if keys[i] <= end+1 {
			// Extend current range
			end = ranges[keys[i]]
		} else {
			// Ending current range and start next
			simplifiedRanges[start] = end
			start = keys[i]
			end = ranges[keys[i]]
		}
	}
	simplifiedRanges[start] = end

	return simplifiedRanges
}
