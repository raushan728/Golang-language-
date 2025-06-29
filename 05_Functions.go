package main

import (
	"fmt"
)

// 1. No parameter, no return
func greet() {
	fmt.Println("Hello from greet()")
}

// 2. With parameter, no return
func printName(name string) {
	fmt.Println("Your name is:", name)
}

// 3. No parameter, with return
func getAppName() string {
	return "GoApp"
}

// 4. With parameter and return
func square(n int) int {
	return n * n
}

// 5. Multiple parameters
func add(a int, b int) int {
	return a + b
}

// 6. Multiple return values
func divide(x int, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}

func main() {
	// Call 1: No param, no return
	greet()

	// Call 2: Param, no return
	printName("Raushan")

	// Call 3: No param, with return
	app := getAppName()
	fmt.Println("App Name is:", app)

	// Call 4: Param, with return
	sq := square(6)
	fmt.Println("Square of 6 is:", sq)

	// Call 5: Multiple parameters
	sum := add(10, 20)
	fmt.Println("Sum of 10 and 20 is:", sum)

	// Call 6: Multiple return values
	q, r := divide(17, 5)
	fmt.Println("Quotient:", q)
	fmt.Println("Remainder:", r)
}
