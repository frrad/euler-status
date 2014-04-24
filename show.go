package main

import (
	"fmt"
	"strconv"
	"strings"
)

func show(set map[int]bool, howHard map[int]int) string {
	ans := ""
	for i := 0; i < max; i++ {
		if set[i] {
			ans += strconv.Itoa(i)
			ans += "("
			ans += colorize(strconv.Itoa(howHard[i]), howHard[i])
			ans += ") "
		}
	}
	return ans
}

func colorize(text string, score int) string {
	if score > EASY {
		return "\033[01;32m" + text + "\033[00m"
	}

	if score > MEDIUM {
		return "\033[1;33m" + text + "\033[00m"
	}

	return "\033[1;31m" + text + "\033[00m"
}

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

//Histogram returns a histogram of the data in supplied list
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

//Given two objects, display a to left or b
func smash(a, sep, b string) (smoosh string) {

	aPieces, bPieces := strings.Split(a, "\n"), strings.Split(b, "\n")
	for aPieces[len(aPieces)-1] == "" {
		aPieces = aPieces[:len(aPieces)-1]
	}
	for bPieces[len(bPieces)-1] == "" {
		bPieces = bPieces[:len(bPieces)-1]
	}

	paddle := 0
	for _, ln := range aPieces {
		if len(ln) > paddle {
			paddle = len(ln)
		}
	}

	for len(aPieces) > len(bPieces) {
		bPieces = append(bPieces, "")
	}
	for len(bPieces) > len(aPieces) {
		aPieces = append(aPieces, "")
	}

	for i := 0; i < len(aPieces); i++ {
		lump := aPieces[i]
		lump += strings.Repeat(" ", paddle-len(aPieces[i]))
		lump += sep
		lump += bPieces[i]
		lump += "\n"
		smoosh += lump
	}

	return
}
