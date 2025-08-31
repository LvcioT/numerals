package main

import (
	"fmt"
	"os"

	"taurino.com/numerals/numerals"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <roman numeral>")
		return
	}

	roman := os.Args[1]
	result := numerals.Interpret(roman)
	fmt.Printf("The arabic numeral for %s is %d\n", roman, result)
}
