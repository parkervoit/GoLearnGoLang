package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Args needed: [username] [password]")
		return
	} else if strings.Contains(os.Args[1], "jack") && strings.Contains(os.Args[2], "1888") {
		fmt.Printf("Access granted to \"%s\"", os.Args[1])
	} else if strings.Contains(os.Args[1], "inanc") && strings.Contains(os.Args[2], "1879") {
		fmt.Printf("Access granted to \"%s\"", os.Args[1])
	} else {
		fmt.Println("Invalid credentials")
	}

}
