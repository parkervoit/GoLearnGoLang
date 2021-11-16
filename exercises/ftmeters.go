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
		n, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		} else {
			meters := float64(n) * 0.3048
			fmt.Printf("%.2f meters\n", meters)
		}
	}
}
