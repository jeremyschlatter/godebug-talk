package main

import (
	"fmt"

	"github.com/mailgun/godebug/lib"
)

func main() {
	ints := make(chan int)
	const n = 50

	for i := 0; i < n; i++ {
		go func(i int) {
			ints <- i
		}(i)
	}

	godebug.SetTrace()
	for i := 0; i < n; i++ {
		next := <-ints
		fmt.Println(next, " ")
	}
}
