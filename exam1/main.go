package main

import (
	"fmt"
)

func sumHirarchyMaxValuePath(hierarchy [][]int) int {
	if len(hierarchy) == 0 {
		return 0
	}
	// first value sub array index 0
	totalSum := hierarchy[0][0]
	index := 0

	for i := 1; i < len(hierarchy); i++ {

		left := hierarchy[i][index]
		right := 0
		if index+1 < len(hierarchy[i]) {
			right = hierarchy[i][index+1]
		}

		// sum max value left or right
		if left > right {
			totalSum += left
		} else {
			totalSum += right
			index++
		}
		fmt.Printf("len(%d) |  left = %d, right = %d ==> totalSumb = %d\n", len(hierarchy[i]), left, right, totalSum)
	}
	return totalSum
}

func main() {
	result := sumHirarchyMaxValuePath(hirarchy2)
	fmt.Printf("input = %v result = %d\n", hirarchy2, result)
}
