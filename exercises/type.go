package main

import "fmt"

func main() {
	num := 3
	floater := 3.14
	greeting := "Hello, there!"
	boolean := true
	fmt.Printf("Type of %d is %T\n", num, num)
	fmt.Printf("Type of %.2f is %T\n", floater, floater)
	fmt.Printf("Type of \"%s\" is %T\n", greeting, greeting)
	fmt.Printf("Type of %t is %T\n", boolean, boolean)
}
