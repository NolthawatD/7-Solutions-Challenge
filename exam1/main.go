package main

import "fmt"

func main() {

	//[[59], [73, 41], [52, 40, 53], [26, 53, 6, 34]]

	hirarchy1 := [][]int{
		{59},
		{73, 41},
		{52, 40, 53},
		{26, 53, 6, 34},
	}

	sumHirarchyMaxValuePath(hirarchy1)

}

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
