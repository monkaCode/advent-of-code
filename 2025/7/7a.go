package main

import (
	lib "monkadev.xyz/aoc"
)

func Day7A() int {
	input := lib.Read("7/input.txt")

	output := make([]string, len(input))
	last := []rune(input[0])
	output[0] = string(last)
	splitted := 0

	for i := 1; i < len(input); i++ {
		row := []rune(input[i])

		for c := range row {
			if last[c] != 'S' && last[c] != '|' {
				continue
			}

			switch row[c] {
			case '.':
				row[c] = '|'
			case '^':
				if c > 0 {
					row[c-1] = '|'
				}
				if c < len(row)-1 {
					row[c+1] = '|'
				}
				splitted += 1
			}
		}
		last = row
		output[i] = string(row)
	}
	output[len(input)-1] = string(last)
	lib.Write("7/output_a.txt", output)

	return splitted
}
