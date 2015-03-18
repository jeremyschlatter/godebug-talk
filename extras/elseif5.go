package main

// Doesn't compile.

func main() {
	// START OMIT
	godebug.Line()
	if myVar || anExpression() {
		godebug.Line()
		doSomething()
	} else { // HL
		godebug.ElseIfInit()                              // HL
		tmp_i := anotherFunc()                            // HL
		if i := tmp_i; godebug.ElseIf(importantFunc(i)) { // HL
			godebug.Line()
			doSomethingElse()
		}
	}
	// END OMIT
}
