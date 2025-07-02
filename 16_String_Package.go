package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "  Hello GoLang World!  "

	// 1. Length (built-in)
	fmt.Println("Length:", len(str))

	// 2. Contains
	fmt.Println("Contains 'Go':", strings.Contains(str, "Go"))

	// 3. Prefix/Suffix
	fmt.Println("HasPrefix '  He':", strings.HasPrefix(str, "  He"))
	fmt.Println("HasSuffix '!  ':", strings.HasSuffix(str, "!  "))

	// 4. Count substring
	fmt.Println("Count of 'l':", strings.Count(str, "l"))

	// 5. Index of substring
	fmt.Println("Index of 'Go':", strings.Index(str, "Go"))

	// 6. Replace
	fmt.Println("Replace 'Go' with 'Java':", strings.ReplaceAll(str, "Go", "Java"))

	// 7. ToLower / ToUpper
	fmt.Println("Lowercase:", strings.ToLower(str))
	fmt.Println("Uppercase:", strings.ToUpper(str))

	// 8. Trim spaces
	fmt.Println("Trimmed:", strings.TrimSpace(str))

	// 9. Split string
	words := strings.Split(str, " ")
	fmt.Println("Split by space:", words)

	// 10. Join slice
	joined := strings.Join([]string{"Raushan", "Singh"}, " ")
	fmt.Println("Joined:", joined)

	// 11. Repeat
	fmt.Println("Repeat '*':", strings.Repeat("*", 5))

	// 12. Compare strings
	fmt.Println("Compare 'Go' vs 'go':", strings.Compare("Go", "go")) // case-sensitive

	// 13. EqualFold (ignore case)
	fmt.Println("EqualFold 'Go' & 'go':", strings.EqualFold("Go", "go"))

	// 14. Fields (auto split by space, remove extra)
	fmt.Println("Fields:", strings.Fields(str))
}
