package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func day8ReadFile() (map[int][]string, map[int][]string) {
	file, _ := os.Open("day8.txt")
	scanner := bufio.NewScanner(file)

	index := 0
	signalPatternMap := make(map[int][]string)
	outputMap := make(map[int][]string)

	for scanner.Scan() {
		strArr := strings.Split(scanner.Text(), "|")
		signalPatterns := strings.Split(strArr[0], " ")
		output := strings.Split(strArr[1], " ")

		signalPatternMap[index] = signalPatterns
		outputMap[index] = output
		index++
	}

	return signalPatternMap, outputMap
}

// Answer: 392
func day8SolutionPartOne(signalPatterns map[int][]string, output map[int][]string) int {
	count := 0

	for i := 0; i < len(output); i++ {
		curOutput := output[i]

		for j := 0; j < len(curOutput); j++ {
			o := len(curOutput[j])
			if o == 2 || o == 4 || o == 3 || o == 7 {
				count++
			}
		}
	}

	return count
}

func day8SolutionPartTwo(signalPatterns map[int][]string, output map[int][]string) int {

	ret := 0

	for i := 0; i < len(signalPatterns); i++ {
		signalPattern := signalPatterns[i]
		digitMap := make(map[int]string)

		sort.Slice(signalPattern, func(i, j int) bool {
			return len(signalPattern[i]) < len(signalPattern[j])
		})
		for _, s := range signalPattern {
			switch len(s) {
			case 2:
				digitMap[1] = s
			case 3:
				digitMap[7] = s
			case 4:
				digitMap[4] = s
			case 7:
				digitMap[8] = s
			case 5:
				x := getFiveThreeOrTwo(digitMap[1], digitMap[4], s)
				digitMap[x] = s
			case 6:
				x := getZeroSixOrNine(digitMap[1], digitMap[3], s)
				digitMap[x] = s
			}
		}

		tempStr := ""
		o := output[i]

		for _, s := range o {
			switch len(s) {
			case 2:
				tempStr += "1"
			case 3:
				tempStr += "7"
			case 4:
				tempStr += "4"
			case 7:
				tempStr += "8"
			case 5:
				tempStr += fmt.Sprint(getFiveThreeOrTwo(digitMap[1], digitMap[4], s))
			case 6:
				tempStr += fmt.Sprint(getZeroSixOrNine(digitMap[1], digitMap[3], s))
			}
		}

		num, _ := strconv.Atoi(tempStr)
		ret += num
	}

	return ret
}

func getFiveThreeOrTwo(oneString string, fourString string, curString string) int {
	oneCount := 0
	fourCount := 0

	for _, s := range oneString {
		if strings.Contains(curString, string(s)) {
			oneCount++
		}
	}

	for _, s := range fourString {
		if strings.Contains(curString, string(s)) {
			fourCount++
		}
	}

	if oneCount == 2 {
		return 3
	} else if fourCount == 3 {
		return 5
	} else {
		return 2
	}
}

func getZeroSixOrNine(oneString string, threeString string, curString string) int {
	oneCount := 0
	threeCount := 0

	for _, s := range oneString {
		if strings.Contains(curString, string(s)) {
			oneCount++
		}
	}

	for _, s := range threeString {
		if strings.Contains(curString, string(s)) {
			threeCount++
		}
	}

	if oneCount == 1 {
		return 6
	} else if threeCount == 4 {
		return 0
	} else {
		return 9
	}
}
