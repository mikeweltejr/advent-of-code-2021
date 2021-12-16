package main

import "fmt"

func main() {
	input := day11ReadFile()
	ret := day11Soution(input, 100)
	fmt.Printf("Day 11 part one: %d\n", ret)
	input = day11ReadFile()
	ret2 := day11SolutionPartTwo(input)
	fmt.Printf("Day 11 part two: %d\n", ret2)
}
