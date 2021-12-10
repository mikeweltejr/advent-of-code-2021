package main

import (
	"bufio"
	"math"
	"os"
	"sort"
	"strconv"
)

func day9ReadFile() map[int][]int {
	file, _ := os.Open("day9.txt")
	scanner := bufio.NewScanner(file)

	intArr := make(map[int][]int)

	mapIndex := 0
	for scanner.Scan() {
		line := scanner.Text()

		for i := 0; i < len(line); i++ {
			curInt, _ := strconv.Atoi(string(line[i]))
			intArr[mapIndex] = append(intArr[mapIndex], curInt)
		}
		mapIndex++
	}

	return intArr
}

func day9Solution(nums map[int][]int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		curArr := nums[i]
		belowarr := []int{}
		topArr := []int{}

		if i < len(nums)-1 {
			belowarr = nums[i+1]
		}

		if i > 0 {
			topArr = nums[i-1]
		}

		for j := 0; j < len(curArr); j++ {
			n := curArr[j]

			if n == 9 {
				continue
			}

			if n == 0 {
				count = count + (n + 1)
				continue
			}

			// Check bottom
			if i < len(nums)-1 && n > belowarr[j] {
				continue
			}

			// Check top
			if i > 0 && n > topArr[j] {
				continue
			}

			// Check left
			if j > 0 && n > curArr[j-1] {
				continue
			}

			// Check right
			if j < len(curArr)-1 && n > curArr[j+1] {
				continue
			}

			count = count + (n + 1)
		}
	}

	return count
}

type Point struct {
	x   int
	y   int
	val int
}

func day9SolutionPartTwo(nums map[int][]int) int {
	width := len(nums[0])
	height := len(nums)

	data := make(map[int]map[int]int)
	retMap := make(map[int]map[int]int)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			v := nums[y][x]

			if data[x] == nil {
				data[x] = make(map[int]int)
			}

			data[x][y] = v
		}
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			currentX := x
			currentY := y
			currentHeight := data[currentX][currentY]
			if currentHeight == 9 {
				continue
			}

			downX, downY := lowestNeighbor(x, y, width, height, data)

			for data[downX][downY] < currentHeight {
				currentX = downX
				currentY = downY
				currentHeight = data[currentX][currentY]
				downX, downY = lowestNeighbor(currentX, currentY, width, height, data)
			}

			if retMap[currentX] == nil {
				retMap[currentX] = make(map[int]int)
			}
			retMap[currentX][currentY] = retMap[currentX][currentY] + 1
		}
	}

	basinVals := []int{}

	for _, v := range retMap {
		for _, vv := range v {
			basinVals = append(basinVals, vv)
		}
	}

	sort.Slice(basinVals, func(i, j int) bool {
		return basinVals[i] > basinVals[j]
	})

	return basinVals[0] * basinVals[1] * basinVals[2]
}

func lowestNeighbor(x int, y int, width int, height int, data map[int]map[int]int) (int, int) {
	min := math.MaxInt32
	minX := x
	minY := y
	val := 0

	if x > 0 {
		val = data[x-1][y]

		if val < min {
			min = val
			minX = x - 1
			minY = y
		}
	}

	if y > 0 {
		val = data[x][y-1]

		if val < min {
			min = val
			minX = x
			minY = y - 1
		}
	}

	if x < width-1 {
		val = data[x+1][y]
		if val < min {
			min = val
			minX = x + 1
			minY = y
		}
	}

	if y < height-1 {
		val = data[x][y+1]
		if val < min {
			min = val
			minX = x
			minY = y + 1
		}
	}

	return minX, minY
}
