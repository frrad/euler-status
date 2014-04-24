package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	path   = "/home/frederick/Projects/project-euler/eulerdata/status.html"
	prizes = 5
)

var totals [prizes]int

var names = [prizes]string{
	"Prime Obsession",
	"Triangle Trophy",
	"Lucky Luke",
	"Decimation II",
	"Ultimate Decimator",
}

var taglines = [prizes]string{
	"prime numbered problems",
	"first triangle numbered problems",
	"lucky numbered problems",
	"rows",
	"rows",
}

var goals = [prizes]int{
	50,
	25,
	50,
	10,
	10,
}

var prizeFns = make([]func(map[int]bool) (int, map[int]bool), prizes)

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

const EASY int = 1000
const MEDIUM int = 500

func colorize(text string, score int) string {
	if score > EASY {
		return "\033[01;32m" + text + "\033[00m"
	}

	if score > MEDIUM {
		return "\033[1;33m" + text + "\033[00m"
	}

	return "\033[1;31m" + text + "\033[00m"

}

func getNum(a string) int {
	probLen := 8 //Length of `Problem '

	starts := strings.Index(a, "Problem ")
	ends := strings.Index(a[starts+probLen:], " ")

	isolated := a[starts+probLen : starts+probLen+ends]
	number, _ := strconv.Atoi(isolated)

	return number
}

func luckySeive(max int) []int {

	//fmt.Printf("Debug: %d\n", max)

	luckyseive := make([]int, max)
	for i := 0; i < max; i++ {
		luckyseive[i] = i + 1
	}

	last := -1
	pointer := 1

	for pointer < len(luckyseive) {

		last = luckyseive[pointer]

		for del := last - 1; del < len(luckyseive); del += last {
			luckyseive[del] = 0
		}
		for i := 0; i < len(luckyseive); i++ {
			if luckyseive[i] == 0 {
				luckyseive = append(luckyseive[:i], luckyseive[i+1:]...)
				i--
			}

		}

		if luckyseive[pointer] == last {
			pointer++
		}

	}

	return luckyseive
}

func howHard(text string) int {
	start := strings.Index(text, "solved by")
	start += 10 //length of "solved by"
	text = text[start:]

	end := strings.Index(text, "members")
	text = text[:end-1]

	//fmt.Printf("%s\n\n\n", text)

	ans, err := strconv.Atoi(text)
	if err == nil {
		//fmt.Printf("$d\n", ans)
		return ans
	}

	fmt.Printf("ERROR: %s\n", err)
	return 0
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

var max int = -1 //number of problems total

func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

//stolen from euler.Import
func inWrapper(filename string) []string {

	// read whole the file
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var output []string

	currentline := ""

	for _, char := range b {
		if char == 10 {
			output = append(output, currentline)
			currentline = ""
		} else {
			currentline += string(char)
		}
	}

	if currentline != "" {
		output = append(output, currentline)
	}

	return output

}

func main() {

	//PRIME NUMBERS (Index = 0)
	prizeFns[0] = func(dict map[int]bool) (ans int, set map[int]bool) {
		set = make(map[int]bool)
		for i := 1; i <= max; i++ {
			if dict[i] {
				if isPrime(i) {
					ans++
				}
			} else if isPrime(i) {
				set[i] = true
			}
		}
		return
	}

	//TRIANGLE NUMBERS (Index = 1)
	prizeFns[1] = func(dict map[int]bool) (ans int, set map[int]bool) {
		set = make(map[int]bool)
		for i := 1; i <= 25; i++ {
			if dict[i*(i+1)/2] {
				ans++
			} else {
				set[i*(i+1)/2] = true
			}
		}
		return
	}

	//LUCKY NUMBER (Index = 2)
	prizeFns[2] = func(dict map[int]bool) (ans int, set map[int]bool) {
		set = make(map[int]bool)
		luckyseive := luckySeive(max)
		for i := 0; i < len(luckyseive); i++ {
			if dict[luckyseive[i]] {
				ans++
			} else {
				set[luckyseive[i]] = true
			}
		}
		return
	}

	//DECIMATION II (Index = 3)
	prizeFns[3] = func(dict map[int]bool) (ans int, set map[int]bool) {
		set = make(map[int]bool)
		decStart := 200
		for i := 0; i < 10; i++ {
			here := 0
			for j := decStart + 10*i + 1; j < decStart+10*(i+1)+1; j++ {
				if dict[j] {
					here++
				}
			}

			if here > 0 {
				ans++
			} else {

				for j := decStart + 10*i + 1; j < decStart+10*(i+1)+1; j++ {
					set[j] = true
				}
			}
		}
		return
	}

	//ULTIMATE DECIMATOR (Index = 4)
	prizeFns[4] = func(dict map[int]bool) (ans int, set map[int]bool) {
		set = make(map[int]bool)
		decStart := 300
		for i := 0; i < 10; i++ {
			here := 0
			for j := decStart + 10*i + 1; j < decStart+10*(i+1)+1; j++ {
				if dict[j] {
					here++
				}
			}

			if here > 0 {
				ans++
			} else {

				for j := decStart + 10*i + 1; j < decStart+10*(i+1)+1; j++ {
					set[j] = true
				}
			}
		}
		return
	}

	lineL := 40

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
