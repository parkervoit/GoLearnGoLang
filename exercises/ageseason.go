package main

import (
	"fmt"
)

func main() {
	age := 40
	if age > 60 {
		fmt.Println("getting older")
	} else if age > 30 {
		fmt.Println("getting wiser")
	} else if age > 20 {
		fmt.Println("adulthood")
	} else if age > 10 {
		fmt.Println("Young blood")
	} else {
		fmt.Println("booting up")
	}
}
