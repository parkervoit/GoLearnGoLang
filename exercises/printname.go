package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[1])

	fmt.Println("Hello", os.Args[1])
	fmt.Println("How are you?")
}
