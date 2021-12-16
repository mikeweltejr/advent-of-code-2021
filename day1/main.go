package main

import "fmt"

func main() {
	nums := readFile()

	ret := day1Solution(nums)
	fmt.Printf("Ret: %d\n", ret)

	ret2 := day1SolutionPartTwo(nums)
	fmt.Printf("Ret: %d\n", ret2)
}
