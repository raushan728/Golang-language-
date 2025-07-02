package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// 1: Create & Write to file
	file, err := os.Create("demo.txt")
	if err != nil {
		fmt.Println("Error creating:", err)
		return
	}
	defer file.Close()

	file.WriteString("Hello Raushan!\nWelcome to Go file handling.\n")
	fmt.Println("File created and written")

	// 2: Read full file
	data, err := os.ReadFile("demo.txt")
	if err == nil {
		fmt.Println("\nFile Content:")
		fmt.Println(string(data))
	}

	// 3: Append to existing file
	file2, err := os.OpenFile("demo.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err == nil {
		defer file2.Close()
		file2.WriteString("This is an appended line.\n")
		fmt.Println("Appended to file")
	}

	// 4: Read line by line
	fmt.Println("\nReading line-by-line:")
	f, err := os.Open("demo.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// 5: Delete file (optional, comment this to keep the file)
	// err = os.Remove("demo.txt")
	// if err == nil {
	// 	fmt.Println("\nFile deleted")
	// }

}
