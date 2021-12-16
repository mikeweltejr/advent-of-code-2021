package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func readDay5File() ([]int, []int) {
	file, _ := os.Open("day5.txt")
	scanner := bufio.NewScanner(file)

	x := []int{}
	y := []int{}

	for scanner.Scan() {
		curLine := strings.Split(scanner.Text(), "->")
		firstLine := strings.Split(curLine[0], ",")
		secondLine := strings.Split(curLine[1], ",")

		x1, _ := strconv.Atoi(strings.TrimSpace(firstLine[0]))
		x2, _ := strconv.Atoi(strings.TrimSpace(secondLine[0]))
		y1, _ := strconv.Atoi(strings.TrimSpace(firstLine[1]))
		y2, _ := strconv.Atoi(strings.TrimSpace(secondLine[1]))

		x = append(x, x1)
		y = append(y, y1)
		x = append(x, x2)
		y = append(y, y2)
	}

	return x, y
}

func day5Solution(x []int, y []int, diagonal bool) int {
	var grid [999][999]int
	countPoints := 0

	for i := 0; i < len(x)-1; i = i + 2 {
		if x[i] == x[i+1] {
			yIndex := 0
			yEnd := 0
			if y[i] < y[i+1] {
				yIndex = y[i]
				yEnd = y[i+1]
			} else {
				yIndex = y[i+1]
				yEnd = y[i]
			}

			for yIndex <= yEnd {
				curVal := grid[x[i]][yIndex]
				curVal += 1
				grid[x[i]][yIndex] = curVal
				yIndex++

				if curVal == 2 {
					countPoints++
				}
			}
		} else if y[i] == y[i+1] {
			xIndex := 0
			xEnd := 0
			if x[i] < x[i+1] {
				xIndex = x[i]
				xEnd = x[i+1]
			} else {
				xIndex = x[i+1]
				xEnd = x[i]
			}

			for xIndex <= xEnd {
				curVal := grid[xIndex][y[i]]
				curVal += 1
				grid[xIndex][y[i]] = curVal
				xIndex++

				if curVal == 2 {
					countPoints++
				}
			}
		} else if diagonal {
			xMinus := x[i] > x[i+1]
			yMinus := y[i] > y[i+1]
			xIndex := x[i]
			yIndex := y[i]

			for xIndex != x[i+1] && yIndex != y[i+1] {
				curVal := grid[xIndex][yIndex]
				curVal += 1
				grid[xIndex][yIndex] = curVal

				if curVal == 2 {
					countPoints++
				}

				if xMinus {
					xIndex--
				} else {
					xIndex++
				}

				if yMinus {
					yIndex--
				} else {
					yIndex++
				}
			}

			curVal := grid[x[i+1]][y[i+1]]
			curVal += 1
			grid[x[i+1]][y[i+1]] = curVal

			if curVal == 2 {
				countPoints++
			}
		}
	}

	return countPoints
}
