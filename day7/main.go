package main

import "fmt"

func main() {
	nums := day7ReadFile()
	ret := day7Solution(nums)
	ret2 := day7SolutionPartTwo(nums)

	fmt.Printf("Day 7 Return: %d\n", ret)
	fmt.Printf("Day 7 Part 2 Return: %d\n", ret2)
}
