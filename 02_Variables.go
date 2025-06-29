package main

import "fmt"

const Pi = 3.14     // Exported constant (public)
const version = 1.0 // Unexported constant (private)

var Name = "Raushan" // Exported variable
var age = 19         // Unexported variable

func main() {
	var city string = "Sheohar" // with type
	var pin = 843329            // type inference
	state := "Bihar"            // short declaration

	age = 23 // overwrite allowed
	// Pi = 3.1415  //  Not allowed (constant)

	fmt.Println(Name)
	fmt.Println(age)
	fmt.Println(city)
	fmt.Println(pin)
	fmt.Println(state)
	fmt.Println(Pi)
}
