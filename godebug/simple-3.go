package main

import (
	"fmt"

	"./godebug"
)

func main() {
	// If you want to see the stepping behavior, uncomment the following line. (hiding this from the presentation) -> OMIT
	// godebug.SetTrace() // OMIT
	godebug.Line()
	fmt.Println("Hello, world!")
}
