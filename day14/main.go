package main

import "fmt"

func main() {
	templateStr, polymerMap := day14ReadFile()
	ret := day14Solution(templateStr, polymerMap, 40)
	fmt.Printf("Day 14 Solution: %d\n", ret)
}
