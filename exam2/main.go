package main

import (
	"fmt"
)

/*

	สัญลักษณ์ “L” หมายความว่า ตัวเลขด้านซ้าย มีค่ามากกว่า ตัวเลขด้านขวา
	สัญลักษณ์ “R” หมายความว่า ตัวเลขด้านขวา มีค่ามากกว่า ตัวเลขด้านซ้าย
	สัญลักษณ์ “=“ หมายความว่า ตัวเลขด้านซ้าย มีค่าเท่ากับ ตัวเลขด้านขวา


	input = LLRR= output = 210122
	input = ==RLL output = 000210
	input = =LLRR output = 221012
	input = RRL=R output = 012001

*/

func decodeString(encoded string) (numberSetString string, numberSetSum int) {
	length := len(encoded)
	result := make([]int, length+1)

	result[0] = 0

	for i := 0; i < length; i++ {
		fmt.Println(i, " - ", string(encoded[i]))
		switch encoded[i] {
		case 'L':
			// fmt.Print("case: L | ")
			// fmt.Printf("result[%d] = %d + 1 = %d \n", i, result[i], result[i]+1)
			if result[i] > result[i+1] {
				continue
			}
			result[i] = result[i] + 1

		case 'R':
			// fmt.Print("case: R | ")
			// fmt.Printf("result[%d] = %d + 1 = %d \n", i+1, result[i], result[i]+1)
			result[i+1] = result[i] + 1

		case '=':
			// fmt.Print("case: = | ")
			// fmt.Printf("result[%d] = %d \n", i, result[i-1])
			// result[i] = result[i]
		}
	}
	// LLRR= 110100
	fmt.Println("result first: ", result, "\n\n")

	fmt.Println("============= reverse =============\n\n")

	//  i + 1 คือเลขตรงนั้น หรือเลขปัจจุบัน
	// fmt.Println(5, " - ", ".", "  ", result[5])
	for i := length - 1; i >= 0; i-- {
		fmt.Println(i, " - ", string(encoded[i]), "  ", result[i])
		switch encoded[i] {
		case 'L':
			// fmt.Print("case: L | ")
			// fmt.Printf("result[%d] <= result[%d+1] === %v \n", i, i, result[i] <= result[i+1])
			if result[i] <= result[i+1] {
				result[i] = result[i+1] + 1
			}

		case 'R':
			// fmt.Print("case: R | ")
			// fmt.Printf("result[%d] = %d + 1 = %d \n", i+1, result[i], result[i]+1)
			if result[i] > result[i+1] {
				result[i] = result[i] - 1
			}
		case '=':
			// fmt.Print("case: = | ")
			// fmt.Printf("result[%d] = %d \n", i, result[i-1])
			// result[i+1] = result[i]
			if result[i+1] < result[i] {
				result[i+1] = result[i]
			} else {
				result[i] = result[i+1]
			}

		}
	}

	fmt.Println("result second: ", result, "\n\n")

	return
	// minVal := result[0]
	// for _, val := range result {
	// 	if val < minVal {
	// 		minVal = val
	// 	}
	// }

	// for i := range result {
	// 	result[i] -= minVal
	// }

	// // fmt.Println("result = ", result)

	// var numberSetBuild strings.Builder
	// for _, num := range result {
	// 	numberSetBuild.WriteString(fmt.Sprintf("%d", num))
	// }
	// numberSetString = numberSetBuild.String()

	// numberSetSum = 0
	// for _, num := range result {
	// 	numberSetSum += num
	// }

	// return
}

func main() {
	// decodeString("LLRR=")
	// decodeString("==RLL")
	// decodeString("=LLRR")
	decodeString("RRL=R")
	// decodeString("LLL=R")
	// decodeString("LLLLL")
	// decodeString("RRRRR")

}
