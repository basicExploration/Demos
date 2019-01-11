//练习 11.3: TestRandomPalindromes测试函数只测试了回文字符串。编写新的随机测试生成器，用于测试随机生成的非回文字符串。
package t

import (
	"math/rand"
	"testing"
	"time"
)

func TestNormal(t *testing.T) {

}
func TestRandomPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}
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
	for i := range s {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func TestNotRandomPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed) // 打印一下 种子
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := NotrandomPalindrome(rng)
		if IsNotPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}
func NotrandomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		r1 := rune(rng.Intn(0x1001))
		runes[i] = r
		runes[len(runes)-i-1] = r1
	}
	return string(runes)
}
func IsNotPalindrome(s string) bool {
	de := 0
	for i := range s {
		if s[i] != s[len(s)-i-1] {
			de++

		}
		if de > 0 {
			return true
		}
	}
	return false
}
