package main

import (
	"fmt"

	"./godebug"
)

func main() {
	godebug.Line()
	fmt.Println("Hello 1")
	godebug.Line()
	if true {
		godebug.Line()
		fmt.Println("Hello 2")
	}
	godebug.Line()
	fmt.Println("Hello 3")
}
