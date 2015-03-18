package main

// An example concurrent program that confuses gdb but not godebug. OMIT
// Remember to import "github.com/mailgun/godebug" and add "godebug.SetTrace()" somewhere if you want to step through it. OMIT

import "fmt"

func main() {
	ints := make(chan int)
	const n = 50

	for i := 0; i < n; i++ {
		go func(i int) {
			ints <- i
		}(i)
	}

	for i := 0; i < n; i++ {
		next := <-ints
		fmt.Println(next, " ")
	}
}
