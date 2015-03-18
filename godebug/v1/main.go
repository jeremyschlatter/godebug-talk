package main

import (
	"fmt"

	"github.com/bradfitz/iter"
)

func main() {
	x := mul(1, 2)
	x = mul(x, x)
	if x == 4 {
		fmt.Println("It works! x == 4.")
	} else if n := 2; n == 3 {
		fmt.Println("Math is broken. Run.")
	} else {
		fmt.Println("What's going on? x ==", x)
	}
}

func add(n, m int) int {
	if n == 0 {
		return m
	}
	if m == 0 {
		return n
	}
	return n + m
}

func mul(n, m int) int {
	var r int
	for range iter.N(m) {
		r += m
	}
	return r
}