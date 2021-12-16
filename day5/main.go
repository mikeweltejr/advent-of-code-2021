package main

import "fmt"

func main() {
	x, y := readDay5File()

	ret := day5Solution(x, y, false)
	ret2 := day5Solution(x, y, true)

	fmt.Printf("Day 5 Part One: %d\n", ret)
	fmt.Printf("Day 5 Part Two: %d\n", ret2)
}
