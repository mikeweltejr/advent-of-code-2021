package main

import "fmt"

func main() {
	nums := day6ReadFile()
	ret := day6Solution(nums, 256)
	fmt.Println(ret)
}
