package main

import (
	"bufio"
	"os"
	"sort"
)

func day10ReadFromFile() []string {
	file, _ := os.Open("day10.txt")
	scanner := bufio.NewScanner(file)

	strArr := []string{}
	for scanner.Scan() {
		strArr = append(strArr, scanner.Text())
	}

	return strArr
}

func day10SolutionPartOne(input []string) (int, int) {
	badCharArr := []rune{}
	closedCharMap := make(map[rune]rune)
	openCharMap := make(map[rune]rune)
	incompleteLines := []string{}

	closedCharMap[')'] = '('
	closedCharMap[']'] = '['
	closedCharMap['}'] = '{'
	closedCharMap['>'] = '<'

	openCharMap['('] = ')'
	openCharMap['['] = ']'
	openCharMap['{'] = '}'
	openCharMap['<'] = '>'

	for i := 0; i < len(input); i++ {
		var s Stack
		isCorrupt := false
		chars := []rune(input[i])

		for j := 0; j < len(chars); j++ {
			c := chars[j]

			if c == '(' || c == '<' || c == '[' || c == '{' {
				s.Push(string(c))
			} else {
				lastChar, found := s.Pop()
				if lastChar != string(closedCharMap[c]) || !found {
					badCharArr = append(badCharArr, c)
					isCorrupt = true
					break
				}
			}
		}

		if !isCorrupt {
			curLine := ""
			curStr, found := s.Pop()
			for found {
				curLine += string(openCharMap[rune(curStr[0])])
				curStr, found = s.Pop()
			}

			if curLine != "" {
				incompleteLines = append(incompleteLines, curLine)
			}
		}
	}

	retVal := 0
	charMap := make(map[rune]int)
	charMap[')'] = 3
	charMap[']'] = 57
	charMap['}'] = 1197
	charMap['>'] = 25137

	for _, c := range badCharArr {
		retVal += charMap[c]
	}

	charMap[')'] = 1
	charMap[']'] = 2
	charMap['}'] = 3
	charMap['>'] = 4
	incompleteScores := []int{}
	for _, str := range incompleteLines {
		chars := []rune(str)
		curScore := 0
		for _, c := range chars {
			curScore = (curScore * 5) + charMap[c]
		}

		incompleteScores = append(incompleteScores, curScore)
	}

	sort.Ints(incompleteScores)

	return retVal, incompleteScores[len(incompleteScores)/2]
}
