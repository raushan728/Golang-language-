package main

import "fmt"

func main() {

	// 1: Basic for loop (like C)
	fmt.Println("Basic for loop:")
	for i := 1; i <= 5; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// 2: For loop as while loop
	fmt.Println("\nFor loop as while:")
	j := 1
	for j <= 3 {
		fmt.Print(j, " ")
		j++
	}
	fmt.Println()

	// 3: Infinite loop with break
	fmt.Println("\nInfinite loop with break:")
	count := 0
	for {
		fmt.Print(count, " ")
		count++
		if count > 2 {
			break // exit the loop
		}
	}
	fmt.Println()

	// 4: Loop with continue
	fmt.Println("\nLoop with continue (skip even numbers):")
	for i := 1; i <= 5; i++ {
		if i%2 == 0 {
			continue // skip even
		}
		fmt.Print(i, " ")
	}
	fmt.Println()

	// 5: Range with array
	fmt.Println("\nLoop with range (array):")
	arr := [4]string{"Go", "is", "super", "fast"}
	for index, value := range arr {
		fmt.Println("Index:", index, "Value:", value)
	}

	// 6:Range with slice
	fmt.Println("\nLoop with range (slice):")
	nums := []int{10, 20, 30, 40}
	for i, val := range nums {
		fmt.Printf("Index %d = %d\n", i, val)
	}

	// 7: Range with only value
	fmt.Println("\nRange with only value:")
	for _, v := range nums {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
