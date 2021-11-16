package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input := os.Args
	s := os.Args[1]
	if len(input[1]) != 1 {
		fmt.Println("Give me a letter.")
	} else if strings.ContainsAny(s, "aeiou") {
		fmt.Printf("\"%v\" is a vowel.\n", s)
	} else if strings.ContainsAny(s, "wy") {
		fmt.Printf("\"%v\" is sometimes a vowel. Sometimes not.\n", s)
	} else {
		fmt.Printf("\"%v\" is a consonant.\n", s)
	}
}
