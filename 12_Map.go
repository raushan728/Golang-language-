package main

import "fmt"

func main() {
	// 1: Declare & Initialize map directly
	person := map[string]string{
		"name":    "Raushan",
		"country": "India",
	}
	fmt.Println("Initial map:", person)

	// 2: Create empty map using make()
	ages := make(map[string]int)
	ages["Amit"] = 25
	ages["Sneha"] = 22
	fmt.Println("Ages map:", ages)

	// 3: Accessing value by key
	fmt.Println("Name:", person["name"])
	fmt.Println("Sneha's age:", ages["Sneha"])

	// 4: Updating a value
	person["name"] = "Raushan Singh"
	fmt.Println("Updated name:", person["name"])

	// 5: Deleting a key
	delete(person, "country")
	fmt.Println("After delete:", person)

	// 6: Check if key exists (important!)
	val, ok := person["country"]
	if ok {
		fmt.Println("Country:", val)
	} else {
		fmt.Println("Key 'country' does not exist")
	}

	// 7: Looping through map with range
	fmt.Println("\nLooping through ages:")
	for name, age := range ages {
		fmt.Printf("%s is %d years old\n", name, age)
	}

	// 8: Getting map length
	fmt.Println("\nLength of ages map:", len(ages))
}
