package data

type Case struct {
	Name string
	From string
	To   uint64
}

var BaseCases = []Case{
	{Name: "empty string to 0", From: "", To: 0},
	{Name: "I to 1", From: "I", To: 1},
	{Name: "II to 2", From: "II", To: 2},
	{Name: "IV to 4", From: "IV", To: 4},
	{Name: "IIII to 4", From: "IIII", To: 4},
	{Name: "V to 5", From: "V", To: 5},
	{Name: "IX to 9", From: "IX", To: 9},
	{Name: "X to 10", From: "X", To: 10},
	{Name: "XL to 40", From: "XL", To: 40},
	{Name: "L to 50", From: "L", To: 50},
	{Name: "XC to 90", From: "XC", To: 90},
	{Name: "C to 100", From: "C", To: 100},
	{Name: "CD to 400", From: "CD", To: 400},
	{Name: "D to 500", From: "D", To: 500},
	{Name: "CM to 900", From: "CM", To: 900},
	{Name: "M to 1000", From: "M", To: 1000},
	{Name: "MCMLXXXIX to 1989", From: "MCMLXXXIX", To: 1989},
	{Name: "MMXXIV to 2024", From: "MMXXIV", To: 2024},
	{Name: "biggest without winculum", From: "MMMCMXCIX", To: 3999},
}

var WithVinculumCases = []Case{
	{Name: "(V)I to 5001", From: "(V)I", To: 5001},
	{Name: "(X)MMXXIV to 12024", From: "(X)MMXXIV", To: 12024},
	{Name: "(IV)CMXCIX to 4999", From: "(IV)CMXCIX", To: 4999},
	{Name: "(IX)I to 9001", From: "(IX)I", To: 9001},
}

// ErrorCases contains test cases that are not valid Roman numerals.
// The current implementation does not fully validate the Roman numeral rules.
var ErrorCases = []Case{
	{Name: "invalid character", From: "IIA", To: 0},
	{Name: "lowercase", From: "ii", To: 2},
	{Name: "IIIII to 5", From: "IIIII", To: 5},
	{Name: "MMMM to 4000", From: "MMMM", To: 4000},
	{Name: "VV to 10", From: "VV", To: 10},
	{Name: "IC to 0", From: "IC", To: 0},
	{Name: "IM to 0", From: "IM", To: 0},
	{Name: "IVI to 5", From: "IVI", To: 5},
	{Name: "IXI to 10", From: "IXI", To: 10},
}
