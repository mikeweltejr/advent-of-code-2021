package main

import (
	"bufio"
	"os"
	"strconv"
)

// Add all values to a map[int]map[int]int

// Each step add 1 to itself

// If == 9 set to 0 and add 1 to every top, bottom, left, right, diaganol value

func day11ReadFile() map[int]map[int]int {
	file, _ := os.Open("day11.txt")
	scanner := bufio.NewScanner(file)
	retMap := make(map[int]map[int]int)
	y := 0

	for scanner.Scan() {
		line := scanner.Text()

		x := 0

		for _, s := range line {
			i, _ := strconv.Atoi(string(s))

			if retMap[x] == nil {
				retMap[x] = make(map[int]int)
			}

			retMap[x][y] = i
			x++
		}

		y++
	}

	return retMap
}

func day11Soution(input map[int]map[int]int, n int) int {
	flashCount := 0

	for i := 1; i <= n; i++ {
		flashMap := make(map[int]map[int]bool)
		for y := 0; y < len(input[0]); y++ {
			for x := 0; x < len(input); x++ {
				flashCount += flashRelatives(input, x, y, flashMap)
			}
		}
	}

	return flashCount
}

func day11SolutionPartTwo(input map[int]map[int]int) int {
	octopusTotal := 0
	index := 1
	allFlashIndex := -1
	for y := 0; y < len(input[0]); y++ {
		for x := 0; x < len(input); x++ {
			octopusTotal += 1
		}
	}

	for allFlashIndex == -1 {
		curFlashCount := 0
		flashMap := make(map[int]map[int]bool)
		for y := 0; y < len(input[0]); y++ {
			for x := 0; x < len(input); x++ {
				curFlashCount += flashRelatives(input, x, y, flashMap)
			}
		}

		if curFlashCount == octopusTotal {
			return index
		}
		index++
	}

	return -1
}

func flashRelatives(input map[int]map[int]int, x int, y int, flashMap map[int]map[int]bool) int {
	count := 0
	if x < 0 || x > len(input)-1 || y < 0 || y > len(input)-1 {
		return 0
	}

	if flashMap[x] == nil {
		flashMap[x] = make(map[int]bool)
	}

	if input[x][y] == 0 && flashMap[x][y] {
		return 0
	}

	if input[x][y] == 9 {
		count++
		input[x][y] = 0
		flashMap[x][y] = true
		count += flashRelatives(input, x, y+1, flashMap)
		count += flashRelatives(input, x, y-1, flashMap)
		count += flashRelatives(input, x-1, y, flashMap)
		count += flashRelatives(input, x+1, y, flashMap)
		count += flashRelatives(input, x-1, y+1, flashMap)
		count += flashRelatives(input, x+1, y+1, flashMap)
		count += flashRelatives(input, x+1, y-1, flashMap)
		count += flashRelatives(input, x-1, y-1, flashMap)
	} else {
		input[x][y] = input[x][y] + 1
	}

	return count
}
