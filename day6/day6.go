package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func day6ReadFile() []int {
	file, _ := os.Open("day6.txt")
	scanner := bufio.NewScanner(file)

	retArr := []int{}

	for scanner.Scan() {
		strArr := strings.Split(scanner.Text(), ",")

		for _, s := range strArr {
			i, _ := strconv.Atoi(s)
			retArr = append(retArr, i)
		}
	}

	return retArr
}

func day6Solution(nums []int, days int) int {
	fishDays := [9]int{}

	for i := 0; i < len(nums); i++ {
		fishDays[nums[i]] = fishDays[nums[i]] + 1
	}

	for day := 0; day < days; day++ {
		zeroDay := fishDays[0]
		for i := 1; i < len(fishDays); i++ {
			fishDays[i-1] = fishDays[i]
		}
		fishDays[6] += zeroDay
		fishDays[8] = zeroDay
	}

	sum := 0
	for _, day := range fishDays {
		sum += day
	}

	return sum
}
