//练习11.5: 用表格驱动的技术扩展TestSplit测试，并打印期望的输出结果。
package main

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	s, sep := "a:b:c", ":"
	words := strings.Split(s, sep)
	if got, want := len(words), 3; got != want {
		t.Errorf("Split(%q, %q) returned %d words, want %d",
			s, sep, got, want)
	}
	// ...
}
