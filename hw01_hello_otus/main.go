package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

// Use reverse without var
func main() {
	fmt.Println(reverse.String("Hello, OTUS!"))
}
