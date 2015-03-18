package main

// (Lines ending in "OMIT" won't show up in the presentation) OMIT
// An example of two goroutines running the same function. OMIT

import (
	"fmt"

	"./godebug"
)

func main() {
	godebug.Line()
	c := make(chan bool)

	godebug.SetTrace()
	godebug.Line()
	go func() {
		godebug.Line()
		Foo()
		godebug.Line()
		close(c)
	}()

	godebug.Line()
	Foo()

	godebug.Line()
	<-c
}

func Foo() {
	godebug.Line()
	fmt.Println("here's")
	godebug.Line()
	fmt.Println("some")
	godebug.Line()
	fmt.Println("code")
}
