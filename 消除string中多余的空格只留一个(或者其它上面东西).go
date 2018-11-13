package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(string(OneSpace("33     33232    3232   ")))

}

func OneSpace(str string) []rune {
	result := make([]rune, 0)
	runeValue := []rune(str)
	for i := 0; i < len(str)-1; i++ {
		if unicode.IsSpace(runeValue[i]) && unicode.IsSpace(runeValue[i+1]) {
			continue
		} else if unicode.IsSpace(runeValue[i]) && !unicode.IsSpace(runeValue[i+1]) {
			result = append(result, runeValue[i])
		} else if !unicode.IsSpace(runeValue[i]) {
			result = append(result, runeValue[i])
		}
	}
	switch runeValue[len(runeValue)-1] {
	case '\t', '\n', '\v', '\f', '\r', ' ', 0x85, 0xA0:
		result = append(result, runeValue[len(runeValue)-1])
	}
	return result
}
