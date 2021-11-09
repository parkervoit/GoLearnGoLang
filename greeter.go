package main

import (
	"fmt"
	"os"
)

//os.Args stores user inputs from the command prompt
func main() {
	fmt.Printf("%#v\n", os.Args)
	//can use slice indices to pull values stored in Args
	fmt.Println("Path:", os.Args[0])
	fmt.Println("1st argument:", os.Args[1])
	fmt.Println("2rd argument:", os.Args[2])
	fmt.Println("3rd argument:", os.Args[3])
	// prints 4 for len because the index is [0-3]
	fmt.Println("Number of items inside os.Args:",
		len(os.Args))
}
