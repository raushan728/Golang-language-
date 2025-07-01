// From    		|	To			|	Method
// string		|	int			|	strconv.Atoi()
// string		|	float64		|	strconv.ParseFloat()
// int			|	string		|	strconv.Itoa()
// float64		|	string		|	fmt.Sprintf()
// int			|	float64		|	float64(x)
// float64		|	int			|	int(x)
// byte/slice	|	string		|	string([]byte)

package main

import (
	"fmt"
	"strconv"
)

func main() {

	// 1. string -> int
	str1 := "123"
	num1, err := strconv.Atoi(str1)
	if err == nil {
		fmt.Println("string to int:", num1)
	}

	// 2. string -> float64
	str2 := "3.14"
	num2, err := strconv.ParseFloat(str2, 64)
	if err == nil {
		fmt.Println("string to float64:", num2)
	}

	// 3. int -> string
	i := 42
	str3 := strconv.Itoa(i)
	fmt.Println("int to string:", str3)

	// 4. float64 -> string
	f := 9.81
	str4 := fmt.Sprintf("%f", f)
	fmt.Println("float64 to string:", str4)

	// 5. int -> float64
	x := 10
	y := float64(x)
	fmt.Println("int to float64:", y)

	// 6. float64 -> int
	a := 5.99
	b := int(a) // decimal part removed
	fmt.Println("float64 to int:", b)

	// 7. byte slice -> string
	bytes := []byte{'H', 'i'}
	s := string(bytes)
	fmt.Println("[]byte to string:", s)

	// 8. string -> []byte
	text := "hello"
	bs := []byte(text)
	fmt.Println("string to []byte:", bs)
}
