package main

import "github.com/mailgun/godebug/lib"

// Doesn't compile.

func main() {
	// START OMIT
	godebug.Line()
	if myVar || anExpression() {
		godebug.Line()
		doSomething()
		godebug.Line() // HL
	} else if importantFunc() { // HL
		godebug.Line()
		doSomethingElse()
	}
	// END OMIT
}
