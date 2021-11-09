package main

import (
	"fmt"
	"path"
)

// splitter for both dir and file
/*func main() {
	var dir, file string

	dir, file = path.Split("go/main.go")

	fmt.Println("dir :", dir)
	fmt.Println("file : ", file)
}*/

// splitter that just returns the file
/*func main() {
	var file string
	_, file = path.Split("go/main.go")
	fmt.Println("file : ", file)
}*/

//splitter using short declaration, no need to declare variable
func main() {
	_, file := path.Split("go/main.go")
	fmt.Println("file : ", file)
}
