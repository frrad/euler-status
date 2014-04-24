package main

import (
	"fmt"
	"strings"
)

const (
	path       = "/home/frederick/Projects/project-euler/eulerdata/status.html"
	lineL      = 40
	EASY   int = 1000
	MEDIUM int = 500
)

var max int = -1 //number of problems total

func main() {
	page := inWrapper(path)

	dict := make(map[int]bool)
	difficulty := make(map[int]int)

	for _, line := range page {
		split := strings.Split(line, "class=\"problem")
		for _, prob := range split {

			if len(prob) >= 9 {
				if prob[:7] == "_solved" {
					//fmt.Printf("Debug (solved): %s\n", prob)
					number := getNum(prob)
					difficulty[number] = howHard(prob)
					dict[number] = true
					if number > max {
						max = number
					}

				} else if prob[:9] == "_unsolved" {
					number := getNum(prob)
					difficulty[number] = howHard(prob)
					if number > max {
						max = number
					}
				}
			}

		}
	}

	for i := 0; i < prizes; i++ {
		totals[i], _ = prizeFns[i](dict)
	}

	fmt.Println(smash(box(dict, lineL), " | ", histogram(dict, difficulty, 140)))

	for i := 0; i < prizes; i++ {
		fmt.Printf("%-20s %2d/%-2d %s\n", names[i], totals[i], goals[i], taglines[i])
	}

	fmt.Print("\n")

	track := make(map[int]int)

	for pNum := 1; pNum <= 4; pNum++ {
		if totals[pNum] < goals[pNum] {
			_, set := prizeFns[pNum](dict)
			fmt.Printf("%s: %s\n\n", names[pNum], show(set, difficulty))

			for i, _ := range set {
				track[i]++
			}

		}
	}

	maxTrack := -1

	set := make(map[int]bool)

	for i := 1; i <= max; i++ {
		if track[i] > 1 {
			set[i] = true
		}
		if track[i] > maxTrack {
			maxTrack = track[i]
		}
	}
	fmt.Printf("Repeats: %s\n", show(set, difficulty))

	set = make(map[int]bool)
	for i := 1; i <= max; i++ {
		if track[i] == maxTrack {
			set[i] = true
		}
	}
	fmt.Printf("Most Repeated (%d): %s", maxTrack, show(set, difficulty))

	fmt.Print("\n")

}
