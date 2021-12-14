package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func day14ReadFile() (string, map[string]string) {
	file, _ := os.Open("day14.txt")
	scanner := bufio.NewScanner(file)
	templateStr := ""
	polymerMap := make(map[string]string)

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) <= 0 {
			continue
		}

		if strings.Contains(line, "->") {
			curStr := strings.Split(line, " -> ")
			key := curStr[0]
			val := curStr[1]

			polymerMap[key] = val
		} else {
			templateStr = line
		}
	}

	return templateStr, polymerMap
}

func day14Solution(templateStr string, polymerMap map[string]string, x int) int {
	charMap := make(map[string]int)

	for i := 0; i < len(templateStr)-1; i++ {
		charMap[templateStr[i:i+2]] += 1
	}
	charMap[templateStr[len(templateStr)-1:]] += 1

	for i := 0; i < x; i++ {
		newCharMap := make(map[string]int)

		for k, v := range charMap {
			if polymerMap[k] != "" {
				newCharMap[k[0:1]+polymerMap[k]] += v
				newCharMap[polymerMap[k]+k[1:2]] += v
			} else {
				newCharMap[k] += v
			}
		}
		charMap = newCharMap
	}

	freq := make(map[string]int)
	for k, v := range charMap {
		freq[k[0:1]] += v
	}

	minVal := math.MaxInt64
	maxVal := 0
	minChar := ""
	maxChar := ""
	for key, val := range freq {
		if val < minVal {
			minVal = val
			minChar = key
		}

		if val > maxVal {
			maxVal = val
			maxChar = key
		}
	}

	fmt.Printf("MaxChar: %v MaxVal: %d ----- MinChar: %v MinVal: %d\n", maxChar, maxVal, minChar, minVal)

	return maxVal - minVal
}
