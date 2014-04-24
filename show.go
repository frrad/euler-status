package main

import (
	"fmt"
	"strings"
)

//Returns a box with Xs in positions corresponding to completed
//questions.
func box(dict map[int]bool, lineL int) (picture string) {
	done := 0
	for i := 1; i <= max; i++ {
		if dict[i] {
			done++
		}
	}

	picture += fmt.Sprintf("Done %d/%d problems\n", done, max)
	picture += strings.Repeat("=", lineL)
	picture += "\n"

	for i := 1; i <= max; i++ {
		if dict[i] {
			picture += "X"
		} else if i == max {
			picture += "O"
		} else {
			picture += " "
		}

		if i%lineL == 0 {
			picture += "\n"
		}
	}
	if max%lineL != 0 {
		picture += "\n"
	}
	picture += strings.Repeat("=", lineL)
	picture += "\n"
	return
}

//histogram returns a histogram of the data in supplied list
func histogram(dict map[int]bool, difficulty map[int]int, width int) (ans string) {
	list := []int{}
	for i := 1; i <= max; i++ {
		if !dict[i] {
			list = append(list, difficulty[i])
		}
	}

	count := func(a, b int) (total int) {
		for _, obj := range list {
			if obj >= a && obj <= b {
				total++
			}
		}
		return
	}

	start := 0

	drawn := 0
	for i := start; drawn < len(list); i += width {
		barLength := count(i, i+width-1)
		drawn += barLength
		bar := colorize(strings.Repeat("+", barLength), i+width/2)
		ans += fmt.Sprintf("%-4d-%4d: %s\n", i, i+width-1, bar)
	}

	return ans
}
