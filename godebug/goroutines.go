package main

// (Lines ending in "OMIT" won't show up in the presentation) OMIT
// An example of two goroutines running the same function. OMIT

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go Foo(&wg)
	Foo(&wg)

	wg.Wait()
}

func Foo(wg *sync.WaitGroup) {
	fmt.Println("here's")
	fmt.Println("some")
	fmt.Println("code")
	wg.Done()
}
