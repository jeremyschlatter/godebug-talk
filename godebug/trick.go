package main

import gofmt "fmt"

var fmt trickPrinter

type trickPrinter struct{}

func (trickPrinter) Println(string) {
	gofmt.Println("Hello, Bug!")
}

func main() {
	fmt.Println("Hello, World!")
}
