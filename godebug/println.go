package main

import "fmt"

func main() {
	someSubroutine()
}

func someSubroutine() {
	if someTrueThing() && someOtherTrueThing() {
		fmt.Println("All is good!")
		return
	}
	panic("impossible")
}

func someTrueThing() bool {
	return true
}

func someOtherTrueThing() bool {
	return 10 == 010 // whoops
}
