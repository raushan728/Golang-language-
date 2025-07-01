package main

import "fmt"

func main() {
	// 1. Basic if
	age := 18
	if age >= 18 {
		fmt.Println("You are eligible to vote.")
	}

	// 2. if-else
	num := 5
	if num%2 == 0 {
		fmt.Println("Even number")
	} else {
		fmt.Println("Odd number")
	}

	// 3. if - else if - else
	score := 75
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 60 {
		fmt.Println("Grade: B")
	} else {
		fmt.Println("Grade: C")
	}

	// 4. if with short statement
	if n := 10; n > 5 {
		fmt.Println("n is greater than 5")
	}

	// 5. if with multiple conditions
	x, y := 10, 20
	if x < y && y < 100 {
		fmt.Println("x is less than y and y is less than 100")
	}

	// 6. if with OR condition
	name := "Raushan"
	if name == "Raushan" || name == "Singh" {
		fmt.Println("Welcome", name)
	}
}
