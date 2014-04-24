package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

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

func getNum(a string) int {
	probLen := 8 //Length of `Problem '

	starts := strings.Index(a, "Problem ")
	ends := strings.Index(a[starts+probLen:], " ")

	isolated := a[starts+probLen : starts+probLen+ends]
	number, _ := strconv.Atoi(isolated)

	return number
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

func parseHTML(path string) (max int, dict map[int]bool, difficulty map[int]int) {
	dict = make(map[int]bool)
	difficulty = make(map[int]int)

	page := inWrapper(path)

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

	return
}
