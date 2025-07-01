package main

import "fmt"

// 1: Defining a struct
type Person struct {
	name   string
	age    int
	gender string
}

// Embedded struct
type Address struct {
	city  string
	state string
}

type Person1 struct {
	name    string
	age     int
	Address // embedded struct
}

func main() {
	// 2: Creating struct instance (field-wise)
	var p1 Person
	p1.name = "Raushan"
	p1.age = 22
	p1.gender = "Male"
	fmt.Println("Person 1:", p1)

	// 3: Creating struct with values (inline)
	p2 := Person{"Sneha", 20, "Female"}
	fmt.Println("Person 2:", p2)

	// 4: Named fields (clear and flexible)
	p3 := Person{
		name:   "Amit",
		age:    25,
		gender: "Male",
	}
	fmt.Println("Person 3:", p3)

	// 5: Accessing individual fields
	fmt.Println("Name:", p1.name)
	fmt.Println("Age:", p1.age)

	// 6: Updating struct field
	p3.age = 26
	fmt.Println("Updated age of Person 3:", p3.age)

	// 7: Pointer to struct
	ptr := &p1
	ptr.name = "Raushan Singh" // auto dereference
	fmt.Println("Updated name using pointer:", p1.name)

	// 8: Passing struct to function
	printPerson(p2)

	// 9: Anonymous struct
	student := struct {
		id    int
		score float64
	}{
		id:    101,
		score: 88.5,
	}
	fmt.Println("Anonymous struct:", student)
	// 10: Embedded struct with direct field access
	p := Person1{
		name: "Sneha",
		age:  24,
		Address: Address{
			city:  "Delhi",
			state: "Delhi",
		},
	}

	// Direct access due to embedding
	fmt.Println("Name:", p.name)
	fmt.Println("City:", p.city)   // from embedded Address
	fmt.Println("State:", p.state) // from embedded Address

}

// Function that takes struct as parameter
func printPerson(p Person) {
	fmt.Printf("Hello %s (%s), you are %d years old.\n", p.name, p.gender, p.age)
}
