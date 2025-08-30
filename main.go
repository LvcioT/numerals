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
	context := &numerals.Context{Input: roman}

	expressions := []numerals.Expression{
		&numerals.VinculumExpression{Roman: "M", Arabic: 1000},
		&numerals.VinculumExpression{Roman: "CM", Arabic: 900},
		&numerals.VinculumExpression{Roman: "D", Arabic: 500},
		&numerals.VinculumExpression{Roman: "CD", Arabic: 400},
		&numerals.VinculumExpression{Roman: "C", Arabic: 100},
		&numerals.VinculumExpression{Roman: "XC", Arabic: 90},
		&numerals.VinculumExpression{Roman: "L", Arabic: 50},
		&numerals.VinculumExpression{Roman: "XL", Arabic: 40},
		&numerals.VinculumExpression{Roman: "X", Arabic: 10},
		&numerals.VinculumExpression{Roman: "IX", Arabic: 9},
		&numerals.VinculumExpression{Roman: "V", Arabic: 5},
		&numerals.VinculumExpression{Roman: "IV", Arabic: 4},
		&numerals.VinculumExpression{Roman: "I", Arabic: 1},
		&numerals.TerminalExpression{Roman: "M", Arabic: 1000},
		&numerals.TerminalExpression{Roman: "CM", Arabic: 900},
		&numerals.TerminalExpression{Roman: "D", Arabic: 500},
		&numerals.TerminalExpression{Roman: "CD", Arabic: 400},
		&numerals.TerminalExpression{Roman: "C", Arabic: 100},
		&numerals.TerminalExpression{Roman: "XC", Arabic: 90},
		&numerals.TerminalExpression{Roman: "L", Arabic: 50},
		&numerals.TerminalExpression{Roman: "XL", Arabic: 40},
		&numerals.TerminalExpression{Roman: "X", Arabic: 10},
		&numerals.TerminalExpression{Roman: "IX", Arabic: 9},
		&numerals.TerminalExpression{Roman: "V", Arabic: 5},
		&numerals.TerminalExpression{Roman: "IV", Arabic: 4},
		&numerals.TerminalExpression{Roman: "I", Arabic: 1},
	}

	for _, expression := range expressions {
		for {
			// Create a temporary context to avoid issues with overlapping patterns
			tempContext := &numerals.Context{Input: context.Input}
			expression.Interpret(tempContext)
			if tempContext.Input != context.Input {
				context.Input = tempContext.Input
				context.Output += tempContext.Output
			} else {
				break
			}
		}
	}

	fmt.Printf("The arabic numeral for %s is %d\n", roman, context.Output)
}
