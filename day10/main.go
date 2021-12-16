package main

import "fmt"

func main() {
	input := day10ReadFromFile()
	ret, ret2 := day10SolutionPartOne(input)

	fmt.Printf("Day 10 solution part one: %d\n", ret)
	fmt.Printf("Day 10 solution part two: %d\n", ret2)
}
