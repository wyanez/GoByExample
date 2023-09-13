package main

import "fmt"

func main() {
	// Strings, which can be added together with +
	fmt.Println("go" + "lang")

	// integers and floats
	fmt.Println("1+1=", 1+1)
	fmt.Println("7.0 / 3.0 = ", 7.0/3.0)
	fmt.Println("7 / 3 = ", 7/3)

	// Booleans, with booleans operators as you'd expect
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(true)
}
