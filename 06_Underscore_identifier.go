package main

import (
	"fmt"
	"strconv"
)

// Function that returns two values: int and error
func convertToInt(input string) (int, error) {
	result, err := strconv.Atoi(input) // string to int
	return result, err
}

func main() {
	// Case 1: We need both values
	num1, err1 := convertToInt("123")
	if err1 != nil {
		fmt.Println("Conversion failed:", err1)
	} else {
		fmt.Println("Converted number:", num1)
	}

	// Case 2: We need only the number, not the error
	num2, _ := convertToInt("456") // ignoring the error using _
	fmt.Println("Number (without checking error):", num2)

	// Case 3: We need only the error, not the result
	_, err3 := convertToInt("abc") // ignoring the int using _
	if err3 != nil {
		fmt.Println("Handled error only:", err3)
	}
}
