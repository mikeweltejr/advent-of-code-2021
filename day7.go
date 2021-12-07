package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func day7ReadFile() []int {
	file, _ := os.Open("day7.txt")
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

func day7Solution(nums []int) int {
	sortedArr := mergeSort(nums)
	median := sortedArr[len(sortedArr)/2]
	retCount := 0

	for _, n := range sortedArr {
		retCount += int(math.Abs(float64(n) - float64(median)))
	}

	return retCount
}

func day7SolutionPartTwo(nums []int) int {
	total := 0

	for _, n := range nums {
		total += n

	}

	minMean := int(math.Floor(float64(total) / float64(len(nums))))
	maxMean := int(math.Ceil((float64(total) / float64(len(nums)))))
	minCount := 0
	maxCount := 0

	for _, n := range nums {
		distMin := int(math.Abs(float64(n) - float64(minMean)))
		minCount += (distMin * (distMin + 1)) / 2

		distMax := int(math.Abs(float64(n) - float64(maxMean)))
		maxCount += (distMax * (distMax + 1)) / 2
	}

	if minCount < maxCount {
		return minCount
	} else {
		return maxCount
	}
}

func mergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	left := mergeSort(nums[:len(nums)/2])
	right := mergeSort(nums[len(nums)/2:])

	return merge(left, right)
}

func merge(left []int, right []int) []int {
	final := []int{}
	i := 0
	j := 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			final = append(final, left[i])
			i++
		} else {
			final = append(final, right[j])
			j++
		}
	}

	for ; i < len(left); i++ {
		final = append(final, left[i])
	}

	for ; j < len(right); j++ {
		final = append(final, right[j])
	}

	return final
}
