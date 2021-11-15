package main

import "fmt"

func main() {
	const (
		Sep = iota + 9
		Oct
		Nov
	)

	fmt.Println(Sep, Oct, Nov)
}
