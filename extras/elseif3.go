package main

// Doesn't compile.

func main() {
	// START OMIT
	godebug.Line()
	if myVar || anExpression() {
		godebug.Line()
		doSomething()
	} else if i := anotherFunc(); godebug.ElseIf(importantFunc(i)) { // HL
		godebug.Line()
		doSomethingElse()
	}
	// END OMIT
}
