package main

// Doesn't compile.

func main() {
	// START OMIT
	godebug.Line()
	if myVar || anExpression() {
		godebug.Line()
		doSomething()
	} else if godebug.ElseIf(importantFunc()) { // HL
		godebug.Line()
		doSomethingElse()
	}
	// END OMIT
}
