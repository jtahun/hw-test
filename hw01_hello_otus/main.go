package main

import (
	"fmt"
	"strings"
)

func main() {
	startString := "Hello, OTUS!"
	reversedString := reverseString(startString)	
	fmt.Println(reversedString)
}

func reverseString(input string) string {	
	runes := []rune(input)	
	var reversed strings.Builder
	for i := len(runes) - 1; i >= 0; i-- {
		reversed.WriteRune(runes[i])
	}
	return reversed.String()
}
