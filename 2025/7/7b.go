package main

import (
	"fmt"

	lib "monkadev.xyz/aoc"
)

func Day7B() int {
	input := lib.Read("7/input.txt")

	field := make([][]rune, len(input))
	for i, s := range input {
		field[i] = []rune(s)
	}

	return simulate(field)
}

func simulate(field [][]rune) int {
	timelines := 1

	for x := range field[0] {
		if field[0][x] != 'S' {
			field[0][x] = '.'
		}
	}

	// out := make([]string, len(field))
	// for i := range field {
	// 	out[i] = string(field[i])
	// }
	// if len(out) > 20 {
	// 	printField(out)
	// }

	last := []rune(field[0])

	for i := 1; i < len(field); i++ {
		row := []rune(field[i])

		for c := range row {
			if last[c] != 'S' && last[c] != '|' {
				continue
			}

			if len(field) == 130 {
				fmt.Println(timelines)
			}

			switch row[c] {
			case '.':
				row[c] = '|'
			case '^':
				timelines += 1
				if c > 0 {
					new := copyField(field)
					new[i][c-1] = 'S'
					timelines += simulate(new[i:])
					return timelines
				}
				if c < len(row)-1 {
					new := copyField(field)
					new[i][c+1] = 'S'
					timelines += simulate(new[i:])
					return timelines
				}
			}
		}
	}
	return timelines
}

func copyField(src [][]rune) [][]rune {
	dst := make([][]rune, len(src))
	for i := range src {
		dst[i] = make([]rune, len(src[i]))
		copy(dst[i], src[i])
	}
	return dst
}

func printField(field []string) {
	lib.Write("7/output_b.txt", field)
}
