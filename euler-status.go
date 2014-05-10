package main

import (
	"fmt"
)

const (
	PATH  = "/home/frederick/.euler-tools/status.html"
	LINEL = 20
)

var (
	MAX    int = -1 //number of problems total
	EASY   int = 1000
	MEDIUM int = 500
)

var totals [prizes]int

func main() {
	temp, dict, difficulty := parseHTML(PATH)
	MAX = temp

	colorThreshold(dict, difficulty)

	for i := 0; i < prizes; i++ {
		totals[i], _ = prizeFns[i](dict)
	}

	chart := box(dict, LINEL)
	hist := histogramHeight(dict, difficulty, height(chart))
	fmt.Println(smash(chart, "\t|\t", hist))

	for i := 0; i < prizes; i++ {
		fmt.Printf("%-20s %2d/%-2d %s\n", names[i], totals[i], goals[i], taglines[i])
	}

	fmt.Print("\n")

	//used to keep track of problem duplication in unfinished prizes
	track := make(map[int]int)

	for pNum := 0; pNum < prizes; pNum++ {
		if totals[pNum] < goals[pNum] {
			_, set := prizeFns[pNum](dict)
			fmt.Printf("%s: %s\n\n", names[pNum], show(set, difficulty))

			for i, _ := range set {
				track[i]++
			}

		}
	}

	maxTrack := -1 //highest frequency repetition

	set := make(map[int]bool) //set of repeats
	for i := 1; i <= MAX; i++ {
		if track[i] > 1 {
			set[i] = true
		}
		if track[i] > maxTrack {
			maxTrack = track[i]
		}
	}
	fmt.Printf("Repeats: %s\n", show(set, difficulty))

	if maxTrack > 2 { //if two or less, this is just the same as 'Repeats'
		set = make(map[int]bool)
		for i := 1; i <= MAX; i++ {
			if track[i] == maxTrack {
				set[i] = true
			}
		}

		fmt.Printf("Most Repeated (%d): %s\n", maxTrack, show(set, difficulty))
	}

}
