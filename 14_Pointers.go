package main

import "fmt"

func main() {
	// 1: Basic variable and pointer
	var x int = 10
	var ptr *int = &x // pointer to x

	fmt.Println("x =", x)
	fmt.Println("Address of x =", &x)
	fmt.Println("ptr (holds address of x) =", ptr)
	fmt.Println("Value at ptr =", *ptr) // dereferencing

	// 2: Modify value using pointer
	*ptr = 20
	fmt.Println("Modified x through pointer =", x)

	// 3: Passing pointer to function (pass by reference)
	num := 50
	increment(&num)
	fmt.Println("After increment:", num)

	// 4: Pointer with struct
	type Person struct {
		name string
		age  int
	}

	p := Person{"Raushan", 22}
	ptrPerson := &p
	ptrPerson.age = 23 // modify using pointer
	fmt.Println("Struct through pointer:", *ptrPerson)

	// 5: Nil pointer check
	var pnil *int
	if pnil == nil {
		fmt.Println("pnil is nil")
	}
}

// Function that accepts pointer
func increment(n *int) {
	*n = *n + 1
}
