package main

import "fmt"

func main() {
	strs := readFileText("day3.txt")
	ret := day3Solution(strs)
	fmt.Printf("Return: %d\n", ret)

	ret2 := day3SolutionPartTwo(strs)
	fmt.Printf("Ret2: %d\n", ret2)
}
