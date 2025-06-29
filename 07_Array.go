package main

import "fmt"

func main() {
	// 1. Integer array with default values (0)
	var intArr [5]int
	fmt.Println("intArr (default):", intArr)

	// 2. String array with default values (empty strings "")
	var strArr [3]string
	fmt.Println("strArr (default):", strArr)

	// 3. Declare and initialize in one line
	nums := [4]int{10, 20, 30, 40}
	fmt.Println("nums:", nums)

	// 4. Initialize without specifying size (Go infers size)
	words := [...]string{"go", "is", "awesome"}
	fmt.Println("words:", words)

	// 5. Accessing elements
	fmt.Println("First element:", nums[0])
	fmt.Println("Second word:", words[1])

	// 6. Updating elements
	nums[1] = 99
	fmt.Println("Updated nums:", nums)

	// 7. Array with index-based initialization
	arr := [5]int{1: 100, 3: 300} // only index 1 and 3 are initialized
	fmt.Println("Index-wise init:", arr)

	// 8. Length of array
	fmt.Println("Length of intArr:", len(intArr))

	// 9. Loop through array using index
	fmt.Print("Loop (with index): ")
	for i := 0; i < len(nums); i++ {
		fmt.Print(nums[i], " ")
	}
	fmt.Println()

	// 10. Loop through array using range
	fmt.Print("Loop (range): ")
	for _, val := range words {
		fmt.Print(val, " ")
	}
	fmt.Println()
}
