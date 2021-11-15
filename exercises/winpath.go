package main

import "fmt"

func main() {
	path1 := "c:\\programfiles\\superduper\\fun.txt\n" +
		"c:\\programfiles\\really\\funny.png"
	fmt.Println("og", path1)
	path2 := `c:\programfiles\\superduper\fun.txt
c:\programfiles\\really\\funny.png`
	fmt.Println("new", path2)
}
