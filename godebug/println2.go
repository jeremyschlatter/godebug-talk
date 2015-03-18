package main

import "fmt"

func main() {
	someSubroutine()
}

func someSubroutine() {
	// START OMIT
	fmt.Println("first thing:", someTrueThing())
	fmt.Println("second thing:", someOtherTrueThing())
	// END OMIT

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
