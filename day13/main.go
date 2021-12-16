package main

func main() {
	coordinates, instructions := day13ReadFile()

	day13Solution(coordinates, instructions, 1)
	day13Solution(coordinates, instructions, 2)
}
