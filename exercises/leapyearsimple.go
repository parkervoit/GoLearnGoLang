package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please give a year number")
		return
	}
	year, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Printf("%q is not a number\n", os.Args[1])
		return
	}

	var leap bool

	if year%400 == 0 || year%4 == 0 {
		leap = true
	} else if year%100 == 0 {
		leap = false
	} else {
		leap = false
	}

	if leap == true {
		fmt.Printf("%d is a leap year\n", year)
	} else {
		fmt.Printf("%d is not a leap year\n", year)
	}
}
