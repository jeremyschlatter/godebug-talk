package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println(`-> fmt.Println("Hello, world!")`)
	bufio.NewScanner(os.Stdin).Scan()
	fmt.Println("Hello, world!")
}
