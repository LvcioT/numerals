package main

import (
	"fmt"
	"os"

	"taurino.com/numerals/pkg"
	"taurino.com/numerals/tools"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: app <roman numeral> <arabic numeral>")
		return
	}

	roman := os.Args[1]
	arabic := os.Args[2]
	var result tools.Context
	var value string

	intRoman, e := pkg.NewInterpreter(tools.Roman)
	if e != nil {
		panic(e)
	}

	result, e = intRoman.Interpret(roman)
	if e != nil {
		panic(e)
	}

	value = result.GetTo()

	fmt.Printf("The arabic numeral for '%s' is '%s'\n", roman, value)

	intArabic, e := pkg.NewInterpreter(tools.Arabic)
	if e != nil {
		panic(e)
	}

	result, e = intArabic.Interpret(arabic)
	if e != nil {
		panic(e)
	}

	value = result.GetTo()

	fmt.Printf("The roman numeral for '%s' is '%s'\n", arabic, value)
}
