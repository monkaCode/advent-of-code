package main

import (
	"fmt"

	lib "monkadev.xyz/aoc"
)

func Day4B() {
	input := lib.Read("4/input.txt")

	var area [139][139]rune

	reachablePaperrolls := 0

	for i := 0; i < len(input) && i < len(area); i++ {
		runes := []rune(input[i])
		for j := 0; j < len(runes) && j < len(area[i]); j++ {
			area[i][j] = runes[j]
		}
	}

	removedPaperrolls := 0

	for i := 0; removedPaperrolls > 0 || i == 0; i++ {
		removedPaperrolls = 0
		area, removedPaperrolls = removeAllPossiblePaperrolls(area)
		reachablePaperrolls += removedPaperrolls
	}

	fmt.Println(reachablePaperrolls)
}

func removeAllPossiblePaperrolls(area [139][139]rune) ([139][139]rune, int) {
	reachablePaperrolls := 0
	var newArea [139][139]rune
	copy(newArea[:], area[:])
	for i := 0; i < len(area); i++ {
		for j := 0; j < len(area[i]); j++ {
			if area[i][j] == '.' {
				continue
			}
			adjacentPaperrolls := 0
			if i > 0 && j > 0 && area[i-1][j-1] == '@' {
				adjacentPaperrolls++
			}
			if i > 0 && area[i-1][j] == '@' {
				adjacentPaperrolls++
			}
			if j > 0 && area[i][j-1] == '@' {
				adjacentPaperrolls++
			}
			if i > 0 && j < len(area)-1 && area[i-1][j+1] == '@' {
				adjacentPaperrolls++
			}
			if i < len(area)-1 && j < len(area)-1 && area[i+1][j+1] == '@' {
				adjacentPaperrolls++
			}
			if i < len(area)-1 && j > 0 && area[i+1][j-1] == '@' {
				adjacentPaperrolls++
			}
			if i < len(area)-1 && area[i+1][j] == '@' {
				adjacentPaperrolls++
			}
			if j < len(area)-1 && area[i][j+1] == '@' {
				adjacentPaperrolls++
			}

			if adjacentPaperrolls < 4 {
				reachablePaperrolls++
				newArea[i][j] = '.'
			}
		}
	}

	return newArea, reachablePaperrolls
}
