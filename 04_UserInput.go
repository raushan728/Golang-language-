package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var word string
	fmt.Print("Enter a single word (using fmt.Scan): ")
	fmt.Scan(&word)
	fmt.Println("You entered (Scan):", word)

	// Flush leftover input (from Scan)
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n') // discard remaining buffer

	// Now take fresh input
	fmt.Print("Enter a full line (using bufio.Reader): ")
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	fmt.Println("You entered (Reader):", line)
}
