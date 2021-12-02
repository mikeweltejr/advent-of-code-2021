package main

import "fmt"

func main() {
	// nums := readFile()

	// ret := day1Solution(nums)
	// fmt.Printf("Ret: %d\n", ret)

	// ret2 := day1SolutionPartTwo(nums)
	// fmt.Printf("Ret: %d\n", ret2)

	input := readFileText()
	day2Sol := day2Solution(input)
	fmt.Println(day2Sol)

	day2SolPartTwo := day2SolutionPartTwo(input)
	fmt.Printf("Part Two Solution: %d\n", day2SolPartTwo)
}
