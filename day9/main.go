package main

import "fmt"

func main() {
	nums := day9ReadFile()
	ret := day9Solution(nums)
	ret2 := day9SolutionPartTwo(nums)

	fmt.Printf("Day 9 part one: %d\n", ret)
	fmt.Printf("Day 9 part two: %d\n", ret2)
}
