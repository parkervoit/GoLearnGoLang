package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	user := os.Args[1]
	pass := os.Args[2]

	if len(os.Args) != 3 {
		fmt.Println("Args needed: [username] [password]")
	} else if strings.Contains(user, "jack") && strings.Contains(pass, "1888") {
		fmt.Printf("Access granted to \"%s\"", user)
	} else {
		fmt.Println("Invalid credentials")
	}

}
