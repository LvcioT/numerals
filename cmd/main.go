package main

import (
	"fmt"
	"os"

	"taurino.com/numerals/pkg"
	"taurino.com/numerals/tools"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <roman numeral>")
		return
	}

	roman := os.Args[1]

	interpreter, e := pkg.NewInterpreter(tools.Arabic)
	if e != nil {
		panic(e)
	}

	result, e := interpreter.Interpret(roman)
	if e != nil {
		panic(e)
	}

	fmt.Printf("The arabic numeral for %s is %d\n", roman, result)
}
