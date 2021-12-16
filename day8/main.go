package main

import "fmt"

func main() {
	signalPatterns, output := day8ReadFile()
	ret := day8SolutionPartOne(signalPatterns, output)

	fmt.Printf("Day 8 Part 1: %d\n", ret)

	ret2 := day8SolutionPartTwo(signalPatterns, output)
	fmt.Printf("Day 8 Part Two: %d\n", ret2)
}
