//修改randomPalindrome函数，以探索IsPalindrome是否对标点和空格做了正确处理。
package main

import (
	"math/rand"
	"strings"
)

func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

func IsPalindrome(s string) bool {
	s = strings.Trim(s, ",")
	s = strings.Trim(s, " ")
	for i := range s {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
