package main

// An example of a debugger implementation that does not handle concurrency correctly. OMIT

import (
	"fmt"
	"sync"

	"./godebug"
)

func main() {
	godebug.Line()
	var wg sync.WaitGroup
	godebug.Line()
	wg.Add(2)
	godebug.SetTrace()
	godebug.Line()
	go Foo(&wg)
	godebug.Line()
	Foo(&wg)

	godebug.Line()
	wg.Wait()
}

func Foo(wg *sync.WaitGroup) {
	godebug.Line()
	fmt.Println("here's")
	godebug.Line()
	fmt.Println("some")
	godebug.Line()
	fmt.Println("code")
	godebug.Line()
	wg.Done()
}
