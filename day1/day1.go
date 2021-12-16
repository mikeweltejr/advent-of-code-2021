package main

import (
	"fmt"
	"io"
	"os"
)

func readFile() []int {
	file, err := os.Open("day1.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var perline int
	var nums []int

	for {

		_, err := fmt.Fscanf(file, "%d\n", &perline) // give a patter to scan

		if err != nil {

			if err == io.EOF {
				break // stop reading the file
			}
			fmt.Println(err)
			os.Exit(1)
		}

		nums = append(nums, perline)
	}

	return nums
}

func day1Solution(nums []int) int {
	count := 0

	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] > nums[i] {
			count++
		}
	}

	return count
}

func day1SolutionPartTwo(nums []int) int {
	count := 0

	for i := 0; i < len(nums)-3; i++ {
		if nums[i+1]+nums[i+2]+nums[i+3] > nums[i]+nums[i+1]+nums[i+2] {
			count++
		}
	}

	return count
}
