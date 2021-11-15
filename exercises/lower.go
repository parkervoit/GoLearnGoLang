package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	word := os.Args[1]
	fmt.Println(strings.ToLower(word))
}
