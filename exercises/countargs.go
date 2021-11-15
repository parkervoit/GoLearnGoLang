package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		args   = os.Args
		length = len(args) - 1
	)
	if length == 0 {
		fmt.Println(" Give me args")
	} else if length == 1 {
		fmt.Printf("there is one: %q\n", args[1])
	} else if length == 2 {
		fmt.Printf("There are two: \"%s %s\"\n", args[1], args[2])
	} else {
		fmt.Printf("There are %d arguments\n", length)
	}
}
