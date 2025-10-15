package data

type Case struct {
	Name   string
	Roman  string
	Arabic string
	Value  uint64
}

var BaseCases = []Case{
	{Name: "empty string to 0", Roman: "", Arabic: "0", Value: 0},
	{Name: "I to 1", Roman: "I", Arabic: "1", Value: 1},
	{Name: "II to 2", Roman: "II", Arabic: "2", Value: 2},
	{Name: "IV to 4", Roman: "IV", Arabic: "4", Value: 4},
	{Name: "V to 5", Roman: "V", Arabic: "5", Value: 5},
	{Name: "IX to 9", Roman: "IX", Arabic: "9", Value: 9},
	{Name: "X to 10", Roman: "X", Arabic: "10", Value: 10},
	{Name: "XL to 40", Roman: "XL", Arabic: "40", Value: 40},
	{Name: "L to 50", Roman: "L", Arabic: "50", Value: 50},
	{Name: "XC to 90", Roman: "XC", Arabic: "90", Value: 90},
	{Name: "C to 100", Roman: "C", Arabic: "100", Value: 100},
	{Name: "CD to 400", Roman: "CD", Arabic: "400", Value: 400},
	{Name: "D to 500", Roman: "D", Arabic: "500", Value: 500},
	{Name: "CM to 900", Roman: "CM", Arabic: "900", Value: 900},
	{Name: "M to 1000", Roman: "M", Arabic: "1000", Value: 1000},
	{Name: "MCMLXXXIX to 1989", Roman: "MCMLXXXIX", Arabic: "1989", Value: 1989},
	{Name: "MMXXIV to 2024", Roman: "MMXXIV", Arabic: "2024", Value: 2024},
	{Name: "biggest without winculum", Roman: "MMMCMXCIX", Arabic: "3999", Value: 3999},
}

var WithVinculumCases = []Case{
	{Name: "(V)I to 5001", Roman: "(V)I", Arabic: "5001", Value: 5001},
	{Name: "(IV)CMXCIX to 4999", Roman: "(IV)CMXCIX", Arabic: "4999", Value: 4999},
	{Name: "(IX)I to 9001", Roman: "(IX)I", Arabic: "9001", Value: 9001},
}

// ErrorCases contains test cases that are not valid Roman numerals.
// The current implementation does not fully validate the Roman numeral rules.
var ErrorCases = []Case{
	{Name: "invalid character", Roman: "IIA", Arabic: "0", Value: 0},
	{Name: "lowercase", Roman: "ii", Arabic: "2", Value: 2},
	{Name: "IIII to 4", Roman: "IIII", Arabic: "4", Value: 4},
	{Name: "IIIII to 5", Roman: "IIIII", Arabic: "5", Value: 5},
	{Name: "MMMM to 4000", Roman: "MMMM", Arabic: "4000", Value: 4000},
	{Name: "VV to 10", Roman: "VV", Arabic: "10", Value: 10},
	{Name: "IC to 0", Roman: "IC", Arabic: "0", Value: 0},
	{Name: "IM to 0", Roman: "IM", Arabic: "0", Value: 0},
	{Name: "IVI to 5", Roman: "IVI", Arabic: "5", Value: 5},
	{Name: "IXI to 10", Roman: "IXI", Arabic: "10", Value: 10},
	{Name: "(X)MMXXIV to 12024", Roman: "(X)MMXXIV", Arabic: "12024", Value: 12024},
}
