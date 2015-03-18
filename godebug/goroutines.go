package main

// (Lines ending in "OMIT" won't show up in the presentation) OMIT
// An example of two goroutines running the same function. OMIT

import "fmt"

func main() {
	c := make(chan bool)

	go func() {
		Foo() // HL
		close(c)
	}()

	Foo() // HL

	<-c
}

func Foo() {
	fmt.Println("here's")
	fmt.Println("some")
	fmt.Println("code")
}
