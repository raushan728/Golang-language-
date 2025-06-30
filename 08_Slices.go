package main

import "fmt"

func main() {
	// 1. Array declaration
	arr := [5]int{10, 20, 30, 40, 50}
	fmt.Println("Array:", arr)

	// 2. Slice from array (no data copy)
	slice1 := arr[1:4] // index 1 to 3 (4 not included)
	fmt.Println("Slice from array (1:4):", slice1)

	// 3. Changing slice changes original array too
	slice1[0] = 99
	fmt.Println("Modified slice:", slice1)
	fmt.Println("Array after slice change:", arr)

	// 4. Declare slice directly (without array)
	slice2 := []string{"go", "is", "awesome"}
	fmt.Println("Direct slice:", slice2)

	// 5. Using make() to create slice
	slice3 := make([]int, 3) // length 3, capacity 3
	fmt.Println("Slice from make:", slice3)

	// 6. Assign values in make slice
	slice3[0] = 1
	slice3[1] = 2
	slice3[2] = 3
	fmt.Println("Updated make slice:", slice3)

	// 7. Append new elements (only with slices)
	slice3 = append(slice3, 4, 5)
	fmt.Println("After append:", slice3)

	// 8. Length and capacity
	fmt.Println("Length:", len(slice3))
	fmt.Println("Capacity:", cap(slice3))
}
