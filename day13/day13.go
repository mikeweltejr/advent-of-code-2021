package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func day13ReadFile() (map[int]map[int]int, []string) {
	file, _ := os.Open("day13.txt")
	scanner := bufio.NewScanner(file)
	retMap := make(map[int]map[int]int)
	instructions := []string{}
	maxX := 0
	maxY := 0

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, ",") {
			strArr := strings.Split(line, ",")

			x, _ := strconv.Atoi(strArr[0])
			y, _ := strconv.Atoi(strArr[1])

			if x > maxX {
				maxX = x
			}

			if y > maxY {
				maxY = y
			}

			if retMap[x] == nil {
				retMap[x] = make(map[int]int)
			}
			retMap[x][y] = 1
		} else {
			if len(line) > 0 {
				instructions = append(instructions, line)
			}

		}
	}

	for i := 0; i <= maxX; i++ {
		for j := 0; j <= maxY; j++ {
			if retMap[i] == nil {
				retMap[i] = make(map[int]int)
			}
			if retMap[i][j] != 1 {
				retMap[i][j] = 0
			}
		}
	}

	return retMap, instructions
}

func day13Solution(codeMap map[int]map[int]int, instructions []string, part int) {
	foldVertical := false
	foldHorizontal := false
	yFold := 0
	xFold := 0

	if part == 1 {
		instructions = instructions[0:1]
	}

	width := len(codeMap)
	height := len(codeMap[0])
	for _, s := range instructions {
		if strings.Contains(s, "y") {
			foldVertical = true
			yFold, _ = strconv.Atoi(strings.Split(s, "=")[1])
		}
		if strings.Contains(s, "x") {
			foldHorizontal = true
			xFold, _ = strconv.Atoi(strings.Split(s, "=")[1])
		}

		if foldVertical {
			curY := yFold + 1

			for y := curY; y < height; y++ {
				for x := 0; x < width; x++ {
					if codeMap[x][y] == 1 {
						codeMap[x][height-y-1] = 1
					}
				}
			}

			height = yFold
		}

		if foldHorizontal {
			curX := xFold + 1

			for y := 0; y < height; y++ {
				for x := curX; x < width; x++ {
					if codeMap[x][y] == 1 {
						codeMap[width-x-1][y] = 1
					}
				}
			}

			width = xFold
		}
	}

	if part == 2 {
		displayMap(codeMap, height, width)
	}

	if part == 1 {
		fmt.Printf("Count after 1 fold: %d\n", countHits(codeMap, height, width))
	}

}

func displayMap(m map[int]map[int]int, height int, width int) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if m[x][y] == 0 {
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}

func countHits(m map[int]map[int]int, height int, width int) int {
	count := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if m[x][y] == 1 {
				count++
			}
		}
	}
	return count
}
