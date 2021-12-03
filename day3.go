package main

import (
	"fmt"
	"strconv"
)

func day3Solution(binaryArray []string) int64 {
	gammaBit := filterBinary(binaryArray, 0, 1)
	epsilonBit := filterBinary(binaryArray, 0, 0)

	gInt, err := strconv.ParseInt(gammaBit, 2, 16)
	eInt, errE := strconv.ParseInt(epsilonBit, 2, 16)

	if err != nil || errE != nil {
		fmt.Println("Error parsing binary")
	}

	return gInt * eInt
}

func filterBinary(strs []string, ind int, strong int) string {
	retStr := ""
	if ind >= len([]rune(strs[0])) {
		return retStr
	}

	zCount := 0
	oCount := 0
	for i := 0; i < len(strs); i++ {
		chars := []rune(strs[i])

		if string(chars[ind]) == "0" {
			zCount++
		} else {
			oCount++
		}
	}

	if oCount > zCount {
		if strong == 1 {
			retStr += "1" + filterBinary(strs, ind+1, strong)
		} else {
			retStr += "0" + filterBinary(strs, ind+1, strong)
		}
	} else {
		if strong == 1 {
			retStr += "0" + filterBinary(strs, ind+1, strong)
		} else {
			retStr += "1" + filterBinary(strs, ind+1, strong)
		}
	}

	return retStr
}

func day3SolutionPartTwo(binaryArray []string) int64 {
	oxygenBinary := filterLists(binaryArray, 0, 1)
	coBinary := filterLists(binaryArray, 0, 0)

	oInt, err := strconv.ParseInt(oxygenBinary, 2, 16)
	cInt, errE := strconv.ParseInt(coBinary, 2, 16)

	if err != nil || errE != nil {
		fmt.Println("Error converting binary")
	}

	return oInt * cInt
}

func filterLists(strs []string, ind int, strong int) string {
	if len(strs) == 1 || ind >= len([]rune(strs[0])) {
		return strs[0]
	}

	zCount := 0
	oCount := 0
	binMap := make(map[int][]string)
	for i := 0; i < len(strs); i++ {
		chars := []rune(strs[i])

		if string(chars[ind]) == "0" {
			binMap[0] = append(binMap[0], strs[i])
			zCount++
		} else {
			binMap[1] = append(binMap[1], strs[i])
			oCount++
		}
	}

	if oCount == zCount {
		if strong == 0 {
			return filterLists(binMap[0], ind+1, strong)
		} else {
			return filterLists(binMap[1], ind+1, strong)
		}
	} else if oCount > zCount {
		if strong == 1 {
			return filterLists(binMap[1], ind+1, strong)
		} else {
			return filterLists(binMap[0], ind+1, strong)
		}
	} else {
		if strong == 1 {
			return filterLists(binMap[0], ind+1, strong)
		} else {
			return filterLists(binMap[1], ind+1, strong)
		}
	}
}
