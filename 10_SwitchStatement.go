package main

import "fmt"

func main() {

	// 1: Switch with single value per case
	day := 3
	fmt.Println("Switch with single value:")
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	default:
		fmt.Println("Invalid day")
	}

	// 2: Switch with multiple values in one case
	fmt.Println("\nSwitch with multiple values in a single case:")
	switch day {
	case 1, 2, 3:
		fmt.Println("It's the beginning of the week (Mon–Wed)")
	case 4, 5:
		fmt.Println("It's midweek (Thu–Fri)")
	case 6, 7:
		fmt.Println("Weekend (Sat–Sun)")
	default:
		fmt.Println("Invalid day number")
	}

	// 3: Switch with string cases (grade example)
	grade := "B+"
	fmt.Println("\nSwitch with string values:")
	switch grade {
	case "A":
		fmt.Println("Excellent")
	case "B", "B+":
		fmt.Println("Good job")
	case "C", "C+":
		fmt.Println("Needs improvement")
	default:
		fmt.Println("Fail or Invalid Grade")
	}

	// 4: Switch without expression (like if-else chain)
	num := 15
	fmt.Println("\nSwitch without expression:")
	switch {
	case num < 0:
		fmt.Println("Negative number")
	case num == 0:
		fmt.Println("Zero")
	case num > 0 && num <= 10:
		fmt.Println("Between 1 and 10")
	default:
		fmt.Println("Greater than 10")
	}
}
