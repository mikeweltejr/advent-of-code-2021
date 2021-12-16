package main

import "fmt"

func main() {
	input := readFileText("day2.txt")
	day2Sol := day2Solution(input)
	fmt.Println(day2Sol)

	day2SolPartTwo := day2SolutionPartTwo(input)
	fmt.Printf("Part Two Solution: %d\n", day2SolPartTwo)
}
