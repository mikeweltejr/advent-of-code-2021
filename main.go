package main

import "fmt"

func main() {
	// nums := readFile()

	// ret := day1Solution(nums)
	// fmt.Printf("Ret: %d\n", ret)

	// ret2 := day1SolutionPartTwo(nums)
	// fmt.Printf("Ret: %d\n", ret2)

	// input := readFileText("day2.txt")
	// day2Sol := day2Solution(input)
	// fmt.Println(day2Sol)

	// day2SolPartTwo := day2SolutionPartTwo(input)
	// fmt.Printf("Part Two Solution: %d\n", day2SolPartTwo)

	// strs := readFileText("day3.txt")
	// ret := day3Solution(strs)
	// fmt.Printf("Return: %d\n", ret)

	// ret2 := day3SolutionPartTwo(strs)
	// fmt.Printf("Ret2: %d\n", ret2)

	// nums := []int{17, 58, 52, 49, 72, 33, 55, 73, 27, 69, 88, 80, 9, 7, 59, 98, 63, 42, 84, 37, 87, 28, 97, 66, 79, 77, 61, 48, 83, 5, 94, 26, 70, 12, 51, 82, 99, 45, 22, 64, 10, 78, 13, 18, 15, 39, 8, 30, 68, 65, 40, 21, 6, 86, 90, 29, 60, 4, 38, 3, 43, 93, 44, 50, 41, 96, 20, 62, 19, 91, 23, 36, 47, 92, 76, 31, 67, 11, 0, 56, 95, 85, 35, 16, 2, 14, 75, 53, 1, 57, 81, 46, 71, 54, 24, 74, 89, 32, 25, 34}
	// boards, boardMap := readBingoNumbersFromFile()
	// ret := day4SolutionPartOne(nums, boardMap, boards)

	// ret2 := day4SolutionPartTwo(nums, boardMap, boards)

	// fmt.Printf("Day 4 Part 1: %d\n", ret)
	// fmt.Printf("Day 4 Part 2: %d\n", ret2)

	// x, y := readDay5File()

	// ret := day5Solution(x, y, false)
	// ret2 := day5Solution(x, y, true)

	// fmt.Printf("Day 5 Part One: %d\n", ret)
	// fmt.Printf("Day 5 Part Two: %d\n", ret2)

	nums := day6ReadFile()
	//retVal := day6SolutionPartOne(nums, 256)
	// retVal2 := day6SolutionPartOne(nums, 1)

	// 362639
	//fmt.Printf("Day 6 Value: %d\n", retVal)
	ret := day6Solution(nums, 256)
	fmt.Println(ret)
	// fmt.Printf("Day 6 Part Two: %d\n", retVal2)
}
