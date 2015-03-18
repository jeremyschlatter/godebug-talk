package main

// Doesn't compile.

func main() {
	// START OMIT
	godebug.Line()
	if myVar || anExpression() {
		godebug.Line()
		doSomething()
	} else if i := func() int { // HL
		godebug.ElseIfInit() // HL
		return anotherFunc() // HL
	}(); godebug.ElseIf(importantFunc(i)) { // HL
		godebug.Line()
		doSomethingElse()
	}
	// END OMIT
}
