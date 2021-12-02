package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFileText() []string {
	file, err := os.Open("day2.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var strs []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		strs = append(strs, scanner.Text())
	}
	return strs
}

func day2Solution(input []string) int {
	h := 0
	d := 0

	for i := 0; i < len(input); i++ {
		strs := strings.Fields(input[i])

		command := strs[0]
		num, err := strconv.Atoi(strs[1])

		if err != nil {
			fmt.Println(err)
		}

		switch command {
		case "forward":
			h += num
		case "up":
			d -= num
		case "down":
			d += num
		}
	}

	return h * d
}

func day2SolutionPartTwo(input []string) int {
	h := 0
	d := 0
	a := 0

	for i := 0; i < len(input); i++ {
		strs := strings.Fields(input[i])

		command := strs[0]
		num, err := strconv.Atoi(strs[1])

		if err != nil {
			fmt.Println(err)
		}

		switch command {
		case "forward":
			h += num
			d = d + (a * num)
		case "up":
			a -= num
		case "down":
			a += num
		}
	}

	return h * d
}
