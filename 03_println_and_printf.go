package main

import "fmt"

func main() {
	name := "Raushan"
	age := 19

	//use of Println
	fmt.Println("Hello", name, "- Age:", age)

	//use of Printf
	fmt.Printf("Hello %s, you are %d years old.\n", name, age)

	// Incorrect use of %d for string
	// This will NOT crash, but give wrong output or warning
	// fmt.Printf("Wrong: Hello %d, age is %d\n", name, age) // name is string, %d is for int

	// Incorrect use of %s for int
	fmt.Printf("Wrong: Age: %s\n", age) // age is int, %s is for string
}
