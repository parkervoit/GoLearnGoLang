package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please enter number of feet")
		return
	} else {
		n, err := strconv.ParseFloat(os.Args[1], 64)
		if err != nil {
			fmt.Printf("ERROR: %s is invalid. Please enter an int or float datatype.\n", os.Args[1])
			return
		} else {
			meters := float64(n) * 0.3048
			fmt.Printf("%.2f meters\n", meters)
		}
	}
}
