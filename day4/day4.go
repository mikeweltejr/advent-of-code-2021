package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readBingoNumbersFromFile() (map[int][25]int, map[int][]int) {
	file, err := os.Open("day4.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	retMap := make(map[int][25]int)
	boardMap := make(map[int][]int)
	index := 0
	count := 0

	scanner := bufio.NewScanner(file)
	var tempIntArr [25]int
	innnerIndex := 0

	for scanner.Scan() {
		s := strings.Fields(scanner.Text())

		if len(s) == 0 {
			continue
		}

		count++

		for _, x := range s {
			num, err := strconv.Atoi(x)

			if err == nil {
				boardMap[num] = append(boardMap[num], index)
				tempIntArr[innnerIndex] = num
				innnerIndex++
			}
		}

		if count == 5 {
			retMap[index] = [25]int{}
			retMap[index] = tempIntArr

			count = 0
			index++
			innnerIndex = 0
			tempIntArr = [25]int{}
		}
	}

	return retMap, boardMap
}

func day4SolutionPartOne(nums []int, boardMap map[int][]int, boards map[int][25]int) int {
	// Cannot have a winnder until at least 5 numbers are called
	// No need to check anything until 5 numbers
	// Having a 2 dimensional array would be tedious
	// Add all numbers to a map key=number val=[]int where []int is the boards it belongs to
	// Add all boards to a map key=index val=board (board=make(map[int]bool)) where bool is marked
	// When number called check map to see which boards it belongs to
	// Loop through boards and mark
	// If any boards just marked match the following (0-4, 5-9, 10-14, 15-19, 20-24) or columns match return that set
	// Multiply the sum of those by the number chosen

	boardMarks := make(map[int][25]bool)

	for i := 0; i < len(nums); i++ {
		boardsWithMark := markNumbers(nums[i], boardMap, boards, boardMarks)

		if i >= 4 {
			isWinner, unmarkedNumbers, _ := checkWinners(nums[i], boardsWithMark, boardMarks, boards)

			if isWinner {
				sum := 0

				for _, n := range unmarkedNumbers {
					sum += n
				}
				return nums[i] * sum
			}
		}
	}

	return -1
}

func day4SolutionPartTwo(nums []int, boardMap map[int][]int, boards map[int][25]int) int {
	boardMarks := make(map[int][25]bool)
	winningBoards := make(map[int]bool)
	var unmarkedNumbers [20]int
	mark := nums[0]

	for i := 0; i < len(nums); i++ {
		boardsWithMark := markNumbers(nums[i], boardMap, boards, boardMarks)

		if i >= 4 {
			tempArr := []int{}
			for j := 0; j < len(boardsWithMark); j++ {
				if !winningBoards[boardsWithMark[j]] {
					tempArr = append(tempArr, boardsWithMark[j])
				}
			}

			isWinner, unmarked, _ := checkWinnersContinuous(nums[i], tempArr, boardMarks, boards, winningBoards)

			if isWinner {
				unmarkedNumbers = unmarked
				mark = nums[i]
			}
		}
	}

	sum := 0

	for _, n := range unmarkedNumbers {
		sum += n
	}

	return mark * sum
}

func checkWinnersContinuous(mark int, boardsWithMark []int, boardMarks map[int][25]bool, boards map[int][25]int, winningBoards map[int]bool) (bool, [20]int, int) {
	foundWinner := false
	retArr := [20]int{}
	boardIndex := 0
	for i := 0; i < len(boardsWithMark); i++ {
		boardIndex = boardsWithMark[i]
		curBoard := boards[boardIndex]

		for j := 0; j < 25; j++ {
			index := 0
			if j%5 == 0 {
				if boardMarks[boardIndex][j] && boardMarks[boardIndex][j+1] && boardMarks[boardIndex][j+2] && boardMarks[boardIndex][j+3] && boardMarks[boardIndex][j+4] {
					for k := 0; k < len(curBoard); k++ {
						if !boardMarks[boardIndex][k] {
							retArr[index] = curBoard[k]
							index++
						}
					}
					foundWinner = true
					winningBoards[boardIndex] = true
				}
			}

			if j < 5 {
				if boardMarks[boardIndex][j] && boardMarks[boardIndex][j+5] && boardMarks[boardIndex][j+10] && boardMarks[boardIndex][j+15] && boardMarks[boardIndex][j+20] {
					for k := 0; k < len(curBoard); k++ {
						if !boardMarks[boardIndex][k] {
							retArr[index] = curBoard[k]
							index++
						}
					}
					foundWinner = true
					winningBoards[boardIndex] = true
				}
			}
		}
	}

	return foundWinner, retArr, boardIndex
}

func checkWinners(mark int, boardsWithMark []int, boardMarks map[int][25]bool, boards map[int][25]int) (bool, [20]int, int) {
	for i := 0; i < len(boardsWithMark); i++ {
		boardIndex := boardsWithMark[i]
		curBoard := boards[boardIndex]

		for j := 0; j < 25; j++ {
			retArr := [20]int{}
			index := 0
			if j%5 == 0 {
				if boardMarks[boardIndex][j] && boardMarks[boardIndex][j+1] && boardMarks[boardIndex][j+2] && boardMarks[boardIndex][j+3] && boardMarks[boardIndex][j+4] {
					for k := 0; k < len(curBoard); k++ {
						if !boardMarks[boardIndex][k] {
							retArr[index] = curBoard[k]
							index++
						}
					}
					return true, retArr, boardIndex
				}
			}

			if j < 5 {
				if boardMarks[boardIndex][j] && boardMarks[boardIndex][j+5] && boardMarks[boardIndex][j+10] && boardMarks[boardIndex][j+15] && boardMarks[boardIndex][j+20] {
					for k := 0; k < len(curBoard); k++ {
						if !boardMarks[boardIndex][k] {
							retArr[index] = curBoard[k]
							index++
						}
					}
					return true, retArr, boardIndex
				}
			}
		}
	}

	return false, [20]int{}, -1
}

// Return boards that are marked
func markNumbers(mark int, boardMap map[int][]int, boards map[int][25]int, boardMarks map[int][25]bool) []int {
	boardsWithNum := boardMap[mark]

	for i := 0; i < len(boardsWithNum); i++ {
		boardIndex := boardsWithNum[i]

		curBoard := boards[boardIndex]

		for j := 0; j < len(curBoard); j++ {
			if curBoard[j] == mark {
				arr := boardMarks[boardIndex]
				arr[j] = true
				boardMarks[boardIndex] = arr
			}
		}
	}

	return boardsWithNum
}
