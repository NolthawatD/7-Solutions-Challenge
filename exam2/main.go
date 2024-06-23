package main

import (
	"fmt"
	"strings"
)

func decodeString(encoded string) (string, int) {
	length := len(encoded)
	result := make([]int, length+1)

	result[0] = 0

	for i := 0; i < length; i++ {
		// fmt.Println(i, " - ", string(encoded[i]))
		switch encoded[i] {
		case 'L':
			// fmt.Printf("result[%d] = %d + 1 = %d \n", i, result[i], result[i]+1)
			if result[i] > result[i+1] {
				continue
			}
			result[i] = result[i] + 1

		case 'R':
			// fmt.Printf("result[%d] = %d + 1 = %d \n", i+1, result[i], result[i]+1)
			if result[i+1] > result[i] {
				continue
			}
			result[i+1] = result[i] + 1

		case '=':
			continue
		}
	}
	// fmt.Println("*** reverse check ***")

	for i := length - 1; i >= 0; i-- {
		// fmt.Println(i, " - ", string(encoded[i]), "  ", result[i])
		switch encoded[i] {
		case 'L':
			if result[i] <= result[i+1] {
				result[i] = result[i+1] + 1
			}

		case 'R':
			if result[i] > result[i+1] {
				result[i] = result[i] - 1
			}
		case '=':

			if i > 0 && i < length-1 {
				if encoded[i-1] != '=' {
					// fmt.Printf("skip index %d, previous encode is %s \n", i, string(encoded[i-1]))
					continue
				}
			}

			if result[i] != result[i+1] {
				if result[i+1] < result[i] {
					result[i+1] = result[i]
				} else {
					result[i] = result[i+1]
				}
			}

		}
	}

	var numberSetBuild strings.Builder
	for _, num := range result {
		numberSetBuild.WriteString(fmt.Sprintf("%d", num))
	}
	resultString := numberSetBuild.String()

	resultSum := 0
	for _, num := range result {
		resultSum += num
	}

	return resultString, resultSum
}

func isValidEncodedString(s string) bool {
	validChars := "=LR"
	for _, char := range s {
		if !strings.ContainsRune(validChars, char) {
			return false
		}
	}
	return true
}

func inputEncodeString() {
	var encoded string
	for {
		fmt.Print("Enter the encoded [ must be  L, R, = ] :  ")
		fmt.Scanln(&encoded)

		if isValidEncodedString(encoded) {
			resultString, _ := decodeString(encoded)
			fmt.Printf("input = %s, output = %s \n", encoded, resultString)
			inputEncodeString()
			// break
		} else {
			fmt.Println("Invalid encoded string. It should only contain '=', 'L', or 'R'. Please try again.")
		}
	}

}

func main() {
	inputEncodeString()
}
