package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type Cave struct {
	neighbors  map[string][]string
	pastVisits map[string][]int
}

func (c *Cave) addEdge(from string, to string) {
	if to == "start" {
		return
	}
	if from == "end" {
		return
	}

	if c.neighbors == nil {
		c.neighbors = make(map[string][]string)
	}

	if c.neighbors[from] == nil {
		c.neighbors[from] = []string{}
	}

	c.neighbors[from] = append(c.neighbors[from], to)
}

func day12ReadFile() *Cave {
	file, _ := os.Open("day12.txt")
	scanner := bufio.NewScanner(file)
	cave := &Cave{}

	for scanner.Scan() {
		line := scanner.Text()
		strArr := strings.Split(line, "-")

		from := strArr[0]
		to := strArr[1]

		cave.addEdge(from, to)
		cave.addEdge(to, from)
	}

	return cave
}

func countPaths(c *Cave, part int) {
	pathCount := new(int)
	searchCave(c, "start", 1, pathCount, part)

	fmt.Println(*pathCount)
}

func searchCave(cave *Cave, curCave string, steps int, pathCount *int, part int) {
	if isSmallCave(curCave) {
		if cave.pastVisits == nil {
			cave.pastVisits = make(map[string][]int)
		}

		if cave.pastVisits[curCave] == nil {
			cave.pastVisits[curCave] = []int{}
		}
		cave.pastVisits[curCave] = append(cave.pastVisits[curCave], steps)
	}

	for _, next := range cave.neighbors[curCave] {
		if next == "end" {
			*pathCount++
			continue
		}

		if isSmallCave(next) {
			if len(cave.pastVisits[next]) > 0 {
				if part == 1 {
					continue
				}

				hasVisitedTwice := false
				for _, c := range cave.pastVisits {
					if len(c) > 1 {
						hasVisitedTwice = true
						break
					}
				}

				if hasVisitedTwice {
					continue
				}
			}
		}

		searchCave(cave, next, steps+1, pathCount, part)

		for key := range cave.pastVisits {
			cave.pastVisits[key] = getPriorSteps(cave.pastVisits[key], steps)
		}
	}
}

func getPriorSteps(n []int, s int) []int {
	ret := []int{}
	for _, v := range n {
		if v <= s {
			ret = append(ret, v)
		}
	}

	return ret
}

func isSmallCave(s string) bool {
	return unicode.IsLower(rune(s[0]))
}
